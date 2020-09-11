package main

import (
	"database/sql"
	"net/http"

	"api/database"
	"api/router"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	router := router.NewRouter()
	db = database.Init()
	defer db.Close()

	http.ListenAndServe(":8080", router)

	/*
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("succesffuly connected")

		sqlStatement := `
			INSERT INTO users (age, email, first_name, last_name)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`
		id := 0
		err = db.QueryRow(sqlStatement, 20, "mrf441@nyu.edu", "Matthew", "Fan").Scan(&id)
		if err != nil {
			panic(err)
		}

		fmt.Println("new record id is: ", id)
	*/
}
