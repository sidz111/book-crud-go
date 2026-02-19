package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	err = c.Service.Update(book)
	if err != nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(updatedBook)
}

func (c *BookController) DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Error(w, "Invalid Id", http.StatusInternalServerError)
	}
	msg, err := c.Service.DeleteBookById(id)
	if err != nil {
		http.Error(w, "Book Not Found", http.StatusInternalServerError)
	}
	defer r.Body.Close()
	json.NewEncoder(w).Encode(map[string]string{
		"message": msg,
	})
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	books, err := c.Service.GetAllBooks()
	if err != nil {
		http.Error(w, "Data Not Found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func (c *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Book Not Found", http.StatusInternalServerError)
		return
	}
	book, err := c.Service.GetBookById(id)
	if err != nil {
		http.Error(w, "Data Not Found", http.StatusNotFound)
		return
	}
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(book)
}
