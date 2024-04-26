package models

import (
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/validator"
)

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Author      string    `json:"author"`
	Tags        []string  `json:"tags" validate:"required"`
	Published   bool      `json:"published"`
	PublishDate string    `json:"publish_date"`
	Created     time.Time `json:"created"`
}

func ValidateArticle(v *validator.Validator, article *Article) {

	v.Check(article.Title != "", "title", "must be provided")
	v.Check(len(article.Title) <= 100, "title", "length must not exceed 100 characters")
	v.Check(article.Author != "", "author", "must be provided")
	v.Check(len(article.Author) <= 50, "author", "length must not exceed 50 characters")
	v.Check(article.Tags != nil, "tags", "must be provided")
	v.Check(validator.Unique(article.Tags), "tags", "must contain unique values")
}
