package repository

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // required for SQL access
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func InitDatabase(ctx context.Context, url string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	_, err = migrate.ExecContext(ctx, db.DB, "postgres", migration(), migrate.Up)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func migration() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "device_01",
				Up: []string{
					// name has a max length of 64 characters because of the longest name is 53 characters
					// owner has a max length of 320 characters because of the max length of an email address
					// id and key have a max length of 26 characters because of the max length of an ulid
					`CREATE TABLE IF NOT EXISTS devices (
						name        VARCHAR(64),
						owner 	    VARCHAR(320),
						id          VARCHAR(26) PRIMARY KEY,
						key         VARCHAR(26) NOT NULL UNIQUE,
						metadata    JSONB,
						created_at  TIMESTAMP,
						updated_at  TIMESTAMP,
						status      SMALLINT NOT NULL DEFAULT 0 CHECK (status >= 0)
					)`,
				},
				Down: []string{
					`DROP TABLE IF EXISTS devices`,
				},
			},
		},
	}
}
