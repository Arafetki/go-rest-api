package models

import (
	"time"
)

type Subscriber struct {
	ID      int       `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}
