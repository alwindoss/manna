package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateReadingSessions, downCreateReadingSessions)
}

func upCreateReadingSessions(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	readingSessionsQuery := `
	CREATE TABLE IF NOT EXISTS reading_sessions (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,                      -- UUID, set after sync
    translation_id  INTEGER NOT NULL REFERENCES bible_translations(id),
    book_id         INTEGER NOT NULL REFERENCES bible_books(id),
    chapter         INTEGER NOT NULL,
    started_at      TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    ended_at        TEXT,
    duration_secs   INTEGER,
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT
	);`
	_, err := tx.ExecContext(ctx, readingSessionsQuery)
	if err != nil {
		return err
	}

	readingSessionsBookIndex := `CREATE INDEX IF NOT EXISTS idx_sessions_book ON reading_sessions (book_id, chapter);`
	_, err = tx.ExecContext(ctx, readingSessionsBookIndex)
	if err != nil {
		return err
	}

	readingSessionsSyncIndex := `CREATE INDEX IF NOT EXISTS idx_sessions_sync ON reading_sessions (sync_status);`
	_, err = tx.ExecContext(ctx, readingSessionsSyncIndex)
	if err != nil {
		return err
	}

	return nil
}

func downCreateReadingSessions(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	readingSessionsSyncIndex := `DROP INDEX IF EXISTS idx_sessions_sync;`
	_, err := tx.ExecContext(ctx, readingSessionsSyncIndex)
	if err != nil {
		return err
	}
	readingSessionsBookIndex := `DROP INDEX IF EXISTS idx_sessions_book;`
	_, err = tx.ExecContext(ctx, readingSessionsBookIndex)
	if err != nil {
		return err
	}

	readingSessionsQuery := `DROP TABLE IF EXISTS reading_sessions;`
	_, err = tx.ExecContext(ctx, readingSessionsQuery)
	if err != nil {
		return err
	}

	return nil
}
