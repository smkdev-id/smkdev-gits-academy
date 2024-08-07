package controllers

import (
	"net/http"

	"github.com/PorcoGalliard/BookStore/pkg/config"
	"github.com/PorcoGalliard/BookStore/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    config.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&book)
    c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
    var book models.Book
    id := c.Param("bookid")
    if err := config.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
    var book models.Book
    id := c.Param("bookid")
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

func DeleteBook(c *gin.Context) {
    var book models.Book
    id := c.Param("bookid")
    if err := config.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    config.DB.Delete(&book)
    c.JSON(http.StatusNoContent, nil)
}
