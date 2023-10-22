package devicemanager

import (
	"context"
	"time"

	"github.com/0x6flab/namegenerator"
	"github.com/oklog/ulid/v2"
)

var _ Service = (*service)(nil)

type service struct {
	deviceRepo Repository
	auth       Auth
	namegen    namegenerator.NameGenerator
}

func NewService(deviceRepo Repository, auth Auth) Service {
	return &service{
		deviceRepo: deviceRepo,
		auth:       auth,
		namegen:    namegenerator.NewNameGenerator(),
	}
}

func (s *service) CreateDevice(ctx context.Context, token string, devices ...Device) error {
	if _, err := s.auth.Identify(token); err != nil {
		return err
	}

	var ds []Device
	for _, device := range devices {
		if device.Name == "" {
			device.Name = s.namegen.Generate()
		}
		device.ID = ulid.Make().String()
		device.Key = ulid.Make().String()
		device.CreatedAt = time.Now()
		device.UpdatedAt = time.Time{}
		device.Status = Enabled

		ds = append(ds, device)
	}

	return s.deviceRepo.Create(ctx, ds...)
}

func (s *service) GetDevice(ctx context.Context, token, id string) (device Device, err error) {
	if _, err := s.auth.Identify(token); err != nil {
		return Device{}, err
	}

	return s.deviceRepo.Read(ctx, id)
}

func (s *service) GetDevices(ctx context.Context, token string, page PageMetadata) (devices Page, err error) {
	if _, err := s.auth.Identify(token); err != nil {
		return Page{}, err
	}

	return s.deviceRepo.ReadAll(ctx, page)
}

func (s *service) UpdateDevice(ctx context.Context, token string, device Device) error {
	if _, err := s.auth.Identify(token); err != nil {
		return err
	}

	device.UpdatedAt = time.Now()
	return s.deviceRepo.Update(ctx, device)
}

func (s *service) DeleteDevice(ctx context.Context, token, id string) error {
	if _, err := s.auth.Identify(token); err != nil {
		return err
	}

	return s.deviceRepo.Delete(ctx, id)
}
