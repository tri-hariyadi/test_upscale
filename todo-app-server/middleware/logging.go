package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"todo-app/common/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.String(),
			"remote": r.RemoteAddr,
		}).Info("\033[32mIncoming request\033[0m")

		next.ServeHTTP(w, r)

		logger.Logger.WithFields(logrus.Fields{
			"method":     r.Method,
			"url":        r.URL.String(),
			"remote":     r.RemoteAddr,
			"duration":   time.Since(start),
		}).Info("Completed request")
	})
}
