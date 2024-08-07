CREATE TABLE IF NOT EXISTS "book"
(
    -- column
    id              VARCHAR PRIMARY KEY NOT NULL,
    title           VARCHAR(255) NOT NULL,
    isbn            VARCHAR(255) NOT NULL,
    author          VARCHAR(50) NOT NULL,
    published_date  TIMESTAMPTZ,
    genre           VARCHAR(50),
    price           BIGINT NOT NULL,
    quantity        INTEGER NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ,
    deleted_at      TIMESTAMPTZ
);