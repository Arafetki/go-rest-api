CREATE TABLE
    IF NOT EXISTS blog.articles (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        body TEXT,
        author VARCHAR(50),
        tags text[] NOT NULL DEFAULT '{}'::text[],
        published BOOLEAN NOT NULL DEFAULT FALSE,
        publish_date DATE NOT NULL DEFAULT '0001-01-01' ,
        created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );