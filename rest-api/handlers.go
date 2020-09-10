package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var todos = Todos{
	{
		Id:   uuid.New(),
		Name: "finish homework",
	},
	{
		Id:   uuid.New(),
		Name: "do laundry",
	},
}

type jsonErr struct {
	ErrorMessage string `json:"errorMessage"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellow owrld")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["todoId"]
	for _, todo := range todos {
		if todo.Id.String() == todoId {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(todo); err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	errorMessage := jsonErr{"Cannot find todo with id " + todoId}

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &newTodo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)

		errorMessage := jsonErr{"Invalid json syntax"}
		if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
			panic(err)
		}
		return
	}
	newTodo.Id = uuid.New()

	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(newTodo); err != nil {
		panic(err)
	}
}
