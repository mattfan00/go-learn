package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "random-password"
	dbname   = "test_demo"
)

var Db *sql.DB

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("succesffuly connected")

	Db = database

	return Db
}
