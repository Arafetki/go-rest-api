package repository

import (
	"github.com/Arafetki/my-portfolio-api/internal/db/store"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type Repository struct {
	Article interface {
		Create(ar *models.Article) error
		Delete(id int) error
		Get(title string, tags []string, filters models.Filters) ([]models.Article, error)
	}
	Subscriber interface {
		Create(sub *models.Subscriber) error
		Delete(email string) error
		GetAllEmails() ([]string, error)
	}
}

func NewRepo(store *store.Store) *Repository {
	return &Repository{
		Article:    ArticleRepo{store},
		Subscriber: SubscriberRepo{store},
	}
}
