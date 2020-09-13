package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"api/database"
	"api/models"

	"github.com/gorilla/mux"
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
		rows.Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		users = append(users, user)
	}

	SendResponse(w, users)
}

func userShow(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	query := "SELECT * FROM users WHERE id = $1"

	var user models.User
	err := database.Db.QueryRow(query, userId).Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		panic(err)
	}

	SendResponse(w, user)
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		SendError(w, "Invalid json format", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		SendError(w, "Must provide an email", http.StatusBadRequest)
		return
	}

	query := `
			INSERT INTO users (age, first_name, last_name, email)
			VALUES ($1, $2, $3, $4)
			RETURNING * 
		`

	err = database.Db.QueryRow(query, user.Age, user.FirstName, user.LastName, user.Email).
		Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendResponse(w, user)
}
