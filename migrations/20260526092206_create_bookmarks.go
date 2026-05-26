package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateBookmarks, downCreateBookmarks)
}

func upCreateBookmarks(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	createBookmarkQuery := `
	CREATE TABLE IF NOT EXISTS bookmarks (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,
    translation_id  INTEGER NOT NULL REFERENCES bible_translations(id),
    book_id         INTEGER NOT NULL REFERENCES bible_books(id),
    chapter         INTEGER NOT NULL,
    verse           INTEGER NOT NULL,
    label           TEXT,
    created_at      TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    deleted_at      TEXT,
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT,
    UNIQUE (translation_id, book_id, chapter, verse)
	);`
	_, err := tx.ExecContext(ctx, createBookmarkQuery)
	if err != nil {
		return err
	}

	createBookmarkIndex := `CREATE INDEX IF NOT EXISTS idx_bookmarks_sync ON bookmarks (sync_status);`
	_, err = tx.ExecContext(ctx, createBookmarkIndex)
	if err != nil {
		return err
	}

	highlightQuery := `
	CREATE TABLE IF NOT EXISTS highlights (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,
    translation_id  INTEGER NOT NULL REFERENCES bible_translations(id),
    book_id         INTEGER NOT NULL REFERENCES bible_books(id),
    chapter         INTEGER NOT NULL,
    verse           INTEGER NOT NULL,
    color           TEXT NOT NULL DEFAULT '#e8c06a',  -- hex colour
    note_id         INTEGER REFERENCES notes(id),     -- optional linked note
    created_at      TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    deleted_at      TEXT,
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT,
    UNIQUE (translation_id, book_id, chapter, verse)
	);`
	_, err = tx.ExecContext(ctx, highlightQuery)
	if err != nil {
		return err
	}

	highlighBookIndex := `CREATE INDEX IF NOT EXISTS idx_highlights_book ON highlights (book_id, chapter);`
	_, err = tx.ExecContext(ctx, highlighBookIndex)
	if err != nil {
		return err
	}

	highlightSyncIndex := `CREATE INDEX IF NOT EXISTS idx_highlights_sync ON highlights (sync_status);`
	_, err = tx.ExecContext(ctx, highlightSyncIndex)
	if err != nil {
		return err
	}

	return nil
}

func downCreateBookmarks(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	highlightSyncIndex := `DROP INDEX IF EXISTS idx_highlights_sync;`
	_, err := tx.ExecContext(ctx, highlightSyncIndex)
	if err != nil {
		return err
	}
	highlighBookIndex := `DROP INDEX IF EXISTS idx_highlights_book;`
	_, err = tx.ExecContext(ctx, highlighBookIndex)
	if err != nil {
		return err
	}
	highlightQuery := `DROP TABLE IF EXISTS highlights;`
	_, err = tx.ExecContext(ctx, highlightQuery)
	if err != nil {
		return err
	}
	bookmarkSyncIndex := `DROP INDEX IF EXISTS idx_bookmarks_sync;`
	_, err = tx.ExecContext(ctx, bookmarkSyncIndex)
	if err != nil {
		return err
	}

	bookmarkQuery := `DROP TABLE IF EXISTS bookmarks;`
	_, err = tx.ExecContext(ctx, bookmarkQuery)
	if err != nil {
		return err
	}
	return nil
}
