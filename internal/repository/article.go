package repository

import (
	"context"
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/db/store"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type ArticleRepo struct {
	store *store.Store
}

func (ar ArticleRepo) Create(article *models.Article) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return ar.store.Queries.CreateArticle(ctx, article)

}

func (ar ArticleRepo) Delete(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := ar.store.Queries.DeleteArticle(ctx, id)
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
