package repository

import (
	"database/sql"
	"fmt"

	"github.com/sidz111/book-crud-go/internal/model"
)

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) Save(book model.Book) error {
	query := "Insert into Book(name, authorname, publishedyear) values(?,?,?)"
	_, err := r.DB.Exec(query, book.Name, book.AuthorName, book.PublishedYear)
	return err
}

func (r *BookRepository) Update(book model.Book) error {
	query := "update book set name=?, authorname=?, publishedyear=? where id=?"
	_, err := r.DB.Exec(query, book.Name, book.AuthorName, book.PublishedYear, book.ID)
	return err
	// var b model.Book
	// row.Scan(&book.Name, &book.AuthorName, &book.PublishedYear, &book.ID)
	// return b
}

func (r *BookRepository) DeleteById(id int) (string, error) {
	query := "delete from book where id =?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return "Book Not Found", err
	} else {
		return "Book Deleted", err
	}
}

func (r *BookRepository) GetAllBooks() ([]model.Book, error) {
	query := "select id, name, authorname, publishedyear from book"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, sql.ErrNoRows
	}
	var books []model.Book

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Name, &book.AuthorName, &book.PublishedYear)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	defer rows.Close()
	return books, err
}

func (r *BookRepository) GetBookById(id int) (model.Book, error) {
	query := "select id, name, authorname, publishedyear from book where id =?"
	row := r.DB.QueryRow(query, id)
	var book model.Book
	err := row.Scan(&book.ID, &book.Name, &book.AuthorName, &book.PublishedYear)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("Book Not Found")
		}
		return book, err
	}
	return book, nil
}
