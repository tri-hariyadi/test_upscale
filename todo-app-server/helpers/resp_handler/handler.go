package resp_handler

import (
	"encoding/json"
	"net/http"
	"todo-app/common/logger"
	"todo-app/model/web"
)

func HandleSuccess(w http.ResponseWriter, code int, data interface{}) {
	response := &web.ResponseWrapper{
		Status:  true,
		Message: "success",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(response)
	Panic(err)
}

func HandleError(w http.ResponseWriter, code int, msg string) {
	response := &web.ResponseWrapper{
		Status:  false,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Logger.Info(err.Error())
	}
}

func HandleRequestBody(r *http.Request, data interface{}) {
	err := json.NewDecoder(r.Body).Decode(data)
	Panic(err, "[HandleRequestBody]")
}
