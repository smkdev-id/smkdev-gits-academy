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

	return utils.SendResponse(c, http.StatusOK, "success", books)
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

	return utils.SendResponse(c, http.StatusOK, "success", book)

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

	return utils.SendResponse(c, http.StatusOK, "success", book)

}
