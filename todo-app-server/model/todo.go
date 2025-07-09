package model

import "time"

type Todo struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Tag         string
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
