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

	newEncoder := json.NewEncoder(w)
	// newEncoder.SetIndent("", "    ")
	newEncoder.Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	todoId := mux.Vars(r)["todoId"]
	fmt.Fprintf(w, "showing todo of id "+todoId)
}
