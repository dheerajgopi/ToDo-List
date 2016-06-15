package main

import (
	"time"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed string    `json:"completed"`
	Due       time.Time `json:due`
}

type Todos []Todo
