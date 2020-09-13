package router

import (
	"fmt"
	"net/http"

	"api/database"
	"api/models"
)

func userIndex(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM users`
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Age, &user.Email, user.FirstName, user.LastName)
		users = append(users, user)
	}

	fmt.Println(users)

}

func userCreate(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `
			INSERT INTO users (age, email, first_name, last_name)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`
	id := 0
	err := database.Db.QueryRow(sqlStatement, 15, "noturemail@gmail.com", "John", "Smith").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("new record id is: ", id)
}
