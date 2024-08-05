package controllers

import (
	"net/http"
	"time"

	"github.com/FuadGrimaldi/bookstore-api/pkg/models"
	"github.com/FuadGrimaldi/bookstore-api/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{DB: db}
}

/*
desc 	: Controller function to retrieve all books from the database.
method 	: GET
*/
func (bc *BookController) GetAllBooks(c echo.Context) error{
	var books []models.Book
	if err := bc.DB.Find(&books).Error; err != nil {
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "Successfully get a book", books)
}
/*
desc 	: Controller function to retrieve a single book by its ID.
method 	: GET
*/
func (bc *BookController) GetOneBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	// Find the book by ID
	if err := bc.DB.First(&book, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, http.StatusNotFound, "Book not found", nil)
		}
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "Succesfully get a book", book)
}
/*
desc 	: Controller function to create a new book entry in the database.
method 	: POST
*/
func (bc *BookController) CreateBook(c echo.Context) error {
	var book models.Book
	book.CreatedAt = time.Now()
	if err := c.Bind(&book); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	book.ID = uuid.New().String() // Generate a new unique ID for the book
	if err := bc.DB.Create(&book).Error; err != nil {
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.JSONResponse(c, http.StatusCreated, "Successfully created", book)
}
/*
desc 	: Controller function to update an existing book entry by its ID.
method 	: PUT
*/
func (bc *BookController) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	book.UpdatedAt = time.Now()
	// Find the book by ID
	if err := bc.DB.First(&book, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, http.StatusNotFound, "Book not found", nil)
		}
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	// Bind the new data to the book object
	if err := c.Bind(&book); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	// Save the updated book
	if err := bc.DB.Save(&book).Error; err != nil {
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "Successfully updated the book", book)
}
/*
desc 	: Controller function to delete a book entry by its ID.
method 	: DELETE
*/
func (bc *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	// Find the book by ID
	if err := bc.DB.First(&book, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.JSONResponse(c, http.StatusNotFound, "Book not found", nil)
		}
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	// Delete the book
	if err := bc.DB.Delete(&book).Error; err != nil {
		return utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "Successfully deleted the book", nil)
}
