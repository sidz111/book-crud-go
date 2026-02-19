package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sidz111/book-crud-go/internal/model"
	"github.com/sidz111/book-crud-go/internal/service"
)

type BookController struct {
	Service *service.BookService
}

func (c *BookController) Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusMethodNotAllowed)
		return
	}

	var book model.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	err = c.Service.Save(book)

	if err != nil {
		http.Error(w, "Failed to Save Book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book Saved Successfully"))

}

func (c *BookController) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid Method type", http.StatusMethodNotAllowed)
		return
	}

	var book model.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updatedBook := c.Service.Update(book)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}
