package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellow owrld")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		{
			Name: "finish homework",
		},
		{
			Name: "do laundry",
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	todoId := mux.Vars(r)["todoId"]
	fmt.Fprintf(w, "showing todo of id "+todoId)
}
