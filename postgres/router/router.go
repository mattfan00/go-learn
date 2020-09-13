package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path        string
	Methods     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Path(route.Path).
			Methods(route.Methods).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}

var routes = []Route{
	{"/", "GET", userIndex},
	{"/{userId}", "GET", userShow},
	{"/", "POST", userCreate},
}
