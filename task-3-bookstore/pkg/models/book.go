package models

import (
	"github.com/firyalhfz/bookstore-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//create a new book
func(b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//get all books
func GetAllBooks() ([]Book) {
    var books []Book
    db.Find(&books)
    return books
}


//get book by id
func GetBookById(id int64) (*Book, *gorm.DB ) {
    var book Book
    db := db.Where("ID=?", id).First(&book)
    return &book, db
}

//delete 
func DeleteBook(id int64) Book {
    var book Book
    db.Where("ID=?", id).Delete(&book)
    return book
}

