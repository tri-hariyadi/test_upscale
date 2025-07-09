package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"todo-app/apps/user/repository"
	"todo-app/common/logger"
	"todo-app/helpers/auth"
	"todo-app/helpers/exception"
	"todo-app/model"
	"todo-app/model/web"
)

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateAccessToken(user model.User) (string, error) {
	return "mock-access-token", nil
}

var userRepositoryMock = &repository.MockUserRepository{Mock: mock.Mock{}}
var mockToken = new(MockTokenService)
var userService = UserServiceImpl{
	UserRepository: userRepositoryMock,
	Validate:       validator.New(),
	TokenService:   mockToken,
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	logger.InitLogger()
	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestLogin(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		loginReq := web.LoginRequest{
			Email:    "test@test.com",
			Password: "secret123",
		}

		hash, _ := auth.HashPassword("secret123")

		mockUser := model.User{
			Password: hash,
		}

		userRepositoryMock.On("FindByEmail", mock.Anything, loginReq.Email).Return(&mockUser, nil)

		resp := userService.Login(context.Background(), loginReq)
		assert.Equal(t, "mock-access-token", resp.AccessToken)
	})

	t.Run("Negative Case Invalid Password", func(t *testing.T) {
		loginReq := web.LoginRequest{
			Email:    "test@test.com",
			Password: "secret1232345",
		}

		hash, _ := auth.HashPassword("secret123")

		mockUser := model.User{
			Password: hash,
		}

		userRepositoryMock.On("FindByEmail", mock.Anything, loginReq.Email).Return(&mockUser, nil)
		defer func() {
			if r := recover(); r != nil {
				assert.Contains(t, fmt.Sprint(r), "Email or password is invalid")
			} else {
				t.Errorf("Expected panic due to invalid password, but no panic occurred")
			}
		}()

		userService.Login(context.Background(), loginReq)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		request := web.CreateRequest{
			Email:    "test@test.com",
			Password: "secret123",
			Name:     "Test User",
		}

		var mockId int64 = 1

		userRepositoryMock.On("Create", mock.Anything, mock.MatchedBy(func(req web.CreateRequest) bool {
			return req.Email == request.Email &&
				req.Name == request.Name &&
				req.Password != request.Password
		})).Return(mockId, nil)

		resp := userService.Register(context.Background(), request)
		assert.Equal(t, "mock-access-token", resp.AccessToken)
	})

	t.Run("Negative Case Validation Error", func(t *testing.T) {
		invalidRequest := web.CreateRequest{
			Email:    "",
			Password: "secret123",
			Name:     "Test User",
		}

		var mockId int64 = 1

		userRepositoryMock.On("Create", mock.Anything, mock.Anything).Return(mockId, nil)
		defer func() {
			if r := recover(); r != nil {
				var validationErrors validator.ValidationErrors
				ok := errors.As(r.(error), &validationErrors)
				assert.True(t, ok, "panic must be validator.ValidationErrors")
			} else {
				t.Error("did not panic as expected")
			}
		}()

		userService.Register(context.Background(), invalidRequest)
	})

	t.Run("Negative Case - Repository Create Failed", func(t *testing.T) {
		validRequest := web.CreateRequest{
			Email:    "test@test.com",
			Password: "secret123",
			Name:     "Test User",
		}

		userRepositoryMock.On("Create", mock.Anything, mock.MatchedBy(func(r web.CreateRequest) bool {
			return r.Email == validRequest.Email
		})).Return(int64(0), exception.NewBadRequestError("DB error"))

		defer func() {
			if r := recover(); r != nil {
				var validationErrors *exception.BadRequestError
				ok := errors.As(r.(error), &validationErrors)
				assert.True(t, ok, "panic must be bad request")
			} else {
				t.Error("did not panic as expected")
			}
		}()
		userService.Register(context.Background(), validRequest)
	})
}
