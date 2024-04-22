package repository

import (
	"context"
	"strings"
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/db/store"
	"github.com/Arafetki/my-portfolio-api/internal/models"
)

type SubscriberRepo struct {
	store *store.Store
}

func (subRepo SubscriberRepo) Create(sub *models.Subscriber) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := subRepo.store.CreateSubscriber(ctx, sub)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateKeyValue
		}
		return err
	}

	return nil
}

func (subRepo SubscriberRepo) Delete(email string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := subRepo.store.DeleteSubscriber(ctx, email)
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

func (subRepo SubscriberRepo) GetAllEmails() ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	emails, err := subRepo.store.Queries.GetAllSubscribers(ctx)

	return emails, err

}
