package main

import (
	"encoding/json"
	"net/http"
)

type jsonErr struct {
	ErrorMessage string `json:"errorMessage"`
}

func SendTodos(w http.ResponseWriter, todos *Todos) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func SendTodo(w http.ResponseWriter, todo *Todo) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
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
