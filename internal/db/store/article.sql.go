package store

import (
	"context"
	"database/sql"
	"fmt"

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

func (q *Queries) GetArticles(ctx context.Context, title string, tags []string, filters models.Filters) ([]models.Article, error) {

	query := fmt.Sprintf(`SELECT * FROM blog.articles
			 			  WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
			 			  AND (tags @> $2 OR $2 = '{}')
			  			  ORDER BY %s %s, id ASC
			  			  LIMIT $3 OFFSET $4;`, filters.SortColumn(), filters.SortDirection())

	rows, err := q.db.QueryxContext(ctx, query, title, pq.Array(tags), filters.Limit(), filters.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []models.Article{}

	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.Author,
			pq.Array(article.Tags),
			&article.Published,
			&article.PublishDate,
			&article.Created,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}
