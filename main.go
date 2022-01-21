package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello world")

	// init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	// r.HandleFunc("/", getBooks).Methods("GET")
	http.ListenAndServe(":8000", r)
}
