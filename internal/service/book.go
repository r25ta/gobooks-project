package service

import (
	"database/sql"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre) VALUES(?,?,?)"

	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)

	if err != nil {
		return err
	}

	lastInserId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	book.ID = int(lastInserId)

	return nil

}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "SELECT id,title, author, genre FROM books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*Book, error) {
	query := "SELECT  id, title, author, genre FROM books WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) UpdateBook(book Book) error {
	query := "UPDATE book SET title=?, author=?, genre=? WHERE id=?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)

	/*CASO ERR = NIL UPDATE OK SENÃO UPDATE NOK*/
	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM book WHERE id = ?"

	_, err := s.db.Exec(query, id)

	/*CASO ERR = NIL DELETE OK SENÃO DELETE NOK*/
	return err

}

func (s *BookService) SearchBooksByName(name string) ([]Book, error) {
	query := "SELECT id, title, author, genre FROM books WHERE title LIKE ?"
	rows, err := s.db.Query(query, "%"+name+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}