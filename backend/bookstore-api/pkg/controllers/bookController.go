package controllers

import (
	"bookstore-api/pkg/models"
	"bookstore-api/pkg/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{DB: db}
}

func (bc *BookController) GetAllBooks(c echo.Context) error {
	var books []models.Book

	if err := bc.DB.Find(&books).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.SendResponse(c, http.StatusOK, "Successfully fetch all books", books)
}

func (bc *BookController) GetBookByID(c echo.Context) error {
	var book []models.Book
	id, errid := utils.StringToInt(c.Param("id"))

	if errid != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Invalid ID", nil)
	}

	if err := bc.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "Not Found", nil)
		}
		return utils.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.SendResponse(c, http.StatusOK, "Successfully fetch book", book)

}

func (bc *BookController) CreateBook(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	book.CreatedAt = time.Now()

	if err := bc.DB.Create(&book).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.SendResponse(c, http.StatusOK, "Successfully create a new book", book)

}

func (bc *BookController) UpdateBook(c echo.Context) error {
	var book models.Book
	id, err := utils.StringToInt(c.Param("id"))

	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Invalid ID", nil)
	}

	if err := bc.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "Book Not Found", nil)
		}
		return utils.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	if err := c.Bind(&book); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	book.UpdatedAt = time.Now()

	if err := bc.DB.Save(&book).Error; err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return utils.SendResponse(c, http.StatusOK, "Successfully update book", book)
}

func (bc *BookController) DeleteBook(c echo.Context) error {
	var book models.Book
	id, err := utils.StringToInt(c.Param("id"))

	if err != nil {
		return utils.SendResponse(c, http.StatusNotFound, "Invalid ID", nil)
	}

	if err := bc.DB.First(&book, id).Delete(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "Book Not Found", nil)
		}
		return utils.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.SendResponse(c, http.StatusOK, "Successfully Delete Book", nil)
}
