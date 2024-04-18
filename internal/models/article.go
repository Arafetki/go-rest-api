package models

import "time"

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Author      string    `json:"author"`
	Published   bool      `json:"published"`
	PublishDate time.Time `json:"publish_date"`
	Created     time.Time `json:"created"`
}
