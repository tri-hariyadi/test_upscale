package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"todo-app/apps/todo/repository"
	"todo-app/common/config"
	"todo-app/helpers/auth"
	"todo-app/helpers/resp_handler"
	"todo-app/model/web"
)

type Quote struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	Validate       *validator.Validate
}

func NewTodoServiceImpl(todoRepository repository.TodoRepository, validate *validator.Validate) *TodoServiceImpl {
	return &TodoServiceImpl{TodoRepository: todoRepository, Validate: validate}
}

func (t TodoServiceImpl) Create(ctx context.Context, request *web.CreateTodoRequest, userId int64) *web.TodoResponse {
	err := t.Validate.Struct(request)
	resp_handler.Panic(err, "[TodoServiceImpl.Create]")

	result, err := t.TodoRepository.Create(ctx, request, userId)
	resp_handler.Panic(err, "[TodoServiceImpl.Create]")

	return result
}

func (t TodoServiceImpl) Get(ctx context.Context, id int64) *web.TodoResponse {
	resp, err := t.TodoRepository.Get(ctx, id)
	resp_handler.Panic(err, "[TodoServiceImpl.Get]")

	return resp
}

func (t TodoServiceImpl) GetAll(ctx context.Context) []*web.TodoResponse {
	claims, ok := ctx.Value("user").(*auth.CustomClaims)
	if !ok {
		resp_handler.Panic(
			errors.New("error parsing context value to CustomClaims"),
			"TodoServiceImpl.GetAll",
		)
	}

	resp, err := t.TodoRepository.GetAll(ctx, claims.UserID)
	resp_handler.Panic(err, "[TodoServiceImpl.GetAll]")

	return resp
}

func (t TodoServiceImpl) Update(ctx context.Context, request *web.CreateTodoRequest, id int64) *web.CreateTodoRequest {
	err := t.Validate.Struct(request)
	resp_handler.Panic(err, "[TodoServiceImpl.Update]")

	result, err := t.TodoRepository.Update(ctx, request, id)
	resp_handler.Panic(err, "[TodoServiceImpl.Update]")

	return result
}

func (t TodoServiceImpl) Delete(ctx context.Context, id int64) {
	err := t.TodoRepository.Delete(ctx, id)
	resp_handler.Panic(err, "[TodoServiceImpl.Delete]")
}

func (t TodoServiceImpl) FilterByTag(ctx context.Context, search string, tag interface{}) []*web.TodoResponse {
	result, err := t.TodoRepository.FilterByTag(ctx, search, tag)
	resp_handler.Panic(err, "[TodoServiceImpl.FilterByTag]")
	return result
}

func (t TodoServiceImpl) GetQuote(ctx context.Context) Quote {
	req, err := http.NewRequest("GET", config.AppConfig.Quotes.ApiUrl, nil)
	resp_handler.Panic(err, "[TodoServiceImpl.GetQuote] Failed to create request")

	req.Header.Set("X-Api-Key", config.AppConfig.Quotes.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	resp_handler.Panic(err, "[TodoServiceImpl.GetQuote] Request failed")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		resp_handler.Panic(err, "[TodoServiceImpl.GetQuote]")
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		resp_handler.Panic(err, "[TodoServiceImpl.GetQuote] API error")
	}

	var quotes []Quote
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		resp_handler.Panic(err, "[TodoServiceImpl.GetQuote] Decode error")
	}

	return quotes[0]
}

