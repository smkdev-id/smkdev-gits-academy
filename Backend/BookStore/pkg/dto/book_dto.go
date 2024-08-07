package dto

import (
	"BookStore/pkg/models"
)

// BookRequest adalah struktur untuk request buku.
type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  uint   `json:"price"`
}

// BookResponse adalah struktur untuk response buku.
type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Price     uint   `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToBook mengubah BookRequest menjadi model Book.
func (t *BookRequest) ToBook() *models.Book {
	return &models.Book{
		Title:  t.Title,
		Author: t.Author,
		Price:  t.Price,
	}
}

// EntityToResponse mengubah model Book menjadi BookResponse.
func EntityToResponse(book *models.Book) *BookResponse {
	return &BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Price:     book.Price,
		CreatedAt: book.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: book.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
