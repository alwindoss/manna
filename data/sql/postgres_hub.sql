-- =============================================================
--  MANNA HUB — PostgreSQL Server Schema
--  Handles: auth, sync mirror of local data, social features
-- =============================================================

-- Extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";       -- fast text search on posts/notes

-- =============================================================
-- 1. USERS & AUTH
-- =============================================================

CREATE TABLE IF NOT EXISTS users (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email           TEXT NOT NULL UNIQUE,
    username        TEXT NOT NULL UNIQUE,
    display_name    TEXT NOT NULL,
    avatar_url      TEXT,
    bio             TEXT,
    password_hash   TEXT NOT NULL,
    email_verified  BOOLEAN NOT NULL DEFAULT FALSE,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_seen_at    TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS user_sessions (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    device_name     TEXT,                             -- 'James's MacBook', 'Desktop'
    refresh_token   TEXT NOT NULL UNIQUE,
    expires_at      TIMESTAMPTZ NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_used_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_sessions_user ON user_sessions (user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_token ON user_sessions (refresh_token);

-- =============================================================
-- 2. SYNC — MIRROR OF LOCAL DATA
--    Each table mirrors its SQLite counterpart.
--    The desktop app pushes via sync_queue → Hub API → these tables.
-- =============================================================

CREATE TABLE IF NOT EXISTS user_notes (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    local_id        INTEGER,                          -- SQLite id on origin device
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title           TEXT NOT NULL DEFAULT 'Untitled',
    body            TEXT NOT NULL DEFAULT '',
    book_id         INTEGER,
    chapter         INTEGER,
    verse_start     INTEGER,
    verse_end       INTEGER,
    is_public       BOOLEAN NOT NULL DEFAULT FALSE,
    tags            TEXT[] DEFAULT '{}',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_unotes_user    ON user_notes (user_id);
CREATE INDEX IF NOT EXISTS idx_unotes_book    ON user_notes (book_id, chapter);
CREATE INDEX IF NOT EXISTS idx_unotes_public  ON user_notes (is_public) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_unotes_search  ON user_notes USING gin(to_tsvector('english', title || ' ' || body));

CREATE TABLE IF NOT EXISTS user_bookmarks (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    local_id        INTEGER,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    translation_code TEXT NOT NULL DEFAULT 'NIV',
    book_id         INTEGER NOT NULL,
    chapter         INTEGER NOT NULL,
    verse           INTEGER NOT NULL,
    label           TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    UNIQUE (user_id, book_id, chapter, verse)
);

CREATE INDEX IF NOT EXISTS idx_bookmarks_user ON user_bookmarks (user_id);

CREATE TABLE IF NOT EXISTS user_highlights (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    local_id        INTEGER,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    translation_code TEXT NOT NULL DEFAULT 'NIV',
    book_id         INTEGER NOT NULL,
    chapter         INTEGER NOT NULL,
    verse           INTEGER NOT NULL,
    color           TEXT NOT NULL DEFAULT '#e8c06a',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    UNIQUE (user_id, book_id, chapter, verse)
);

CREATE INDEX IF NOT EXISTS idx_highlights_user ON user_highlights (user_id);

CREATE TABLE IF NOT EXISTS user_reading_sessions (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    local_id        INTEGER,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    translation_code TEXT NOT NULL DEFAULT 'NIV',
    book_id         INTEGER NOT NULL,
    chapter         INTEGER NOT NULL,
    started_at      TIMESTAMPTZ NOT NULL,
    ended_at        TIMESTAMPTZ,
    duration_secs   INTEGER,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_rsessions_user ON user_reading_sessions (user_id);
CREATE INDEX IF NOT EXISTS idx_rsessions_time ON user_reading_sessions (user_id, started_at);

CREATE TABLE IF NOT EXISTS user_reading_plans (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    local_id        INTEGER,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name            TEXT NOT NULL,
    description     TEXT,
    total_days      INTEGER NOT NULL,
    started_at      TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    is_active       BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_search_history (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    local_id        INTEGER,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    query           TEXT NOT NULL,
    result_count    INTEGER,
    searched_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_search_user ON user_search_history (user_id, searched_at);

-- =============================================================
-- 3. SOCIAL — MANNA HUB
-- =============================================================

-- Posts (devotionals, reflections, prayer requests, etc.)
CREATE TABLE IF NOT EXISTS posts (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body            TEXT NOT NULL,
    post_type       TEXT NOT NULL DEFAULT 'reflection'
                        CHECK (post_type IN ('reflection','prayer','testimony','question','announcement')),
    -- Optional Bible reference
    book_id         INTEGER,
    chapter         INTEGER,
    verse_start     INTEGER,
    verse_end       INTEGER,
    translation_code TEXT,
    -- Optional linked note
    note_id         UUID REFERENCES user_notes(id) ON DELETE SET NULL,
    is_published    BOOLEAN NOT NULL DEFAULT TRUE,
    like_count      INTEGER NOT NULL DEFAULT 0,
    comment_count   INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_posts_user      ON posts (user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_posts_feed      ON posts (created_at DESC) WHERE is_published AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_posts_book      ON posts (book_id, chapter);
CREATE INDEX IF NOT EXISTS idx_posts_search    ON posts USING gin(to_tsvector('english', body));
CREATE INDEX IF NOT EXISTS idx_posts_type      ON posts (post_type, created_at DESC);

-- Comments on posts
CREATE TABLE IF NOT EXISTS comments (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id         UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    parent_id       UUID REFERENCES comments(id) ON DELETE CASCADE,  -- threading
    body            TEXT NOT NULL,
    like_count      INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_comments_post   ON comments (post_id, created_at);
CREATE INDEX IF NOT EXISTS idx_comments_parent ON comments (parent_id);

-- Likes (posts + comments)
CREATE TABLE IF NOT EXISTS likes (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id         UUID REFERENCES posts(id) ON DELETE CASCADE,
    comment_id      UUID REFERENCES comments(id) ON DELETE CASCADE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK (
        (post_id IS NOT NULL AND comment_id IS NULL) OR
        (post_id IS NULL AND comment_id IS NOT NULL)
    ),
    UNIQUE (user_id, post_id),
    UNIQUE (user_id, comment_id)
);

CREATE INDEX IF NOT EXISTS idx_likes_post    ON likes (post_id);
CREATE INDEX IF NOT EXISTS idx_likes_comment ON likes (comment_id);

-- Follows
CREATE TABLE IF NOT EXISTS follows (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    follower_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    following_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (follower_id, following_id),
    CHECK (follower_id != following_id)
);

CREATE INDEX IF NOT EXISTS idx_follows_follower  ON follows (follower_id);
CREATE INDEX IF NOT EXISTS idx_follows_following ON follows (following_id);

-- Shared notes (when a user shares a note to the Hub)
CREATE TABLE IF NOT EXISTS shared_notes (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    note_id         UUID NOT NULL REFERENCES user_notes(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title_override  TEXT,                             -- optional different public title
    like_count      INTEGER NOT NULL DEFAULT 0,
    comment_count   INTEGER NOT NULL DEFAULT 0,
    shared_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (note_id)                                  -- a note can only be shared once
);

CREATE INDEX IF NOT EXISTS idx_shared_notes_user ON shared_notes (user_id, shared_at DESC);
CREATE INDEX IF NOT EXISTS idx_shared_notes_feed ON shared_notes (shared_at DESC);

-- Comments on shared notes
CREATE TABLE IF NOT EXISTS note_comments (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    shared_note_id  UUID NOT NULL REFERENCES shared_notes(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    parent_id       UUID REFERENCES note_comments(id) ON DELETE CASCADE,
    body            TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_note_comments_note ON note_comments (shared_note_id, created_at);

-- Notifications
CREATE TABLE IF NOT EXISTS notifications (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    actor_id        UUID REFERENCES users(id) ON DELETE SET NULL,
    type            TEXT NOT NULL
                        CHECK (type IN ('like','comment','follow','mention','plan_invite','note_share')),
    post_id         UUID REFERENCES posts(id) ON DELETE CASCADE,
    note_id         UUID REFERENCES shared_notes(id) ON DELETE CASCADE,
    comment_id      UUID REFERENCES comments(id) ON DELETE CASCADE,
    is_read         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_notif_user   ON notifications (user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_notif_unread ON notifications (user_id) WHERE is_read = FALSE;

-- Hub reading plans (community plans anyone can follow)
CREATE TABLE IF NOT EXISTS hub_reading_plans (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    author_id       UUID NOT NULL REFERENCES users(id),
    name            TEXT NOT NULL,
    description     TEXT,
    total_days      INTEGER NOT NULL,
    is_published    BOOLEAN NOT NULL DEFAULT FALSE,
    follower_count  INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS hub_reading_plan_days (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    plan_id         UUID NOT NULL REFERENCES hub_reading_plans(id) ON DELETE CASCADE,
    day_number      INTEGER NOT NULL,
    book_id         INTEGER NOT NULL,
    chapter_start   INTEGER NOT NULL,
    chapter_end     INTEGER NOT NULL,
    title           TEXT,
    devotional      TEXT,
    UNIQUE (plan_id, day_number)
);

-- =============================================================
-- 4. UPDATED_AT TRIGGERS
-- =============================================================

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_posts_updated_at
    BEFORE UPDATE ON posts
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_user_notes_updated_at
    BEFORE UPDATE ON user_notes
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_shared_notes_updated_at
    BEFORE UPDATE ON shared_notes
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_hub_plans_updated_at
    BEFORE UPDATE ON hub_reading_plans
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- =============================================================
-- 5. DENORMALISED COUNTER TRIGGERS (keep like/comment counts fast)
-- =============================================================

CREATE OR REPLACE FUNCTION update_post_like_count()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        UPDATE posts SET like_count = like_count + 1 WHERE id = NEW.post_id;
    ELSIF TG_OP = 'DELETE' THEN
        UPDATE posts SET like_count = GREATEST(0, like_count - 1) WHERE id = OLD.post_id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_post_likes
    AFTER INSERT OR DELETE ON likes
    FOR EACH ROW WHEN (NEW.post_id IS NOT NULL OR OLD.post_id IS NOT NULL)
    EXECUTE FUNCTION update_post_like_count();

CREATE OR REPLACE FUNCTION update_post_comment_count()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        UPDATE posts SET comment_count = comment_count + 1 WHERE id = NEW.post_id;
    ELSIF TG_OP = 'UPDATE' AND NEW.deleted_at IS NOT NULL AND OLD.deleted_at IS NULL THEN
        UPDATE posts SET comment_count = GREATEST(0, comment_count - 1) WHERE id = NEW.post_id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_post_comments
    AFTER INSERT OR UPDATE ON comments
    FOR EACH ROW EXECUTE FUNCTION update_post_comment_count();