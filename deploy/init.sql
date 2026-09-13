-- Example schema for the template users module.
-- Rename schema / tables to match your domain.

CREATE SCHEMA IF NOT EXISTS app;

CREATE TABLE IF NOT EXISTS app.users (
    id         BIGSERIAL PRIMARY KEY,
    email      VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL DEFAULT '',
    last_name  VARCHAR(100) NOT NULL DEFAULT '',
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

INSERT INTO app.users (email, first_name, last_name, is_active)
VALUES
    ('alice@example.com', 'Alice', 'Example', TRUE),
    ('bob@example.com', 'Bob', 'Example', TRUE)
ON CONFLICT (email) DO NOTHING;
