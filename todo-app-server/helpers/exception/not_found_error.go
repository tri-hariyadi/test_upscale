package exception

import "net/http"

type NotFoundError struct {
	StatusCode int
	Message    string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{StatusCode: http.StatusNotFound, Message: message}
}

func (n NotFoundError) Error() string {
	return n.Message
}

