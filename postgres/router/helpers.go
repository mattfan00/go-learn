package router

import (
	"encoding/json"
	"net/http"
)

type jsonErr struct {
	ErrorMessage string `json:"errorMessage"`
}

func SendResponse(w http.ResponseWriter, info interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(info); err != nil {
		panic(err)
	}
}

func SendError(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	returnErr := jsonErr{errorMessage}
	if err := json.NewEncoder(w).Encode(returnErr); err != nil {
		panic(err)
	}
}
