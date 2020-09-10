package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellow owrld")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	SendTodos(w, &todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["todoId"]
	for _, todo := range todos {
		if todo.Id.String() == todoId {
			SendTodo(w, &todo)
			return
		}
	}

	SendError(w, "Cannot find todo with id "+todoId, http.StatusNotFound)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &newTodo); err != nil {
		SendError(w, "Invalid json syntax", http.StatusBadRequest)
		return
	}

	newTodo.Id = uuid.New()
	todos = append(todos, newTodo)
	SendTodo(w, &newTodo)
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo

	todoId := mux.Vars(r)["todoId"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &updatedTodo); err != nil {
		SendError(w, "Invalid json syntax", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.Id.String() == todoId {
			todos[i] = updatedTodo
			SendTodo(w, &updatedTodo)
			return
		}
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["todoId"]

	for i, todo := range todos {
		if todo.Id.String() == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			SendTodo(w, &todo)
			return
		}
	}

	SendError(w, "Cannot find todo with id "+todoId, http.StatusNotFound)
}
