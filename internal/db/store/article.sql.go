package store

import (
	"context"
	"database/sql"

	"github.com/Arafetki/my-portfolio-api/internal/models"
)

func (q *Queries) CreateArticle(ctx context.Context, article *models.Article) error {

	query := `INSERT INTO blog.articles (title,body,author,published,publish_date)
			  VALUES ($1,$2,$3,$4,$5)
			  RETURNING id,created;`

	args := []any{article.Title, article.Body, article.Author, article.Published, article.PublishDate}
	return q.db.QueryRowxContext(ctx, query, args...).Scan(
		&article.ID,
		&article.Created,
	)
}

func (q *Queries) AddArticleToCategories(ctx context.Context, articleID int, catIds []int) error {

	query := `INSERT INTO blog.article_categories (article_id,category_id)
			  VALUES ($1,$2);`

	for _, v := range catIds {
		_, err := q.db.ExecContext(ctx, query, articleID, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *Queries) DeleteArticle(ctx context.Context, id int) (sql.Result, error) {

	query := `DELETE FROM blog.articles
			  WHERE id=$1`

	res, err := q.db.ExecContext(ctx, query, id)

	return res, err

}

