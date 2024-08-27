package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	//	db, err := sql.Open("sqlite3", "./books.db")
	db, err := sql.Open("postgres", "postgresql://postgres:admin@my_postgres:5432/sandboxdb?sslmode=disable")

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		panic(err)
	}

	defer db.Close()

	bookService := service.NewBookService(db)
	bookHandlers := web.NewBoolHandlers(bookService)

	//Criando servidor WEB
	router := http.NewServeMux()

	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	// Iniciando o servidor
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
