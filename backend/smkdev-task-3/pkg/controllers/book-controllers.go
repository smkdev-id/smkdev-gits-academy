package controllers

import (
	"BookStore/pkg/models"
	"BookStore/pkg/service"
	"BookStore/pkg/utils"
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
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to retrieve books", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Books retrieved successfully", books)
}

func (bc *BookController) GetBooksByID(c *gin.Context) {
	var book models.BookRequest
	if err := c.ShouldBindUri(&book); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id, err := uuid.Parse(book.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid UUID format", nil)
		return
	}

	booksbyid, err := bc.bookService.GetBooksByID(c, id)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to retrieve books", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Book retrieved successfully", booksbyid)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.CreateBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdBook, err := bc.bookService.CreateBook(c, &book)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to create book", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Book created successfully", createdBook)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var book models.UpdateBookRequest

	if err := c.ShouldBindUri(&book); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	updatedBook, err := bc.bookService.UpdateBook(c, &book)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Book updated successfully", updatedBook)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	var book models.DeleteBookRequest
	if err := c.ShouldBindUri(&book); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id, err := uuid.Parse(book.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid UUID format", nil)
		return
	}

	if err := bc.bookService.DeleteBook(c, id); err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Book deleted successfully", nil)
}
