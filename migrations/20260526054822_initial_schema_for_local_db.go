package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInitialSchemaForLocalDb, downInitialSchemaForLocalDb)
}

func upInitialSchemaForLocalDb(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func downInitialSchemaForLocalDb(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS users;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
