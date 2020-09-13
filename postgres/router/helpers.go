package router

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, info interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(info); err != nil {
		panic(err)
	}
}
