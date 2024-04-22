package models

import "time"

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,max=100"`
	Body        string    `json:"body"`
	Author      string    `json:"author" validate:"required,max=50"`
	Tags        []string  `json:"tags" validate:"required"`
	Published   bool      `json:"published"`
	PublishDate string    `json:"publish_date"`
	Created     time.Time `json:"created"`
}
