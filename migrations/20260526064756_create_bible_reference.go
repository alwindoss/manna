package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateBibleReference, downCreateBibleReference)
}

func upCreateBibleReference(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	bibleTranslationsQuery := `
	CREATE TABLE IF NOT EXISTS bible_translations (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    code            TEXT NOT NULL UNIQUE,             -- 'NIV', 'ESV', 'KJV', 'NLT'
    name            TEXT NOT NULL,
    language        TEXT NOT NULL DEFAULT 'en',
    is_installed    INTEGER NOT NULL DEFAULT 0,
    installed_at    TEXT
)	;`
	_, err := tx.ExecContext(ctx, bibleTranslationsQuery)
	if err != nil {
		return err
	}

	bibleBooksQuery := `
	CREATE TABLE IF NOT EXISTS bible_books (
		id              INTEGER PRIMARY KEY,              -- 1–66
		testament       TEXT NOT NULL CHECK (testament IN ('OT', 'NT')),
		name            TEXT NOT NULL UNIQUE,
		short_name      TEXT NOT NULL UNIQUE,             -- 'Gen', 'Exo', etc.
		chapter_count   INTEGER NOT NULL,
		sort_order      INTEGER NOT NULL
		);`
	_, err = tx.ExecContext(ctx, bibleBooksQuery)
	if err != nil {
		return err
	}

	bibleVersesQuery := `
	CREATE TABLE IF NOT EXISTS bible_verses (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    translation_id  INTEGER NOT NULL REFERENCES bible_translations(id),
    book_id         INTEGER NOT NULL REFERENCES bible_books(id),
    chapter         INTEGER NOT NULL,
    verse           INTEGER NOT NULL,
    text            TEXT NOT NULL,
    UNIQUE (translation_id, book_id, chapter, verse)
	);`
	_, err = tx.ExecContext(ctx, bibleVersesQuery)
	if err != nil {
		return err
	}

	bibleVersesIndex := `
	CREATE INDEX IF NOT EXISTS idx_verses_location
    ON bible_verses (translation_id, book_id, chapter, verse);`
	_, err = tx.ExecContext(ctx, bibleVersesIndex)
	if err != nil {
		return err
	}
	return nil
}

func downCreateBibleReference(ctx context.Context, tx *sql.Tx) error {
	bibleVersesIndex := `DROP INDEX IF EXISTS idx_verses_location;`
	_, err := tx.ExecContext(ctx, bibleVersesIndex)
	if err != nil {
		return err
	}
	bibleVersesQuery := `DROP TABLE IF EXISTS bible_verses;`
	_, err = tx.ExecContext(ctx, bibleVersesQuery)
	if err != nil {
		return err
	}
	bibleBooksQuery := `DROP TABLE IF EXISTS bible_books;`
	_, err = tx.ExecContext(ctx, bibleBooksQuery)
	if err != nil {
		return err
	}
	bibleTranslationsQuery := `DROP TABLE IF EXISTS bible_translations;`
	_, err = tx.ExecContext(ctx, bibleTranslationsQuery)
	return err
}
