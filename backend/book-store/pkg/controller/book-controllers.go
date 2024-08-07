package controller

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var logger = utils.InitializeLogger()

func GetBooks(c echo.Context) error {
	var books []models.Book
	result := config.DB.Find(&books)
	if result.Error != nil {
		logger.Println("Error:", result.Error)
		return utils.RespondWithError(c, http.StatusInternalServerError, "Error fetching books")
	}

	return utils.NewSuccessResponse(c, "Success get books", books, http.StatusOK)
}

func CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		logger.Info("Error:", err)
		return utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
	}

	result := config.DB.Create(book)
	if result.Error != nil {
		logger.Info("Error:", result.Error)
		return utils.RespondWithError(c, http.StatusInternalServerError, "Error creating book")
	}

	return utils.NewSuccessResponse(c, "Success create book", book, http.StatusCreated)
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("bookId"))
	var book models.Book

	result := config.DB.First(&book, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return utils.RespondWithError(c, http.StatusNotFound, "Book not found")
		}
		logger.Println("Error:", result.Error)
		return utils.RespondWithError(c, http.StatusInternalServerError, result.Error.Error())
	}
	return utils.NewSuccessResponse(c, "Success get book by id", book, http.StatusOK)
}

func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("bookId"))
	var book models.Book
	result := config.DB.Delete(&book, id)
	if result.Error != nil {
		logger.Println("Error:", result.Error)
		return utils.RespondWithError(c, http.StatusInternalServerError, "Error deleting book")
	}
	return utils.NewSuccessResponse(c, "Success delete book", nil, http.StatusOK)
}

func UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("bookId"))
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return utils.RespondWithError(c, http.StatusNotFound, "Book not found")
	}
	if err := c.Bind(&book); err != nil {
		logger.Println("Error:", err)
		return utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	config.DB.Save(&book)
	return utils.NewSuccessResponse(c, "Success update book", book, http.StatusOK)
}
