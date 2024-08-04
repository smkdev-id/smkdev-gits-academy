package controllers

import (
	"book-sales-management/config"
	"book-sales-management/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetBooks retrieves all books
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"message": "Data retrieved successfully", "data": books})
}

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// GetBookByID retrieves a book by its ID
func GetBookByID(c *gin.Context) {
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("bookId")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("bookId")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("bookId")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted", "data": book})
}
