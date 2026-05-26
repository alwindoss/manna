package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateNotes, downCreateNotes)
}

func upCreateNotes(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	notesQuery := `
	CREATE TABLE IF NOT EXISTS notes (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,                      -- UUID, set after sync
    title           TEXT NOT NULL DEFAULT 'Untitled',
    body            TEXT NOT NULL DEFAULT '',
    book_id         INTEGER REFERENCES bible_books(id),
    chapter         INTEGER,
    verse_start     INTEGER,
    verse_end       INTEGER,
    is_shared       INTEGER NOT NULL DEFAULT 0,       -- shared to Hub
    is_public       INTEGER NOT NULL DEFAULT 0,       -- visible on Hub feed
    tags            TEXT,                             -- JSON array: '["love","gospel"]'
    created_at      TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    updated_at      TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    deleted_at      TEXT,                             -- soft delete for sync
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT
	);`
	_, err := tx.ExecContext(ctx, notesQuery)
	if err != nil {
		return err
	}
	/*


	   CREATE INDEX IF NOT EXISTS idx_notes_deleted  ON notes (deleted_at);
	*/
	notesBookIndex := `CREATE INDEX IF NOT EXISTS idx_notes_book     ON notes (book_id, chapter);`
	_, err = tx.ExecContext(ctx, notesBookIndex)
	if err != nil {
		return err
	}

	notesSyncIndex := `CREATE INDEX IF NOT EXISTS idx_notes_sync     ON notes (sync_status);`
	_, err = tx.ExecContext(ctx, notesSyncIndex)
	if err != nil {
		return err
	}

	notesDeletedIndex := `CREATE INDEX IF NOT EXISTS idx_notes_deleted  ON notes (deleted_at);`
	_, err = tx.ExecContext(ctx, notesDeletedIndex)
	if err != nil {
		return err
	}
	return nil
}

func downCreateNotes(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	notesDeletedIndex := `DROP INDEX IF EXISTS idx_notes_deleted;`
	_, err := tx.ExecContext(ctx, notesDeletedIndex)
	if err != nil {
		return err
	}
	notesSyncIndex := `DROP INDEX IF EXISTS idx_notes_sync;`
	_, err = tx.ExecContext(ctx, notesSyncIndex)
	if err != nil {
		return err
	}
	notesBookIndex := `DROP INDEX IF EXISTS idx_notes_book;`
	_, err = tx.ExecContext(ctx, notesBookIndex)
	if err != nil {
		return err
	}

	notesQuery := `DROP TABLE IF EXISTS notes;`
	_, err = tx.ExecContext(ctx, notesQuery)
	if err != nil {
		return err
	}
	return nil
}
