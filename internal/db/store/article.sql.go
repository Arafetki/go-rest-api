package store

import (
	"context"
	"database/sql"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"github.com/lib/pq"
)

func (q *Queries) CreateArticle(ctx context.Context, article *models.Article) error {

	query := `INSERT INTO blog.articles (title,body,author,tags,published,publish_date)
			  VALUES ($1,$2,$3,$4,$5,$6)
			  RETURNING id,created;`

	args := []any{article.Title, article.Body, article.Author, pq.Array(article.Tags), article.Published, article.PublishDate}
	return q.db.QueryRowxContext(ctx, query, args...).Scan(
		&article.ID,
		&article.Created,
	)
}

func (q *Queries) DeleteArticle(ctx context.Context, id int) (sql.Result, error) {

	query := `DELETE FROM blog.articles
			  WHERE id=$1`

	res, err := q.db.ExecContext(ctx, query, id)

	return res, err

}
