package repository

import (
	database "github.com/Arafetki/my-portfolio-api/internal/db"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type ArticleRepo struct {
	db *database.DB
}

func (ar ArticleRepo) Create(article *models.Article) error {
	return nil
}
