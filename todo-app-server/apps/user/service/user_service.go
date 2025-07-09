package service

import (
	"context"
	"todo-app/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.CreateRequest) web.LoginResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	VerifyToken(ctx context.Context) web.VerifyResponse
}
