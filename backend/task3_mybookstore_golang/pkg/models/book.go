package models

import (
    "task3_mybookstore_golang/pkg/config"
    "gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
    gorm.Model
    Title       string `json:"title"`
    Author      string `json:"author"`
    Publisher   string `json:"publisher"`
}

func init() {
    config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
    db.Create(&b)
    return b
}

func GetAllBooks() []Book {
    var Books []Book
    db.Find(&Books)
    return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
    var getBook Book
    db := db.Where("ID=?", Id).Find(&getBook)
    return &getBook, db
}

func DeleteBook(ID int64) Book {
    var book Book
    db.Where("ID=?", ID).Delete(book)
    return book
}
