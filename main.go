package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Todo creates a basic model
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

// Todos is a slice (an ordered collection) of the Todo model
type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index serves the root of the site
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex serves the todos route
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

// TodoShow serves the todoID on the todos route
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoID := vars["todoID"]

	fmt.Fprintln(w, "Todo show:", todoID)
}
