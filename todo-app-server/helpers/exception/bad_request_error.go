package exception

import "net/http"

type BadRequestError struct {
	StatusCode int
	Message    string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{StatusCode: http.StatusBadRequest, Message: message}
}

func (b BadRequestError) Error() string {
	return b.Message
}

