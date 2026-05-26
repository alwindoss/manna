package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateSyncQueue, downCreateSyncQueue)
}

func upCreateSyncQueue(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	syncQueueQuery := `
	CREATE TABLE IF NOT EXISTS sync_queue (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    entity_table    TEXT NOT NULL,                    -- 'notes', 'bookmarks', etc.
    entity_id       INTEGER NOT NULL,
    operation       TEXT NOT NULL
                        CHECK (operation IN ('insert','update','delete')),
    payload         TEXT,                             -- JSON snapshot for the Hub API
    attempts        INTEGER NOT NULL DEFAULT 0,
    last_error      TEXT,
    queued_at       TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    sent_at         TEXT
	);`
	_, err := tx.ExecContext(ctx, syncQueueQuery)
	if err != nil {
		return err
	}

	syncQueuePendingIndex := `
	CREATE INDEX IF NOT EXISTS idx_sync_queue_pending
    ON sync_queue (entity_table, entity_id) WHERE sent_at IS NULL;`
	_, err = tx.ExecContext(ctx, syncQueuePendingIndex)
	if err != nil {
		return err
	}

	return nil
}

func downCreateSyncQueue(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	syncQueuePendingIndex := `DROP INDEX IF EXISTS idx_sync_queue_pending;`
	_, err := tx.ExecContext(ctx, syncQueuePendingIndex)
	if err != nil {
		return err
	}

	syncQueueQuery := `DROP TABLE IF EXISTS sync_queue;`
	_, err = tx.ExecContext(ctx, syncQueueQuery)
	if err != nil {
		return err
	}

	return nil
}
