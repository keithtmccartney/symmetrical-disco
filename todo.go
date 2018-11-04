package main

import "time"

// Todo creates a basic model
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is a slice (an ordered collection) of the Todo model
type Todos []Todo
