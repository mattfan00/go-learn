package models

type User struct {
	Id        int    `json:"id"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
