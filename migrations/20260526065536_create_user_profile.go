package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserProfile, downCreateUserProfile)
}

func upCreateUserProfile(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	userProfileQuery := `
	CREATE TABLE IF NOT EXISTS user_profile (
    id                  INTEGER PRIMARY KEY,          -- always 1 for single-user
    hub_user_id         TEXT UNIQUE,                  -- UUID from Hub after login
    display_name        TEXT NOT NULL DEFAULT 'User',
    email               TEXT,
    avatar_url          TEXT,                         -- local path or Hub CDN URL
    bio                 TEXT,
    preferred_translation TEXT NOT NULL DEFAULT 'KJV',
    theme               TEXT NOT NULL DEFAULT 'parchment',  -- 'parchment' | 'dark' | 'light'
    font_size           INTEGER NOT NULL DEFAULT 16,
    reading_goal_chapters INTEGER NOT NULL DEFAULT 1, -- chapters per day
    is_hub_connected    INTEGER NOT NULL DEFAULT 0,   -- 0=false, 1=true
    hub_access_token    TEXT,
    hub_refresh_token   TEXT,
    token_expires_at    TEXT,                         -- ISO8601
    created_at          TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    updated_at          TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now'))
	);`
	_, err := tx.ExecContext(ctx, userProfileQuery)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUserProfile(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	userProfileQuery := `DROP TABLE IF EXISTS user_profile;`
	_, err := tx.ExecContext(ctx, userProfileQuery)
	if err != nil {
		return err
	}
	return nil
}
