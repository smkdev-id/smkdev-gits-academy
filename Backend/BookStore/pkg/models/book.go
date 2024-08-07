package models

import (
	"gorm.io/gorm"
)

// Book adalah model untuk tabel buku dalam database.
type Book struct {
	gorm.Model
	Title  string `json:"title" validate:"required,min=3"` // Judul buku yang harus diisi dan minimal 3 karakter
	Author string `json:"author" validate:"required,min=3"` // Penulis buku yang harus diisi dan minimal 3 karakter
	Price  uint   `json:"price" validate:"required,gte=1"`  // Harga buku yang harus diisi dan minimal 1
}
