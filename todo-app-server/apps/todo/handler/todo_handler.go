package handler

import (
	"net/http"
)

type TodoHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FilterByTag(w http.ResponseWriter, r *http.Request)
	GetQuote(w http.ResponseWriter, r *http.Request)
}
