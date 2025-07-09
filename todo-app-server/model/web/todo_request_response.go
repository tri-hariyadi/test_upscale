package web

import "time"

type CreateTodoRequest struct {
	//UserID      int64     `json:"user_id"`
	Title       string    `validate:"required,min=3" json:"title"`
	Description string    `validate:"required,min=10" json:"description"`
	Tag         string    `validate:"required" json:"tag"`
	DueDate     time.Time `validate:"required" json:"due_date"`
}

type TodoResponse struct {
	ID          int64     `validate:"required" json:"id"`
	Title       string    `validate:"required,min=3" json:"title"`
	Description string    `validate:"required,min=10" json:"description"`
	Tag         string    `validate:"required" json:"tag"`
	DueDate     time.Time `validate:"required" json:"due_date"`
	CreatedAt   time.Time `validate:"required" json:"created_at"`
	UpdatedAt   time.Time `validate:"required" json:"updated_at"`
}
