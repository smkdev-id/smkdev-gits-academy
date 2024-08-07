package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name     string    `json:"name"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
}

var db *gorm.DB
var err error

func main() {
	e := echo.New()

	db, err = gorm.Open(sqlite.Open("perpustakaan.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&Book{})

	// Routes
	e.GET("/books", getAllBooks)
	e.POST("/books", addBook)
	e.GET("/books/:id", getBook)
	e.PUT("/books/:id", updateBook)
	e.DELETE("/books/:id", deleteBook)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func addBook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	book.Id = uuid.New()
	if err := db.Create(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create book"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book successfully created",
		"data":    book,
	})
}

func getAllBooks(c echo.Context) error {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not retrieve books"})
	}

	return c.JSON(http.StatusOK, books)
}

func getBook(c echo.Context) error {
	id := c.Param("id")
	book := new(Book)
	if err := db.First(&book, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	return c.JSON(http.StatusOK, book)
}

func updateBook(c echo.Context) error {
	id := c.Param("id")
	book := new(Book)
	if err := db.First(&book, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := db.Save(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update book"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book successfully updated",
		"data":    book,
	})
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	book := new(Book)
	if err := db.First(&book, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	if err := db.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete book"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Book successfully deleted"})
}
