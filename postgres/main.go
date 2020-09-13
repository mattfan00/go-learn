package main

import (
	// "database/sql"
	"net/http"

	"api/database"
	"api/router"

	_ "github.com/lib/pq"
)

func main() {
	router := router.NewRouter()
	db := database.Init()
	defer db.Close()

	http.ListenAndServe(":8080", router)

	/*
	 */
}
