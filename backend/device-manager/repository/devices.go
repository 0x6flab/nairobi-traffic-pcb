package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	"github.com/goccy/go-json"
	_ "github.com/jackc/pgx/v5/stdlib" // required for SQL access
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewRepository(db *sqlx.DB, logger *slog.Logger) devicemanager.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (repo *repository) Create(ctx context.Context, devices ...devicemanager.Device) error {
	query := `INSERT INTO devices (name, owner, id, key, metadata, created_at, updated_at, status) VALUES (:name, :owner, :id, :key, :metadata, :created_at, :updated_at, :status)`

	tx, err := repo.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Join([]error{devicemanager.ErrDB, err}...)
	}

	for _, device := range devices {
		dbDevice, err := toDBDevice(device)
		if err != nil {
			return err
		}

		if _, err := tx.NamedExecContext(ctx, query, dbDevice); err != nil {
			return errors.Join([]error{devicemanager.ErrQueryFailed, err}...)
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Join([]error{devicemanager.ErrDB, err}...)
	}

	return nil
}

func (repo *repository) Read(ctx context.Context, id string) (device devicemanager.Device, err error) {
	query := `SELECT name, owner, id, key, metadata, created_at, updated_at, status FROM devices WHERE id = :id`
	dbDevice := dbDevice{
		ID: id,
	}
	row, err := repo.db.NamedQueryContext(ctx, query, dbDevice)
	if err != nil {
		return devicemanager.Device{}, errors.Join([]error{devicemanager.ErrQueryFailed, err}...)
	}
	defer row.Close()

	if !row.Next() {
		return devicemanager.Device{}, devicemanager.ErrNotFound
	}

	if err := row.StructScan(&dbDevice); err != nil {
		return devicemanager.Device{}, errors.Join([]error{devicemanager.ErrDB, err}...)
	}

	return dbDevice.toDevice()
}

func (repo *repository) ReadAll(ctx context.Context, page devicemanager.PageMetadata) (devicesPage devicemanager.Page, err error) {
	query := `SELECT name, owner, id, key, metadata, created_at, updated_at, status FROM devices LIMIT :limit OFFSET :offset`

	rows, err := repo.db.NamedQueryContext(ctx, query, page)
	if err != nil {
		return devicemanager.Page{}, errors.Join([]error{devicemanager.ErrQueryFailed, err}...)
	}
	defer rows.Close()

	var devices []devicemanager.Device
	for rows.Next() {
		var dbDevice dbDevice
		if err := rows.StructScan(&dbDevice); err != nil {
			return devicemanager.Page{}, errors.Join([]error{devicemanager.ErrDB, err}...)
		}
		device, err := dbDevice.toDevice()
		if err != nil {
			return devicemanager.Page{}, err
		}
		devices = append(devices, device)
	}

	query = `SELECT COUNT(*) FROM devices`
	var total uint64
	if err := repo.db.GetContext(ctx, &total, query); err != nil {
		return devicemanager.Page{}, errors.Join([]error{devicemanager.ErrQueryFailed, err}...)
	}

	return devicemanager.Page{
		Total:   total,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Devices: devices,
	}, nil
}

func (repo *repository) Update(ctx context.Context, device devicemanager.Device) error {
	query := `UPDATE devices SET name = :name, owner = :owner, key = :key, metadata = :metadata, updated_at = :updated_at, status = :status WHERE id = :id`
	dbDevice, err := toDBDevice(device)
	if err != nil {
		return err
	}

	if _, err := repo.db.NamedExecContext(ctx, query, dbDevice); err != nil {
		return errors.Join([]error{devicemanager.ErrQueryFailed, err}...)
	}

	return nil
}

func (repo *repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM devices WHERE id = :id`
	dbDevice := dbDevice{
		ID: id,
	}

	if _, err := repo.db.NamedExecContext(ctx, query, dbDevice); err != nil {
		return errors.Join([]error{devicemanager.ErrQueryFailed, err}...)
	}

	return nil
}

type dbDevice struct {
	Name      string               `db:"name,omitempty"`
	Owner     string               `db:"owner,omitempty"`
	ID        string               `db:"id"`
	Key       string               `db:"key"`
	Metadata  []byte               `db:"metadata,omitempty"`
	CreatedAt time.Time            `db:"created_at"`
	UpdatedAt sql.NullTime         `db:"updated_at,omitempty"`
	Status    devicemanager.Status `db:"status"`
}

func (d *dbDevice) toDevice() (devicemanager.Device, error) {
	var metadata map[string]interface{}
	if d.Metadata != nil {
		if err := json.Unmarshal(d.Metadata, &metadata); err != nil {
			return devicemanager.Device{}, devicemanager.ErrMalformedEntity
		}
	}
	var updatedAt time.Time
	if d.UpdatedAt.Valid {
		updatedAt = d.UpdatedAt.Time
	}

	return devicemanager.Device{
		Name:      d.Name,
		Owner:     d.Owner,
		ID:        d.ID,
		Key:       d.Key,
		Metadata:  metadata,
		CreatedAt: d.CreatedAt,
		UpdatedAt: updatedAt,
		Status:    d.Status,
	}, nil
}

func toDBDevice(d devicemanager.Device) (dbDevice, error) {
	metadata := []byte("{}")
	if d.Metadata != nil {
		var err error
		metadata, err = json.Marshal(d.Metadata)
		if err != nil {
			return dbDevice{}, devicemanager.ErrMalformedEntity
		}
	}
	var updatedAt sql.NullTime
	if !d.UpdatedAt.IsZero() {
		updatedAt = sql.NullTime{
			Time:  d.UpdatedAt,
			Valid: true,
		}
	}

	return dbDevice{
		Name:      d.Name,
		Owner:     d.Owner,
		ID:        d.ID,
		Key:       d.Key,
		Metadata:  metadata,
		CreatedAt: d.CreatedAt,
		UpdatedAt: updatedAt,
		Status:    d.Status,
	}, nil
}
