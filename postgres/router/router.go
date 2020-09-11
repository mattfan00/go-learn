package router

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Path("/").
		Methods("GET").
		HandlerFunc(userIndex)

	return router
}
