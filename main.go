package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var current Book
	for _, item := range books {
		if item.Id == params["id"] {
			current = item
		}
	}

	json.NewEncoder(w).Encode(current)
}

func main() {
	fmt.Println("hello world")

	// init Router
	// asd
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{
		Id:   "1",
		Name: "how we should sex with our partner",
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/getBooks", getBooks).Methods("GET")
	r.HandleFunc("/getBook/{id}", getBook).Methods("GET")
	http.ListenAndServe(":8000", r)
	// log.Fatal()
}
