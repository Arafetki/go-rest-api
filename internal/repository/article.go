package repository

import (
	"context"
	"strings"
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/db/store"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type ArticleRepo struct {
	store *store.Store
}

type createArticleTxResult struct {
	Article models.Article
}

func (ar ArticleRepo) Create(article *models.Article, catIds []int) (*createArticleTxResult, error) {

	var result createArticleTxResult

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := ar.store.ExecTx(ctx, func(q *store.Queries) error {
		var err error
		err = q.CreateArticle(ctx, article)
		if err != nil {
			return err
		}

		err = q.AddArticleToCategories(ctx, article.ID, catIds)
		if err != nil {
			if strings.Contains(err.Error(), "violates foreign key constraint") {
				return ErrForeignKeyViolation
			}
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	result.Article = *article

	return &result, nil

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
