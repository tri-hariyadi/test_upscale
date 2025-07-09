package middleware

import (
	"net/http"
	"todo-app/helpers/resp_handler"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recover from panic
		defer func() {
			if err := recover(); err != nil {
				resp_handler.ErrorHandler(w, err.(error))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
