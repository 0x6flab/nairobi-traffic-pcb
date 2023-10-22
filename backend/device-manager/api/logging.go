package api

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
)

var _ devicemanager.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    devicemanager.Service
}

func LoggingMiddleware(svc devicemanager.Service, logger *slog.Logger) devicemanager.Service {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) CreateDevice(ctx context.Context, token string, devices ...devicemanager.Device) (err error) {
	defer func(begin time.Time) {
		attr := slog.Group(
			"create_device",
			slog.String("took", time.Since(begin).String()),
			slog.String("count", fmt.Sprintf("%d", len(devices))),
		)
		if err != nil {
			lm.logger.Error(fmt.Sprintf("failed to creat device: %s", err), attr)
			return
		}
		lm.logger.Info("device created", attr)
	}(time.Now())

	return lm.svc.CreateDevice(ctx, token, devices...)
}

func (lm *loggingMiddleware) GetDevice(ctx context.Context, token, id string) (device devicemanager.Device, err error) {
	defer func(begin time.Time) {
		attr := slog.Group(
			"get_device",
			slog.String("took", time.Since(begin).String()),
			slog.String("id", id),
		)
		if err != nil {
			lm.logger.Error(fmt.Sprintf("failed to get device: %s", err), attr)
			return
		}
		lm.logger.Info("device retrieved", attr)
	}(time.Now())

	return lm.svc.GetDevice(ctx, token, id)
}

func (lm *loggingMiddleware) GetDevices(ctx context.Context, token string, page devicemanager.PageMetadata) (devices devicemanager.Page, err error) {
	defer func(begin time.Time) {
		attr := slog.Group(
			"get_devices",
			slog.String("took", time.Since(begin).String()),
			slog.String("offset", fmt.Sprintf("%d", page.Offset)),
			slog.String("limit", fmt.Sprintf("%d", page.Limit)),
		)
		if err != nil {
			lm.logger.Error(fmt.Sprintf("failed to get devices: %s", err), attr)
			return
		}
		lm.logger.Info("devices retrieved", attr)
	}(time.Now())

	return lm.svc.GetDevices(ctx, token, page)
}

func (lm *loggingMiddleware) UpdateDevice(ctx context.Context, token string, device devicemanager.Device) (err error) {
	defer func(begin time.Time) {
		attr := slog.Group(
			"update_device",
			slog.String("took", time.Since(begin).String()),
			slog.String("id", device.ID),
		)
		if err != nil {
			lm.logger.Error(fmt.Sprintf("failed to update device: %s", err), attr)
			return
		}
		lm.logger.Info("device updated", attr)
	}(time.Now())

	return lm.svc.UpdateDevice(ctx, token, device)
}

func (lm *loggingMiddleware) DeleteDevice(ctx context.Context, token, id string) (err error) {
	defer func(begin time.Time) {
		attr := slog.Group(
			"delete_device",
			slog.String("took", time.Since(begin).String()),
			slog.String("id", id),
		)
		if err != nil {
			lm.logger.Error(fmt.Sprintf("failed to delete device: %s", err), attr)
			return
		}
		lm.logger.Info("device deleted", attr)
	}(time.Now())

	return lm.svc.DeleteDevice(ctx, token, id)
}
