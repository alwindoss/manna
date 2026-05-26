-- =============================================================
--  MANNA — DuckDB Analytics Schema
--
--  DuckDB reads directly from SQLite or from exported Parquet/CSV.
--  These views and tables power the analytics dashboard:
--  reading streaks, chapter completion, verse engagement, etc.
--
--  Two patterns supported:
--    A) Attach SQLite directly:
--       ATTACH 'manna_local.db' AS local (TYPE SQLITE, READ_ONLY);
--    B) Load from exported Parquet snapshots (for Hub-wide analytics)
-- =============================================================

-- =============================================================
-- PATTERN A — Attach SQLite and create analytical views
-- =============================================================

-- Attach the local SQLite DB (run this at session start)
-- ATTACH 'path/to/manna_local.db' AS local (TYPE SQLITE, READ_ONLY);

-- =============================================================
-- 1. READING STREAK  (consecutive days with at least one session)
-- =============================================================

CREATE OR REPLACE VIEW v_daily_reading AS
SELECT
    DATE_TRUNC('day', CAST(started_at AS TIMESTAMPTZ)) AS read_date,
    COUNT(*)                                             AS session_count,
    SUM(duration_secs)                                  AS total_secs,
    COUNT(DISTINCT book_id)                             AS books_read,
    COUNT(DISTINCT chapter)                             AS chapters_read
FROM local.reading_sessions
WHERE ended_at IS NOT NULL
GROUP BY 1
ORDER BY 1;

-- Current streak (consecutive days up to today)
CREATE OR REPLACE VIEW v_reading_streak AS
WITH daily AS (
    SELECT read_date,
           read_date - INTERVAL (ROW_NUMBER() OVER (ORDER BY read_date)) DAY AS grp
    FROM v_daily_reading
),
streaks AS (
    SELECT grp,
           MIN(read_date) AS streak_start,
           MAX(read_date) AS streak_end,
           COUNT(*)       AS streak_days
    FROM daily
    GROUP BY grp
)
SELECT
    streak_start,
    streak_end,
    streak_days,
    streak_end = CURRENT_DATE OR streak_end = CURRENT_DATE - INTERVAL '1' DAY AS is_active
FROM streaks
ORDER BY streak_end DESC;

-- =============================================================
-- 2. READING PROGRESS — chapters read per book
-- =============================================================

CREATE OR REPLACE VIEW v_book_progress AS
SELECT
    rs.book_id,
    bb.name                             AS book_name,
    bb.testament,
    bb.chapter_count,
    COUNT(DISTINCT rs.chapter)          AS chapters_read,
    ROUND(COUNT(DISTINCT rs.chapter) * 100.0 / bb.chapter_count, 1) AS pct_complete,
    MIN(CAST(rs.started_at AS TIMESTAMPTZ)) AS first_read_at,
    MAX(CAST(rs.started_at AS TIMESTAMPTZ)) AS last_read_at,
    SUM(rs.duration_secs)               AS total_secs
FROM local.reading_sessions rs
JOIN local.bible_books bb ON bb.id = rs.book_id
WHERE rs.ended_at IS NOT NULL
GROUP BY rs.book_id, bb.name, bb.testament, bb.chapter_count
ORDER BY bb.sort_order;

-- =============================================================
-- 3. TIME OF DAY READING PATTERNS
-- =============================================================

CREATE OR REPLACE VIEW v_reading_by_hour AS
SELECT
    EXTRACT(HOUR FROM CAST(started_at AS TIMESTAMPTZ)) AS hour_of_day,
    COUNT(*)                                            AS session_count,
    AVG(duration_secs)                                  AS avg_duration_secs,
    SUM(duration_secs)                                  AS total_secs
FROM local.reading_sessions
WHERE ended_at IS NOT NULL
GROUP BY 1
ORDER BY 1;

CREATE OR REPLACE VIEW v_reading_by_weekday AS
SELECT
    DAYOFWEEK(CAST(started_at AS TIMESTAMPTZ))  AS weekday,   -- 0=Sun
    DAYNAME(CAST(started_at AS TIMESTAMPTZ))    AS weekday_name,
    COUNT(*)                                     AS session_count,
    AVG(duration_secs)                           AS avg_duration_secs
FROM local.reading_sessions
WHERE ended_at IS NOT NULL
GROUP BY 1, 2
ORDER BY 1;

-- =============================================================
-- 4. NOTE ACTIVITY
-- =============================================================

CREATE OR REPLACE VIEW v_note_activity AS
SELECT
    DATE_TRUNC('week', CAST(created_at AS TIMESTAMPTZ)) AS week,
    COUNT(*)                                              AS notes_created,
    COUNT(*) FILTER (WHERE is_shared = 1)                AS notes_shared,
    AVG(LENGTH(body))                                    AS avg_body_length
FROM local.notes
WHERE deleted_at IS NULL
GROUP BY 1
ORDER BY 1;

CREATE OR REPLACE VIEW v_notes_by_book AS
SELECT
    n.book_id,
    bb.name     AS book_name,
    COUNT(*)    AS note_count,
    SUM(CASE WHEN n.is_shared = 1 THEN 1 ELSE 0 END) AS shared_count
FROM local.notes n
JOIN local.bible_books bb ON bb.id = n.book_id
WHERE n.deleted_at IS NULL AND n.book_id IS NOT NULL
GROUP BY n.book_id, bb.name
ORDER BY note_count DESC;

-- =============================================================
-- 5. HIGHLIGHT HEATMAP — most highlighted verses
-- =============================================================

CREATE OR REPLACE VIEW v_highlight_heatmap AS
SELECT
    h.book_id,
    bb.name     AS book_name,
    h.chapter,
    h.verse,
    COUNT(*)    AS highlight_count,
    STRING_AGG(DISTINCT h.color, ', ') AS colors_used
FROM local.highlights h
JOIN local.bible_books bb ON bb.id = h.book_id
WHERE h.deleted_at IS NULL
GROUP BY h.book_id, bb.name, h.chapter, h.verse
ORDER BY highlight_count DESC;

-- =============================================================
-- 6. SEARCH PATTERNS
-- =============================================================

CREATE OR REPLACE VIEW v_top_searches AS
SELECT
    LOWER(query) AS query,
    COUNT(*)     AS search_count,
    AVG(result_count) AS avg_results,
    MAX(CAST(searched_at AS TIMESTAMPTZ)) AS last_searched
FROM local.search_history
GROUP BY 1
ORDER BY 2 DESC
LIMIT 50;

-- =============================================================
-- 7. READING PLAN COMPLETION RATE
-- =============================================================

CREATE OR REPLACE VIEW v_plan_progress AS
SELECT
    rp.id,
    rp.name,
    rp.total_days,
    COUNT(rpd.id)                                       AS total_readings,
    SUM(rpd.completed)                                  AS completed_readings,
    ROUND(SUM(rpd.completed) * 100.0 / COUNT(rpd.id), 1) AS pct_complete,
    rp.started_at,
    rp.completed_at
FROM local.reading_plans rp
JOIN local.reading_plan_days rpd ON rpd.plan_id = rp.id
GROUP BY rp.id, rp.name, rp.total_days, rp.started_at, rp.completed_at
ORDER BY pct_complete DESC;

-- =============================================================
-- PATTERN B — Hub-wide analytics (Parquet snapshots from PostgreSQL)
--   Export from Postgres nightly:
--     COPY user_reading_sessions TO '/exports/sessions.parquet' (FORMAT parquet);
--   Then load in DuckDB:
-- =============================================================

CREATE OR REPLACE VIEW v_hub_reading_sessions AS
SELECT * FROM read_parquet('exports/hub_reading_sessions.parquet', hive_partitioning=true);

CREATE OR REPLACE VIEW v_hub_engagement AS
SELECT
    DATE_TRUNC('day', started_at)   AS day,
    COUNT(DISTINCT user_id)         AS active_users,
    COUNT(*)                        AS total_sessions,
    AVG(duration_secs)              AS avg_session_secs,
    SUM(duration_secs)              AS total_reading_secs
FROM v_hub_reading_sessions
GROUP BY 1
ORDER BY 1 DESC;

CREATE OR REPLACE VIEW v_hub_popular_books AS
SELECT
    book_id,
    COUNT(DISTINCT user_id)   AS unique_readers,
    COUNT(*)                  AS total_sessions,
    AVG(duration_secs)        AS avg_secs_per_session
FROM v_hub_reading_sessions
GROUP BY book_id
ORDER BY unique_readers DESC;

CREATE OR REPLACE VIEW v_hub_note_engagement AS
SELECT * FROM read_parquet('exports/hub_user_notes.parquet');

CREATE OR REPLACE VIEW v_hub_top_notes AS
SELECT
    n.id,
    n.user_id,
    n.title,
    n.book_id,
    n.chapter,
    n.created_at
FROM v_hub_note_engagement n
WHERE n.is_public = true
  AND n.deleted_at IS NULL
ORDER BY n.updated_at DESC;