package controllers

import (
	"book_sell/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBooks mengembalikan semua buku
func GetBooks(c *gin.Context) {
	books := models.GetAllBooks()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook menambahkan buku baru
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book = models.AddBook(book) 
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GetBookByID mengembalikan buku berdasarkan ID
func GetBookByID(c *gin.Context) {
	id := c.Param("bookId")
	book, found := models.GetBookByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook memperbarui buku berdasarkan ID
func UpdateBook(c *gin.Context) {
	id := c.Param("bookId")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook, err := models.UpdateBook(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedBook})
}

// DeleteBook menghapus buku berdasarkan ID
func DeleteBook(c *gin.Context) {
	id := c.Param("bookId")
	if err := models.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
