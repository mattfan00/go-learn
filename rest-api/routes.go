package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Logger(fn http.HandlerFunc, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s\t%s\t%s\t%s\t", r.Method, r.RequestURI, name, time.Since(start))

		fn(w, r)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)

		router.
			Path(route.Path).
			Methods(route.Method).
			Name(route.Name).
			HandlerFunc(handler)
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
	{
		"TodoCreate",
		"/todos",
		"POST",
		TodoCreate,
	},
	{
		"TodoUpdate",
		"/todos/{todoId}",
		"PUT",
		TodoUpdate,
	},
}
