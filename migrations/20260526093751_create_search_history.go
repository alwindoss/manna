package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateSearchHistory, downCreateSearchHistory)
}

func upCreateSearchHistory(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	searchHistoryQuery := `
	CREATE TABLE IF NOT EXISTS search_history (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,
    query           TEXT NOT NULL,
    result_count    INTEGER,
    searched_at     TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT
	);`
	_, err := tx.ExecContext(ctx, searchHistoryQuery)
	if err != nil {
		return err
	}

	searchHistorySyncIndex := `CREATE INDEX IF NOT EXISTS idx_search_history_sync ON search_history (sync_status);`
	_, err = tx.ExecContext(ctx, searchHistorySyncIndex)
	if err != nil {
		return err
	}
	return nil
}

func downCreateSearchHistory(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	searchHistorySyncIndex := `DROP INDEX IF EXISTS idx_search_history_sync;`
	_, err := tx.ExecContext(ctx, searchHistorySyncIndex)
	if err != nil {
		return err
	}

	searchHistoryQuery := `DROP TABLE IF EXISTS search_history;`
	_, err = tx.ExecContext(ctx, searchHistoryQuery)
	if err != nil {
		return err
	}

	return nil
}
