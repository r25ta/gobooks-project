package main

import (
	"fmt"
	"gobooks/internal/service"
)

func main() {
	book := service.Book{
		ID:     1,
		Title:  "The Hobbit",
		Author: "J.J.R. Tolkien",
		Genre:  "Fantasy",
	}
	fmt.Println("The book is ", book.Title)
}
