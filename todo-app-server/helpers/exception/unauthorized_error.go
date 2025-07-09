package exception

import "net/http"

type UnauthorizedError struct {
	StatusCode int
	Message    string
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{StatusCode: http.StatusUnauthorized, Message: message}
}

func (u UnauthorizedError) Error() string {
	return u.Message
}

