package repository

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrForeignKeyViolation = errors.New("foreign key violated")
	ErrDuplicateKeyValue   = errors.New("duplicate key value violates unique constraint")
)
