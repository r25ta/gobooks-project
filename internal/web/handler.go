package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBoolHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}

}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "failed to get books", http.StatusInternalServerError)
		return
	}

	if len(books) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	//STRUCT
	var book service.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return

	}

	if err := h.service.CreateBook(&book); err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}
