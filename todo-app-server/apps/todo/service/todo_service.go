package service

import (
	"context"
	"todo-app/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request *web.CreateTodoRequest, userId int64) *web.TodoResponse
	Get(ctx context.Context, id int64) *web.TodoResponse
	GetAll(ctx context.Context) []*web.TodoResponse
	Update(ctx context.Context, request *web.CreateTodoRequest, id int64) *web.CreateTodoRequest
	Delete(ctx context.Context, id int64)
	FilterByTag(ctx context.Context, search string, tag interface{}) []*web.TodoResponse
	GetQuote(ctx context.Context) Quote
}
