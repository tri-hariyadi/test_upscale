package resp_handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"todo-app/common/logger"
	exception2 "todo-app/helpers/exception"
)

func Panic(err error, tag ...string) {
	if err != nil {
		message := ""
		if len(tag) > 0 {
			message = "\033[31m" + strings.Join(tag, " ") + ":\033[0m "
		}
		logger.Logger.Error(message + err.Error())
		panic(err)
	}
}

func ErrorHandler(w http.ResponseWriter, err error) {
	if notFoundError(w, err) {
		return
	}
	if validationError(w, err) {
		return
	}
	if badRequestError(w, err) {
		return
	}
	if unauthorizedError(w, err) {
		return
	}
	internalServerError(w)
}

func internalServerError(w http.ResponseWriter) {
	HandleError(w, http.StatusInternalServerError, "Internal Server Error")
}

func notFoundError(w http.ResponseWriter, err error) bool {
	var exception *exception2.NotFoundError
	if ok := errors.As(err, &exception); ok {
		HandleError(w, exception.StatusCode, exception.Error())
		return true
	}
	return false
}

func validationError(w http.ResponseWriter, err error) bool {
	var exception validator.ValidationErrors
	if ok := errors.As(err, &exception); ok {
		HandleError(w, http.StatusBadRequest, exception.Error())
		return true
	}
	return false
}

func badRequestError(w http.ResponseWriter, err error) bool {
	var exception *exception2.BadRequestError
	if ok := errors.As(err, &exception); ok {
		HandleError(w, exception.StatusCode, exception.Error())
		return true
	}
	return false
}

func unauthorizedError(w http.ResponseWriter, err error) bool {
	var exception *exception2.UnauthorizedError
	if ok := errors.As(err, &exception); ok {
		HandleError(w, exception.StatusCode, exception.Error())
		return true
	}
	return false
}
