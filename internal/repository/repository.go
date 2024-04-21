package repository

import (
	database "github.com/Arafetki/my-portfolio-api/internal/db"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type Repository struct {
	Article interface {
		Create(ar *models.Article) error
		GetByID(id int) (models.Article, error)
	}
}

func NewRepo(db *database.DB) *Repository {
	return &Repository{
		Article: ArticleRepo{db},
	}
}
