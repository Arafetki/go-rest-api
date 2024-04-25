package store

import (
	"database/sql"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"golang.org/x/net/context"
)

func (q *Queries) CreateSubscriber(ctx context.Context, sub *models.Subscriber) error {

	query := `INSERT INTO blog.subscribers (email)
			  VALUES ($1)
			  RETURNING id,created;`

	return q.db.QueryRowxContext(ctx, query, sub.Email).Scan(&sub.ID, &sub.Created)
}

func (q *Queries) DeleteSubscriber(ctx context.Context, email string) (sql.Result, error) {

	query := `DELETE FROM blog.subscribers
			  WHERE email=$1;`

	res, err := q.db.ExecContext(ctx, query, email)

	return res, err
}

func (q *Queries) GetAllSubscribers(ctx context.Context) ([]string, error) {

	query := `SELECT email FROM blog.subscribers`

	emails := []string{}

	rows, err := q.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var email string
	for rows.Next() {
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}
