package store

import (
	"context"
	"fmt"

	database "github.com/Arafetki/my-portfolio-api/internal/db"
)

type Store struct {
	*Queries
	db *database.DB
}

func NewStore(db *database.DB) *Store {
	return &Store{
		db:      db,
		Queries: NewQueries(db),
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	q := NewQueries(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err

	}

	return tx.Commit()
}
