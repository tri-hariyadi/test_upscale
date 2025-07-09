package repository

import (
	"context"
	"todo-app/model"
	"todo-app/model/web"
)

type UserRepository interface {
	Create(ctx context.Context, data web.CreateRequest) (int64, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
