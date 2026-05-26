package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateReadingPlans, downCreateReadingPlans)
}

func upCreateReadingPlans(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	readingPlanQuery := `
	CREATE TABLE IF NOT EXISTS reading_plans (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,
    name            TEXT NOT NULL,
    description     TEXT,
    total_days      INTEGER NOT NULL,
    source          TEXT NOT NULL DEFAULT 'local'
                        CHECK (source IN ('local','hub','imported')),
    hub_plan_id     TEXT,                             -- if sourced from Hub
    started_at      TEXT,
    completed_at    TEXT,
    is_active       INTEGER NOT NULL DEFAULT 0,
    created_at      TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT
	);`
	_, err := tx.ExecContext(ctx, readingPlanQuery)
	if err != nil {
		return err
	}

	readingPlanDaysQuery := `
	CREATE TABLE IF NOT EXISTS reading_plan_days (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    plan_id         INTEGER NOT NULL REFERENCES reading_plans(id) ON DELETE CASCADE,
    day_number      INTEGER NOT NULL,
    book_id         INTEGER NOT NULL REFERENCES bible_books(id),
    chapter_start   INTEGER NOT NULL,
    chapter_end     INTEGER NOT NULL,
    completed       INTEGER NOT NULL DEFAULT 0,
    completed_at    TEXT,
    UNIQUE (plan_id, day_number)
	);`
	_, err = tx.ExecContext(ctx, readingPlanDaysQuery)
	if err != nil {
		return err
	}

	readingPlanDaysIndex := `CREATE INDEX IF NOT EXISTS idx_plan_days_plan ON reading_plan_days (plan_id, day_number);`
	_, err = tx.ExecContext(ctx, readingPlanDaysIndex)
	if err != nil {
		return err
	}

	return nil
}

func downCreateReadingPlans(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	readingPlanDaysIndex := `DROP INDEX IF EXISTS idx_plan_days_plan;`
	_, err := tx.ExecContext(ctx, readingPlanDaysIndex)
	if err != nil {
		return err
	}

	readingPlanDaysQuery := `DROP TABLE IF EXISTS reading_plan_days`
	_, err = tx.ExecContext(ctx, readingPlanDaysQuery)
	if err != nil {
		return err
	}

	readingPlanQuery := `DROP TABLE IF EXISTS reading_plans`
	_, err = tx.ExecContext(ctx, readingPlanQuery)
	if err != nil {
		return err
	}

	return nil
}
