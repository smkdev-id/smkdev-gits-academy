package models

import (
	"errors"
	"book_sell/config"
	"gorm.io/gorm"
)

// Book adalah model data untuk buku
type Book struct {
	gorm.Model
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// GetAllBooks mengembalikan semua buku dari database
func GetAllBooks() []Book {
	db := config.ConnectDatabase()
	var books []Book
	db.Find(&books)
	return books
}

// AddBook menambahkan buku baru ke database
func AddBook(book Book) Book {
	db := config.ConnectDatabase()
	db.Create(&book)
	return book
}

// GetBookByID mengembalikan buku berdasarkan ID
func GetBookByID(id string) (Book, bool) {
	db := config.ConnectDatabase()
	var book Book
	if err := db.First(&book, "id = ?", id).Error; err != nil {
		return Book{}, false
	}
	return book, true
}

// UpdateBook memperbarui buku berdasarkan ID
func UpdateBook(id string, updatedData Book) (Book, error) {
	db := config.ConnectDatabase()
	var book Book
	if err := db.First(&book, "id = ?", id).Error; err != nil {
		return Book{}, errors.New("book not found")
	}

	book.Title = updatedData.Title
	book.Author = updatedData.Author
	book.Price = updatedData.Price

	db.Save(&book)
	return book, nil
}

// DeleteBook menghapus buku berdasarkan ID
func DeleteBook(id string) error {
	db := config.ConnectDatabase()
	var book Book
	if err := db.First(&book, "id = ?", id).Error; err != nil {
		return errors.New("book not found")
	}
	db.Delete(&book)
	return nil
}
