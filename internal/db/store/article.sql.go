package store

import (
	"context"

	"github.com/Arafetki/my-portfolio-api/internal/models"
)

const createArticleQuery = `INSERT INTO blog.articles (title,body,author,published,publish_date)
							VALUES ($1,$2,$3,$4,$5)
							RETURNING id,created;`

func (q *Queries) CreateArticle(ctx context.Context, article *models.Article) error {

	args := []any{article.Title, article.Body, article.Author, article.Published, article.PublishDate}
	return q.db.QueryRowxContext(ctx, createArticleQuery, args...).Scan(
		&article.ID,
		&article.Created,
	)
}

const createArticleCategoriesQuery = `INSERT INTO blog.article_categories (article_id,category_id)
									  VALUES ($1,$2);`

func (q *Queries) CreateArticleCategories(ctx context.Context, articleID int, catIds []int) error {

	for _, v := range catIds {
		_, err := q.db.ExecContext(ctx, createArticleCategoriesQuery, articleID, v)
		if err != nil {
			return err
		}
	}

	return nil
}
