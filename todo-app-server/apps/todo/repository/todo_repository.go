package repository

import (
	"context"
	"todo-app/model/web"
)

type TodoRepository interface {
	Create(ctx context.Context, req *web.CreateTodoRequest, userId int64) (*web.TodoResponse, error)
	Get(ctx context.Context, id int64) (*web.TodoResponse, error)
	GetAll(ctx context.Context, userId int64) ([]*web.TodoResponse, error)
	Update(ctx context.Context, req *web.CreateTodoRequest, id int64) (*web.CreateTodoRequest, error)
	Delete(ctx context.Context, id int64) error
	FilterByTag(ctx context.Context, search string, tag interface{}) ([]*web.TodoResponse, error)
}
