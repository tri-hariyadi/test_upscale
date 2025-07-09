package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"todo-app/model/web"
)

type UserServiceMock struct {
	mock.Mock
}

func (u *UserServiceMock) Register(ctx context.Context, request web.CreateRequest) web.LoginResponse {
	args := u.Called(ctx, request)
	return args.Get(0).(web.LoginResponse)
}

func (u *UserServiceMock) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	args := u.Called(ctx, request)
	return args.Get(0).(web.LoginResponse)
}

func (u *UserServiceMock) VerifyToken(ctx context.Context) web.VerifyResponse {
	args := u.Called(ctx)
	return args.Get(0).(web.VerifyResponse)
}

