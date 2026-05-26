package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTriggers, downCreateTriggers)
}

func upCreateTriggers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	notesInsertTrigger := `
	CREATE TRIGGER IF NOT EXISTS trg_notes_insert
	AFTER INSERT ON notes
	BEGIN
		INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
		VALUES ('notes', NEW.id, 'insert', json_object(
			'id', NEW.id, 'title', NEW.title, 'body', NEW.body,
			'book_id', NEW.book_id, 'chapter', NEW.chapter,
			'verse_start', NEW.verse_start, 'verse_end', NEW.verse_end,
			'tags', NEW.tags, 'is_public', NEW.is_public,
			'created_at', NEW.created_at
		));
	END;`
	_, err := tx.ExecContext(ctx, notesInsertTrigger)
	if err != nil {
		return err
	}

	notesUpdateTrigger := `
	CREATE TRIGGER IF NOT EXISTS trg_notes_update
	AFTER UPDATE ON notes WHEN NEW.sync_status != 'synced'
	BEGIN
		UPDATE notes SET updated_at = strftime('%Y-%m-%dT%H:%M:%SZ','now'),
						sync_status = 'pending'
		WHERE id = NEW.id;
		INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
		VALUES ('notes', NEW.id, 'update', json_object(
			'id', NEW.id, 'title', NEW.title, 'body', NEW.body,
			'tags', NEW.tags, 'is_public', NEW.is_public,
			'updated_at', NEW.updated_at
		));
	END;`
	_, err = tx.ExecContext(ctx, notesUpdateTrigger)
	if err != nil {
		return err
	}

	notesDeleteTrigger := `
	CREATE TRIGGER IF NOT EXISTS trg_notes_delete
	AFTER UPDATE OF deleted_at ON notes WHEN NEW.deleted_at IS NOT NULL
	BEGIN
		INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
		VALUES ('notes', NEW.id, 'delete', json_object('id', NEW.id, 'hub_id', NEW.hub_id));
	END;`
	_, err = tx.ExecContext(ctx, notesDeleteTrigger)
	if err != nil {
		return err
	}

	bookmarksInsertTrigger := `
	CREATE TRIGGER IF NOT EXISTS trg_bookmarks_insert
	AFTER INSERT ON bookmarks
	BEGIN
		INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
		VALUES ('bookmarks', NEW.id, 'insert', json_object(
			'id', NEW.id, 'book_id', NEW.book_id, 'chapter', NEW.chapter,
			'verse', NEW.verse, 'label', NEW.label, 'created_at', NEW.created_at
		));
	END;`
	_, err = tx.ExecContext(ctx, bookmarksInsertTrigger)
	if err != nil {
		return err
	}

	highlightsInsertTrigger := `
	CREATE TRIGGER IF NOT EXISTS trg_highlights_insert
	AFTER INSERT ON highlights
	BEGIN
		INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
		VALUES ('highlights', NEW.id, 'insert', json_object(
			'id', NEW.id, 'book_id', NEW.book_id, 'chapter', NEW.chapter,
			'verse', NEW.verse, 'color', NEW.color, 'created_at', NEW.created_at
		));
	END;`
	_, err = tx.ExecContext(ctx, highlightsInsertTrigger)
	if err != nil {
		return err
	}

	sessionsInsertTrigger := `
	CREATE TRIGGER IF NOT EXISTS trg_sessions_insert
	AFTER INSERT ON reading_sessions
	BEGIN
		INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
		VALUES ('reading_sessions', NEW.id, 'insert', json_object(
			'id', NEW.id, 'book_id', NEW.book_id, 'chapter', NEW.chapter,
			'started_at', NEW.started_at
		));
	END;`
	_, err = tx.ExecContext(ctx, sessionsInsertTrigger)
	if err != nil {
		return err
	}

	return nil
}

func downCreateTriggers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	sessionsInsertTrigger := `DROP TRIGGER IF EXISTS trg_sessions_insert;`
	_, err := tx.ExecContext(ctx, sessionsInsertTrigger)
	if err != nil {
		return err
	}

	highlightsInsertTrigger := `DROP TRIGGER IF EXISTS trg_highlights_insert;`
	_, err = tx.ExecContext(ctx, highlightsInsertTrigger)
	if err != nil {
		return err
	}

	bookmarksInsertTrigger := `DROP TRIGGER IF EXISTS trg_bookmarks_insert;`
	_, err = tx.ExecContext(ctx, bookmarksInsertTrigger)
	if err != nil {
		return err
	}

	notesDeleteTrigger := `DROP TRIGGER IF EXISTS trg_notes_delete;`
	_, err = tx.ExecContext(ctx, notesDeleteTrigger)
	if err != nil {
		return err
	}

	notesUpdateTrigger := `DROP TRIGGER IF EXISTS trg_notes_update;`
	_, err = tx.ExecContext(ctx, notesUpdateTrigger)
	if err != nil {
		return err
	}

	notesInsertTrigger := `DROP TRIGGER IF EXISTS trg_notes_insert;`
	_, err = tx.ExecContext(ctx, notesInsertTrigger)
	if err != nil {
		return err
	}

	/////////////////////////////////////////////////////////////////

	return nil
}
