package controllers

import (
	"books_store/PKG/config"
	"books_store/PKG/models"
	"books_store/PKG/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return utils.ResponseError(c, http.StatusBadRequest, "Invalid input")
	}

	book.ID = uuid.New()
	if err := config.DB.Create(&book).Error; err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Could not create book")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book successfully created",
		"data":    book,
	})
}

func GetBooks(c echo.Context) error {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Could not retrieve books")
	}

	return c.JSON(http.StatusOK, books)
}

func GetBookById(c echo.Context) error {
	id := c.Param("id")
	book := new(models.Book)
	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		return utils.ResponseError(c, http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	book := new(models.Book)
	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		return utils.ResponseError(c, http.StatusNotFound, "Book not found")
	}

	if err := c.Bind(book); err != nil {
		return utils.ResponseError(c, http.StatusBadRequest, "Invalid input")
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Could not update book")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book successfully updated",
		"data":    book,
	})
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	book := new(models.Book)
	if err := config.DB.First(&book, "id = ?", id).Error; err != nil {
		return utils.ResponseError(c, http.StatusNotFound, "Book not found")
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Could not delete book")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Book successfully deleted"})
}
