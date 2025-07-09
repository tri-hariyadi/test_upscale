package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"todo-app/apps/user/repository"
	"todo-app/helpers/auth"
	"todo-app/helpers/exception"
	"todo-app/helpers/resp_handler"
	"todo-app/model"
	"todo-app/model/web"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
	TokenService   auth.TokenService
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, Validate: validate, TokenService: auth.NewTokenAuthService()}
}

func (u UserServiceImpl) Register(ctx context.Context, request web.CreateRequest) web.LoginResponse {
	err := u.Validate.Struct(request)
	resp_handler.Panic(err, "[UserServiceImpl.Register]")

	password, err := auth.HashPassword(request.Password)
	resp_handler.Panic(err, "[UserServiceImpl.Register]")

	dataRequest := web.CreateRequest{
		Email:    request.Email,
		Password: password,
		Name:     request.Name,
	}

	id, err := u.UserRepository.Create(ctx, dataRequest)
	resp_handler.Panic(err, "[UserServiceImpl.Register]")

	token, err := u.TokenService.GenerateAccessToken(model.User{
		ID:    id,
		Email: request.Email,
		Name:  request.Name,
	})
	resp_handler.Panic(err, "[UserServiceImpl.Register]")

	return web.LoginResponse{
		AccessToken: token,
	}
}

func (u UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := u.Validate.Struct(request)
	resp_handler.Panic(err, "UserServiceImpl.Login")

	dataUser, err := u.UserRepository.FindByEmail(ctx, request.Email)
	resp_handler.Panic(err, "UserServiceImpl.Login")

	if !auth.VerifyPassword(request.Password, dataUser.Password) {
		resp_handler.Panic(exception.NewBadRequestError("Email or password is invalid"), "UserServiceImpl.Login")
	}

	token, err := u.TokenService.GenerateAccessToken(model.User{
		ID:    dataUser.ID,
		Email: dataUser.Email,
		Name:  dataUser.Name,
	})
	resp_handler.Panic(err, "UserServiceImpl.Login")

	return web.LoginResponse{
		AccessToken: token,
	}
}

func (u UserServiceImpl) VerifyToken(ctx context.Context) web.VerifyResponse {
	claims, ok := ctx.Value("user").(*auth.CustomClaims)
	if !ok {
		resp_handler.Panic(
			errors.New("error parsing context value to CustomClaims"),
			"UserServiceImpl.VerifyToken",
		)
	}

	return web.VerifyResponse{
		ID: claims.UserID,
		Name: claims.Name,
		Email: claims.Email,
	}
}
