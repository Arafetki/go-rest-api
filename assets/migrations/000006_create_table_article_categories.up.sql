CREATE TABLE
    IF NOT EXISTS blog_schema.article_categories (
        article_id INT,
        category_id INT,
        PRIMARY KEY (article_id, category_id),
        FOREIGN KEY (article_id) REFERENCES blog_schema.articles (id),
        FOREIGN KEY (category_id) REFERENCES blog_schema.categories (id)
    );