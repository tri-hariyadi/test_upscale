package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo-app/helpers/exception"
	"todo-app/middleware"
)

type RegisterHandlerFunc func(r *mux.Router)
type UserRegisterHandlerFunc RegisterHandlerFunc
type TodoRegisterHandlerFunc RegisterHandlerFunc

func NewRouter(
	userHandler UserRegisterHandlerFunc,
	todoHandler TodoRegisterHandlerFunc,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	r := router.PathPrefix("/api/v1").Subrouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoverMiddleware)
	r.NotFoundHandler = middleware.RecoverMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(exception.NewNotFoundError("Page is not found"))
	}))

	userHandler(r)
	todoHandler(r)

	return r
}
