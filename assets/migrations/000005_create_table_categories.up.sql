CREATE TABLE
    IF NOT EXISTS blog_schema.categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) UNIQUE NOT NULL
    );