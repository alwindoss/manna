package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateCrossReferences, downCreateCrossReferences)
}

func upCreateCrossReferences(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	crossReferencesQuery := `
		CREATE TABLE IF NOT EXISTS cross_references (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
				from_book TEXT,
				from_chapter INTEGER,
				from_verse INTEGER,
				to_book TEXT,
				to_chapter INTEGER,
				to_verse_start INTEGER,
				to_verse_end INTEGER,
				votes INTEGER
		);
	`
	_, err := tx.ExecContext(ctx, crossReferencesQuery)
	if err != nil {
		return err
	}

	crossReferencesIndex := `CREATE INDEX IF NOT EXISTS idx_cross_references_from_to ON cross_references (
		from_book,
		from_chapter,
		from_verse,
		to_book,
		to_chapter,
		to_verse_start,
		to_verse_end
	);`
	_, err = tx.ExecContext(ctx, crossReferencesIndex)
	if err != nil {
		return err
	}

	return nil
}

func downCreateCrossReferences(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	crossReferencesIndex := `DROP INDEX IF EXISTS idx_cross_references_from_to;`
	_, err := tx.ExecContext(ctx, crossReferencesIndex)
	if err != nil {
		return err
	}

	crossReferencesQuery := `DROP TABLE IF EXISTS cross_references;`
	_, err = tx.ExecContext(ctx, crossReferencesQuery)
	if err != nil {
		return err
	}

	return nil
}
