package controllers

import (
	"net/http"
	"task3/PKG/config"
	"task3/PKG/models"

	// "net/http"

	"github.com/gin-gonic/gin"
)

// create data book
func CREATEBOOK(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

// get all data (book)
func GETBOOK(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// get book by id
func GETBOOKBYID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// update book using id
func UPDATEBOOK(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
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

// delete book using id
func DELETEBOOK(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
