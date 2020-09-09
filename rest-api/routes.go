package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Path(route.Path).
			Methods(route.Method).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}

type Route struct {
	Name        string
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{
		"Index",
		"/",
		"GET",
		Index,
	},
	{
		"TodoIndex",
		"/todos",
		"GET",
		TodoIndex,
	},
	{
		"TodoShow",
		"/todos/{todoId}",
		"GET",
		TodoShow,
	},
}
