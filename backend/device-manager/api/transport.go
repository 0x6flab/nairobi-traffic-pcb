package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/goccy/go-json"
	"github.com/hellofresh/health-go/v5"
)

const contentType = "application/json"

var errUnsupportedContentType = errors.New("unsupported content type")

func MakeHandler(svc devicemanager.Service, logger *slog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(loggingErrorEncoder(logger, encodeError)),
	}

	r := chi.NewRouter()

	r.Route("/devices", func(r chi.Router) {
		r.Post("/", kithttp.NewServer(
			createDevicesEndpoint(svc),
			decodeCreateDevicesReq,
			encodeResponse,
			opts...,
		).ServeHTTP)
		r.Get("/", kithttp.NewServer(
			getDevicesEndpoint(svc),
			decodeGetDevicesReq,
			encodeResponse,
			opts...,
		).ServeHTTP)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", kithttp.NewServer(
				getDeviceEndpoint(svc),
				decodeGetDeviceReq,
				encodeResponse,
				opts...,
			).ServeHTTP)
			r.Put("/", kithttp.NewServer(
				updateDeviceEndpoint(svc),
				decodeUpdateDeviceReq,
				encodeResponse,
				opts...,
			).ServeHTTP)
			r.Delete("/", kithttp.NewServer(
				deleteDeviceEndpoint(svc),
				decodeDeleteDeviceReq,
				encodeResponse,
				opts...,
			).ServeHTTP)
		})
	})

	r.Get("/health", Health())

	return r
}

func decodeCreateDevicesReq(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedContentType
	}

	var req createDevicesReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.token = ExtractBearerToken(r)

	return req, nil
}

func decodeGetDeviceReq(_ context.Context, r *http.Request) (interface{}, error) {
	return getDeviceReq{
		token: ExtractBearerToken(r),
		id:    chi.URLParam(r, "id"),
	}, nil
}

func decodeGetDevicesReq(_ context.Context, r *http.Request) (interface{}, error) {
	offset := r.URL.Query().Get("offset")
	o, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		o = 0
	}

	limit := r.URL.Query().Get("limit")
	l, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		l = 10
	}

	return getDevicesReq{
		token: ExtractBearerToken(r),
		page: devicemanager.PageMetadata{
			Offset: o,
			Limit:  l,
		},
	}, nil
}

func decodeUpdateDeviceReq(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedContentType
	}

	var req updateDeviceReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.token = ExtractBearerToken(r)

	return req, nil
}

func decodeDeleteDeviceReq(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteDeviceReq{
		token: ExtractBearerToken(r),
		id:    chi.URLParam(r, "id"),
	}, nil
}

func ExtractBearerToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if !strings.HasPrefix(token, "Bearer ") {
		return ""
	}

	return strings.TrimPrefix(token, "Bearer ")
}

func Health() http.HandlerFunc {
	h, _ := health.New(
		health.WithComponent(
			health.Component{
				Name:    devicemanager.ServiceName,
				Version: devicemanager.Version,
			},
		), health.WithChecks(
			health.Config{
				Name:      "cockroachdb",
				Timeout:   time.Second * 5,
				SkipOnErr: true,
				Check: func(ctx context.Context) error {
					// TODO: cockroachdb health check implementation goes here
					return nil
				},
			},
		), health.WithSystemInfo(),
	)

	return h.HandlerFunc
}

func loggingErrorEncoder(logger *slog.Logger, enc kithttp.ErrorEncoder) kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, err, w)
	}
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)
	switch {
	case errors.Is(err, devicemanager.ErrMalformedEntity):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, devicemanager.ErrAuthentication):
		w.WriteHeader(http.StatusUnauthorized)
	case errors.Is(err, devicemanager.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, errUnsupportedContentType):
		w.WriteHeader(http.StatusUnsupportedMediaType)
	case errors.Is(err, devicemanager.ErrDB):
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
