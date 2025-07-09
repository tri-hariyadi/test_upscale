package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"todo-app/common/config"
	"todo-app/common/logger"
	"todo-app/helpers/auth"
	"todo-app/helpers/exception"
	"todo-app/helpers/resp_handler"
)

func TokenVerifyMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var secret []byte = []byte(config.AppConfig.JWTSecret.AccessSecret)

			//cookie, err := r.Cookie("access_token")
			//if err != nil {
			//	logger.Logger.Error(err, cookie)
			//	resp_handler.Panic(exception.NewUnauthorizedError("You don't have access, token is invalid"), "[TokenVerifyMiddleware]")
			//}

			authHeader := r.Header.Get("Authorization")
			bearerToken := strings.Split(authHeader, " ")

			if len(bearerToken) != 2 {
				resp_handler.Panic(exception.NewUnauthorizedError("You don't have access, token is invalid"), "[TokenVerifyMiddleware]")
			}

			token, err := jwt.ParseWithClaims(cookie.Value, &auth.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("internal server error")
				}
				return secret, nil
			})
			if err != nil || !token.Valid {
				logger.Logger.Error(err, token.Valid)
				resp_handler.Panic(exception.NewUnauthorizedError("You don't have access, token is invalid"), "[TokenVerifyMiddleware]")
			}

			claims, ok := token.Claims.(*auth.CustomClaims)
			if !ok {
				resp_handler.Panic(errors.New("internal server error"), "[TokenVerifyMiddleware]")
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
