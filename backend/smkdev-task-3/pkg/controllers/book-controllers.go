package controllers

import (
	"BookStore/pkg/models"
	"BookStore/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{bookService: bookService}
}
func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.bookService.GetBooks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	c.JSON(http.StatusOK, books)
}
func (bc *BookController) GetBooksByID(c *gin.Context) {
	var book models.BookRequest
	if err := c.ShouldBindUri(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if err := c.ShouldBindJSON(&book); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	id, err := uuid.Parse(book.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	booksbyid, err := bc.bookService.GetBooksByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	c.JSON(http.StatusOK, booksbyid)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.CreateBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := bc.bookService.CreateBook(c, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	c.JSON(http.StatusOK, createdBook)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var book models.UpdateBookRequest

	if err := c.ShouldBindUri(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook, err := bc.bookService.UpdateBook(c, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	var book models.DeleteBookRequest
	if err := c.ShouldBindUri(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(book.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	if err := bc.bookService.DeleteBook(c, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
