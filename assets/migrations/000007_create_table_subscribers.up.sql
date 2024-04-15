CREATE TABLE
    IF NOT EXISTS blog_schema.subscribers (
        id SERIAL PRIMARY KEY,
        email CITEXT NOT NULL UNIQUE,
        created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );