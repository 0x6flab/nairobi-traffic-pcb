package mqtt

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type Hook struct {
	mqtt.HookBase
	topic string
	serve *mqtt.Server
	repo  devicemanager.Repository
}

func NewHook(topic string, repo devicemanager.Repository) *Hook {
	return &Hook{
		topic: topic,
		repo:  repo,
	}
}

func (h *Hook) SetServer(s *mqtt.Server) {
	h.serve = s
}

func (h *Hook) ID() string {
	return devicemanager.ServiceName + "-mqtt-hook"
}

func (h *Hook) Provides(b byte) bool {
	return bytes.Contains( // nolint: gocritic
		[]byte{
			mqtt.OnConnect,
			mqtt.OnDisconnect,
			mqtt.OnConnectAuthenticate,
			mqtt.OnACLCheck,
			mqtt.OnSubscribed,
			mqtt.OnUnsubscribed,
		}, []byte{b})
}

func (h *Hook) Init(_ interface{}) error {
	h.Log.Info("initializing mqtt hook")

	return nil
}

func (h *Hook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	h.Log.Info(
		"mqtt client connected",
		slog.String("client_id", cl.ID),
		slog.String("username", string(cl.Properties.Username)),
	)

	return nil
}

func (h *Hook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	var attr = slog.Group(
		"mqtt client disconnected",
		slog.String("client_id", cl.ID),
		slog.String("username", string(cl.Properties.Username)),
		slog.Bool("expired", expire),
	)
	if err != nil {
		h.Log.Error(fmt.Sprintf("disconnect mqtt client with error: %s", err), attr)
		return
	}

	h.Log.Info("mqtt client disconnected", attr)
}

func (h *Hook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	device, err := h.repo.Read(context.Background(), string(cl.Properties.Username))
	if err != nil {
		return false
	}

	if device.Key != string(pk.Connect.Password) {
		return false
	}

	return true
}

func (h *Hook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	return topic == h.topic
}

func (h *Hook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, _ []byte) {
	var topics []string
	for _, t := range pk.Filters {
		topics = append(topics, string(t.Filter))
	}

	h.Log.Info(
		"mqtt client subscribed",
		slog.String("client_id", cl.ID),
		slog.String("username", string(cl.Properties.Username)),
		slog.String("topics", strings.Join(topics, ",")),
	)
}

func (h *Hook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	var topics []string
	for _, t := range pk.Filters {
		topics = append(topics, string(t.Filter))
	}

	h.Log.Info(
		"mqtt client unsubscribed",
		slog.String("client_id", cl.ID),
		slog.String("username", string(cl.Properties.Username)),
		slog.String("topics", strings.Join(topics, ",")),
	)
}
