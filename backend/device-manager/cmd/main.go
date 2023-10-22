package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	"github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager/api"
	mqttadapter "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager/mqtt"
	"github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager/repository"
	"github.com/caarlos0/env/v9"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

type Config struct {
	LogLevel   int    `env:"NTP_DEVICE_MANAGER_LOG_LEVEL"            envDefault:"0"`
	DBURL      string `env:"NTP_DEVICE_MANAGER_DB_URL,required"`
	UserID     string `env:"NTP_DEVICE_MANAGER_USER_ID,required"`
	UserToken  string `env:"NTP_DEVICE_MANAGER_USER_TOKEN,required"`
	HTTPURL    string `env:"NTP_DEVICE_MANAGER_HTTP_URL"             envDefault:"localhost:9000"`
	MQTTURL    string `env:"NTP_DEVICE_MANAGER_MQTT_URL"             envDefault:"localhost:1883"`
	MQTTTopic  string `env:"NTP_DEVICE_MANAGER_MQTT_TOPIC"           envDefault:"ntp/maps/data"`
	ServerCert string `env:"NTP_DEVICE_MANAGER_SERVER_CERT"          envDefault:""`
	ServerKey  string `env:"NTP_DEVICE_MANAGER_SERVER_KEY"           envDefault:""`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(cfg.LogLevel)}))
	slog.SetDefault(logger)

	db, err := connectDB(ctx, cfg.DBURL, logger)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	auth := devicemanager.NewAuth(cfg.UserID, cfg.UserToken)

	svc := devicemanager.NewService(db, auth)
	svc = api.LoggingMiddleware(svc, logger)

	errs := make(chan error, 1)

	go startHTTPServer(api.MakeHandler(svc, logger), cfg, logger, errs)
	go startMQTTServer(cfg, logger, db, errs)

	go func() {
		errs <- StopSignalHandler(ctx, cancel)
	}()

	err = <-errs
	logger.Error(fmt.Sprintf("%s service stopped with error: %s", devicemanager.ServiceName, err))
}

func connectDB(ctx context.Context, url string, logger *slog.Logger) (devicemanager.Repository, error) {
	db, err := repository.InitDatabase(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	repo := repository.NewRepository(db, logger)

	return repo, nil
}

func startHTTPServer(handler http.Handler, cfg Config, logger *slog.Logger, errs chan error) {
	if cfg.ServerCert != "" || cfg.ServerKey != "" {
		logger.Info(fmt.Sprintf("%s service started using https on %s with cert %s key %s", devicemanager.ServiceName, cfg.HTTPURL, cfg.ServerCert, cfg.ServerKey))
		errs <- http.ListenAndServeTLS(cfg.HTTPURL, cfg.ServerCert, cfg.ServerKey, handler)
		return
	}
	logger.Info(fmt.Sprintf("%s service started using http on %s", devicemanager.ServiceName, cfg.HTTPURL))
	errs <- http.ListenAndServe(cfg.HTTPURL, handler)
}

func StopSignalHandler(ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()

	var err error
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	select {
	case sig := <-c:
		err = fmt.Errorf("shutting down by signal: %s", sig)
	case <-ctx.Done():
		return nil
	}

	return err
}

func startMQTTServer(cfg Config, logger *slog.Logger, repo devicemanager.Repository, errs chan error) {
	server := mqtt.New(&mqtt.Options{
		InlineClient: true,
		Logger:       logger,
	})

	hook := mqttadapter.NewHook(cfg.MQTTTopic, repo)
	hook.SetServer(server)

	if err := server.AddHook(hook, nil); err != nil {
		errs <- fmt.Errorf("failed to add %s hook: %s", devicemanager.ServiceName, err)
	}

	mqtt := listeners.NewTCP(fmt.Sprintf("%s-mqtt", devicemanager.ServiceName), cfg.MQTTURL, nil)
	if err := server.AddListener(mqtt); err != nil {
		errs <- fmt.Errorf("failed to add mqtt listener: %w", err)
	}

	if err := server.Serve(); err != nil {
		errs <- fmt.Errorf("failed to start mqtt server: %w", err)
	}
}
