package repository

import (
	"context"
	"errors"
	"time"

	database "github.com/Arafetki/my-portfolio-api/internal/db"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type ArticleRepo struct {
	db *database.DB
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func (ar ArticleRepo) Create(article *models.Article) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

func (ar ArticleRepo) GetByID(id int) (models.Article, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM blog.articles
			WHERE id=$1;`

	var article models.Article

	err := ar.db.QueryRowxContext(ctx, query, id).Scan(
		&article.ID,
		&article.Title,
		&article.Body,
		&article.Author,
		&article.Published,
		&article.PublishDate,
		&article.Created,
	)

	if err != nil {
		return models.Article{}, err
	}

	return article, nil

}

func (ar ArticleRepo) Delete(id int) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM blog.articles
			  WHERE id=$1;`

	res, err := ar.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
