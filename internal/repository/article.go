package repository

import (
	"context"
	"time"

	database "github.com/Arafetki/my-portfolio-api/internal/db"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type ArticleRepo struct {
	db *database.DB
}

func (ar ArticleRepo) Create(article *models.Article) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO blog.articles (title,body,author,published,publish_date)
			  VALUES ($1,$2,$3,$4,$5)
			  RETURNING id,created;`

	args := []any{article.Title, article.Body, article.Author, article.Published, article.PublishDate}

	err := ar.db.QueryRowxContext(ctx, query, args...).Scan(
		&article.ID,
		&article.Created,
	)
	if err != nil {
		return err
	}

	return nil
}
