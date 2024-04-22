CREATE TABLE
    IF NOT EXISTS blog.articles (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        body TEXT,
        author VARCHAR(50),
        tags text[] DEFAULT '{}'::text[],
        published BOOLEAN NOT NULL DEFAULT FALSE,
        publish_date DATE,
        created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );