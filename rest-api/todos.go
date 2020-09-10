package main

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

var todos = Todos{
	{
		Id:   uuid.New(),
		Name: "finish homework",
	},
	{
		Id:   uuid.New(),
		Name: "do laundry",
	},
}
