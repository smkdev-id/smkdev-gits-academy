package controllers

import (
    "net/http"
    "bookstore/pkg/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type BookController struct {
    DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
    return &BookController{DB: db}
}

func (bc *BookController) GetBooks(c *gin.Context) {
    var books []models.Book
    bc.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

func (bc *BookController) CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    bc.DB.Create(&book)
    c.JSON(http.StatusCreated, book)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := bc.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := bc.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    bc.DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := bc.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    bc.DB.Delete(&book)
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}