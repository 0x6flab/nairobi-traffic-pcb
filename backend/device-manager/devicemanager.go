package devicemanager

import (
	"context"
	"time"

	"github.com/goccy/go-json"
)

const ServiceName = "device-manager"

var (
	// Version represents the last service git tag in git history.
	// -ldflags "-X 'github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager.Version=0.0.0'".
	Version = "0.0.0"
	// Commit represents the service git commit hash.
	// -ldflags "-X 'github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager.Commit=ffffffff'".
	Commit = "ffffffff"
	// BuildTime represetns the service build time.
	// -ldflags "-X 'github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager.BuildTime=1970-01-01_00:00:00'".
	BuildTime = "1970-01-01_00:00:00"
)

// Status represents the status of a device.
type Status uint8

const (
	// Enabled represents the enabled status of a device. It is able to connect to the backend.
	Enabled Status = iota
	// Disabled represents the disabled status of a device. It is not able to connect to the backend.
	Disabled
)

// Device represents a device.
type Device struct {
	Name      string                 `json:"name"`
	Owner     string                 `json:"owner"`
	ID        string                 `json:"id"`
	Key       string                 `json:"key"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Status    Status                 `json:"status"`
}

type Page struct {
	Total   uint64   `json:"total" db:"total"`
	Offset  uint64   `json:"offset" db:"offset"`
	Limit   uint64   `json:"limit" db:"limit"`
	Devices []Device `json:"devices" db:"devices"`
}

type PageMetadata struct {
	Offset uint64 `json:"offset" db:"offset"`
	Limit  uint64 `json:"limit" db:"limit"`
}

type Service interface {
	// CreateDevice adds a new device to the device manager.
	CreateDevice(ctx context.Context, token string, device ...Device) error

	// GetDevice retrieves a device from the device manager.
	GetDevice(ctx context.Context, token, id string) (device Device, err error)

	// GetDevices retrieves all devices from the device manager.
	GetDevices(ctx context.Context, token string, page PageMetadata) (devices Page, err error)

	// UpdateDevice updates a device in the device manager.
	UpdateDevice(ctx context.Context, token string, device Device) error

	// DeleteDevice deletes a device from the device manager.
	DeleteDevice(ctx context.Context, token, id string) error
}

// Repository represents a device repository.
type Repository interface {
	// Create creates a new device.
	Create(ctx context.Context, device ...Device) error

	// Read retrieves a device.
	Read(ctx context.Context, id string) (device Device, err error)

	// ReadAll retrieves all devices.
	ReadAll(ctx context.Context, page PageMetadata) (devices Page, err error)

	// Update updates a device.
	Update(ctx context.Context, device Device) error

	// Delete deletes a device.
	Delete(ctx context.Context, id string) error
}

func (status Status) String() string {
	switch status {
	case Enabled:
		return "enabled"
	case Disabled:
		return "disabled"
	default:
		return "unknown"
	}
}

func (p Page) MarshalJSON() ([]byte, error) {
	if p.Devices == nil {
		p.Devices = []Device{}
	}

	type Alias Page
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(p),
	})

}
