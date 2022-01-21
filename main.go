package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	fun "github.com/alirezakargar1380/go_basics/test"
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
	var x int = 34
	fmt.Printf("hello world = %v moda fucker", x)

	// if
	uio := 8
	if uio > 9 {
		fmt.Println("im bigger than 9")
	} else {
		fmt.Println("im smaller than 9")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("i is = %v my darlin\n", i)
	}

	// fmt.Println(some_fun())
	fmt.Println(fun.Some_fun())
	// fmt.Println(fun.abc)

	// init Router
	// asd
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{
		Id:   "1",
		Name: "how we should sex with our partneradsasdf",
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/getBooks", getBooks).Methods("GET")
	r.HandleFunc("/getBook/{id}", getBook).Methods("GET")
	http.ListenAndServe(":8000", r)
	// log.Fatal()
}
