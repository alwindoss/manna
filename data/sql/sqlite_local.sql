-- =============================================================
--  MANNA — SQLite Local Database Schema
--  All tables include sync columns for Manna Hub replication
--  Sync status: 'local' | 'pending' | 'synced' | 'conflict'
-- =============================================================

PRAGMA journal_mode = WAL;
PRAGMA foreign_keys = ON;

-- =============================================================
-- 1. USER / PROFILE
-- =============================================================

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
);

-- =============================================================
-- 2. BIBLE REFERENCE DATA
-- =============================================================

CREATE TABLE IF NOT EXISTS bible_translations (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    code            TEXT NOT NULL UNIQUE,             -- 'NIV', 'ESV', 'KJV', 'NLT'
    name            TEXT NOT NULL,
    language        TEXT NOT NULL DEFAULT 'en',
    is_installed    INTEGER NOT NULL DEFAULT 0,
    installed_at    TEXT
);

CREATE TABLE IF NOT EXISTS bible_books (
    id              INTEGER PRIMARY KEY,              -- 1–66
    testament       TEXT NOT NULL CHECK (testament IN ('OT', 'NT')),
    name            TEXT NOT NULL UNIQUE,
    short_name      TEXT NOT NULL UNIQUE,             -- 'Gen', 'Exo', etc.
    chapter_count   INTEGER NOT NULL,
    sort_order      INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS bible_verses (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    translation_id  INTEGER NOT NULL REFERENCES bible_translations(id),
    book_id         INTEGER NOT NULL REFERENCES bible_books(id),
    chapter         INTEGER NOT NULL,
    verse           INTEGER NOT NULL,
    text            TEXT NOT NULL,
    UNIQUE (translation_id, book_id, chapter, verse)
);

CREATE INDEX IF NOT EXISTS idx_verses_location
    ON bible_verses (translation_id, book_id, chapter, verse);

-- =============================================================
-- 3. READING HISTORY  (used by DuckDB analytics)
-- =============================================================

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
);

CREATE INDEX IF NOT EXISTS idx_sessions_book ON reading_sessions (book_id, chapter);
CREATE INDEX IF NOT EXISTS idx_sessions_sync ON reading_sessions (sync_status);

-- =============================================================
-- 4. NOTES
-- =============================================================

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
);

CREATE INDEX IF NOT EXISTS idx_notes_book     ON notes (book_id, chapter);
CREATE INDEX IF NOT EXISTS idx_notes_sync     ON notes (sync_status);
CREATE INDEX IF NOT EXISTS idx_notes_deleted  ON notes (deleted_at);

-- =============================================================
-- 5. BOOKMARKS
-- =============================================================

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
);

CREATE INDEX IF NOT EXISTS idx_bookmarks_sync ON bookmarks (sync_status);

-- =============================================================
-- 6. HIGHLIGHTS
-- =============================================================

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
);

CREATE INDEX IF NOT EXISTS idx_highlights_book ON highlights (book_id, chapter);
CREATE INDEX IF NOT EXISTS idx_highlights_sync ON highlights (sync_status);

-- =============================================================
-- 7. READING PLANS
-- =============================================================

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
);

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
);

CREATE INDEX IF NOT EXISTS idx_plan_days_plan ON reading_plan_days (plan_id, day_number);

-- =============================================================
-- 8. SEARCH HISTORY  (used by DuckDB analytics)
-- =============================================================

CREATE TABLE IF NOT EXISTS search_history (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    hub_id          TEXT UNIQUE,
    query           TEXT NOT NULL,
    result_count    INTEGER,
    searched_at     TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%SZ', 'now')),
    sync_status     TEXT NOT NULL DEFAULT 'local'
                        CHECK (sync_status IN ('local','pending','synced','conflict')),
    synced_at       TEXT
);

CREATE INDEX IF NOT EXISTS idx_search_history_sync ON search_history (sync_status);

-- =============================================================
-- 9. SYNC QUEUE  (outbox pattern — tracks all pending Hub ops)
-- =============================================================

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
);

CREATE INDEX IF NOT EXISTS idx_sync_queue_pending
    ON sync_queue (entity_table, entity_id) WHERE sent_at IS NULL;

-- =============================================================
-- 10. TRIGGERS — auto-queue changes for Hub sync
-- =============================================================

-- Notes
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
END;

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
END;

CREATE TRIGGER IF NOT EXISTS trg_notes_delete
AFTER UPDATE OF deleted_at ON notes WHEN NEW.deleted_at IS NOT NULL
BEGIN
    INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
    VALUES ('notes', NEW.id, 'delete', json_object('id', NEW.id, 'hub_id', NEW.hub_id));
END;

-- Bookmarks
CREATE TRIGGER IF NOT EXISTS trg_bookmarks_insert
AFTER INSERT ON bookmarks
BEGIN
    INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
    VALUES ('bookmarks', NEW.id, 'insert', json_object(
        'id', NEW.id, 'book_id', NEW.book_id, 'chapter', NEW.chapter,
        'verse', NEW.verse, 'label', NEW.label, 'created_at', NEW.created_at
    ));
END;

-- Highlights
CREATE TRIGGER IF NOT EXISTS trg_highlights_insert
AFTER INSERT ON highlights
BEGIN
    INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
    VALUES ('highlights', NEW.id, 'insert', json_object(
        'id', NEW.id, 'book_id', NEW.book_id, 'chapter', NEW.chapter,
        'verse', NEW.verse, 'color', NEW.color, 'created_at', NEW.created_at
    ));
END;

-- Reading sessions
CREATE TRIGGER IF NOT EXISTS trg_sessions_insert
AFTER INSERT ON reading_sessions
BEGIN
    INSERT INTO sync_queue (entity_table, entity_id, operation, payload)
    VALUES ('reading_sessions', NEW.id, 'insert', json_object(
        'id', NEW.id, 'book_id', NEW.book_id, 'chapter', NEW.chapter,
        'started_at', NEW.started_at
    ));
END;