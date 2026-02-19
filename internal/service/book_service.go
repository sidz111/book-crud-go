package service

import (
	"github.com/sidz111/book-crud-go/internal/model"
	"github.com/sidz111/book-crud-go/internal/repository"
)

type BookService struct {
	Repo *repository.BookRepository
}

func (s *BookService) Save(book model.Book) error {
	return s.Repo.Save(book)
}

func (s *BookService) Update(book model.Book) error {
	return s.Repo.Update(book)
}

func (s *BookService) DeleteBookById(id int) (string, error) {
	return s.Repo.DeleteById(id)
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.Repo.GetAllBooks()
}

func (s *BookService) GetBookById(id int) (model.Book, error) {
	return s.Repo.GetBookById(id)
}
