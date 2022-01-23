package fun

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Books []Book

func Initial() {
	// Mock Data - @todo - implement DB
	Books = append(Books, Book{
		Id:   "1",
		Name: "how we should sex with our partneradsasdf",
	})
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var current Book
	for _, item := range Books {
		if item.Id == params["id"] {
			current = item
		}
	}

	json.NewEncoder(w).Encode(current)
}
