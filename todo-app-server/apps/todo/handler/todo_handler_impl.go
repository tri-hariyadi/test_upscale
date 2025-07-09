package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo-app/apps/todo/service"
	"todo-app/common/router"
	"todo-app/helpers/auth"
	"todo-app/helpers/resp_handler"
	"todo-app/middleware"
	"todo-app/model/web"
)

type TodoHandlerImpl struct {
	TodoService service.TodoService
}

func NewTodoHandlerImpl(todoService service.TodoService) router.TodoRegisterHandlerFunc {
	return func(router *mux.Router) {
		todoHandler := &TodoHandlerImpl{TodoService: todoService}
		authRouter := router.PathPrefix("/todos").Subrouter()
		authRouter.Use(middleware.TokenVerifyMiddleware())

		authRouter.HandleFunc("", todoHandler.Create).Methods("POST")
		authRouter.HandleFunc("", todoHandler.GetAll).Methods("GET")
		authRouter.HandleFunc("/{id}", todoHandler.Update).Methods("PATCH")
		authRouter.HandleFunc("/{id}", todoHandler.Delete).Methods("DELETE")
		authRouter.HandleFunc("/filter", todoHandler.FilterByTag).Methods("GET")
		authRouter.HandleFunc("/today/quotes", todoHandler.GetQuote).Methods("GET")
	}
}

func (t TodoHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	createRequest := &web.CreateTodoRequest{}
	claims, ok := r.Context().Value("user").(*auth.CustomClaims)
	if !ok {
		resp_handler.Panic(
			errors.New("error parsing context value to CustomClaims"),
			"UserServiceImpl.VerifyToken",
		)
	}
	resp_handler.HandleRequestBody(r, createRequest)

	result := t.TodoService.Create(r.Context(), createRequest, claims.UserID)
	resp_handler.HandleSuccess(w, http.StatusOK, result)
}

func (t TodoHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	result := t.TodoService.GetAll(r.Context())
	resp_handler.HandleSuccess(w, http.StatusOK, result)
}

func (t TodoHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	todoID, err := strconv.Atoi(todoId)
	resp_handler.Panic(err, "TodoHandlerImpl.Update")

	createRequest := &web.CreateTodoRequest{}
	resp_handler.HandleRequestBody(r, createRequest)

	result := t.TodoService.Update(r.Context(), createRequest, int64(todoID))
	resp_handler.HandleSuccess(w, http.StatusOK, result)
}

func (t TodoHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	todoID, err := strconv.Atoi(todoId)
	resp_handler.Panic(err, "TodoHandlerImpl.Delete")

	t.TodoService.Delete(r.Context(), int64(todoID))
	resp_handler.HandleSuccess(w, http.StatusOK, nil)
}

func (t TodoHandlerImpl) FilterByTag(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	tag := r.URL.Query().Get("tag")

	var tagParam interface{}
	if tag == "" {
		tagParam = nil
	} else {
		tagParam = tag
	}

	result := t.TodoService.FilterByTag(r.Context(), search, tagParam)
	resp_handler.HandleSuccess(w, http.StatusOK, result)
}

func (t TodoHandlerImpl) GetQuote(w http.ResponseWriter, r *http.Request) {
	result := t.TodoService.GetQuote(r.Context())
	resp_handler.HandleSuccess(w, http.StatusOK, result)
}




