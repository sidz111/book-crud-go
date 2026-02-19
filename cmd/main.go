package main

import (
	"fmt"
	"net/http"

	"github.com/sidz111/book-crud-go/config"
	"github.com/sidz111/book-crud-go/internal/controller"
	"github.com/sidz111/book-crud-go/internal/repository"
	"github.com/sidz111/book-crud-go/internal/service"
)

func main() {
	fmt.Println("Started...")
	db := config.ConnectDB()

	repo := &repository.BookRepository{DB: db}
	service := &service.BookService{Repo: repo}
	controller := &controller.BookController{Service: service}
	mux := http.NewServeMux()

	mux.HandleFunc("/book", controller.Save)
	mux.HandleFunc("/book/update", controller.Update)
	mux.HandleFunc("/book/delete", controller.DeleteById)
	mux.HandleFunc("/books", controller.GetAllBooks)
	mux.HandleFunc("/book/get", controller.GetBookById)

	fmt.Println("Server started at Port 8080")
	http.ListenAndServe(":8080", mux)
}
