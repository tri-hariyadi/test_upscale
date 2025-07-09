//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
	todo_handler "todo-app/apps/todo/handler"
	todo_repository "todo-app/apps/todo/repository"
	todo_service "todo-app/apps/todo/service"
	user_handler "todo-app/apps/user/handler"
	user_repository "todo-app/apps/user/repository"
	user_service "todo-app/apps/user/service"
	"todo-app/common/database"
	"todo-app/common/router"
	"todo-app/common/server"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

var userSet = wire.NewSet(
	user_repository.NewUserRepositoryImpl,
	wire.Bind(new(user_repository.UserRepository), new(*user_repository.UserRepositoryImpl)),
	user_service.NewUserServiceImpl,
	wire.Bind(new(user_service.UserService), new(*user_service.UserServiceImpl)),
	user_handler.NewUserHandlerImpl,
)

var todoSet = wire.NewSet(
	todo_repository.NewTodoRepositoryImpl,
	wire.Bind(new(todo_repository.TodoRepository), new(*todo_repository.TodoRepositoryImpl)),
	todo_service.NewTodoServiceImpl,
	wire.Bind(new(todo_service.TodoService), new(*todo_service.TodoServiceImpl)),
	todo_handler.NewTodoHandlerImpl,
)

func InitializedServer() (*http.Server, func()) {
	wire.Build(
		database.ConnectDB,
		NewValidator,
		userSet,
		todoSet,
		router.NewRouter,
		server.NewServer,
	)
	return nil, nil
}
