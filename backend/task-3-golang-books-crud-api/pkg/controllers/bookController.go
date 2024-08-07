package controllers

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	var books []models.Book
	result := config.DB.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	if len(books) == 0 {
		// Tambahkan logging jika tidak ada data
		c.Logger().Info("No books found in the database.")
	}
	return c.JSON(http.StatusOK, books)
}
func CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return err
	}
	config.DB.Create(book)
	return c.JSON(http.StatusCreated, book)
}

func GetBookByID(c echo.Context) error {
	id := c.Param("bookid")
	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, book)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("bookid")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(&book); err != nil {
		return err
	}
	config.DB.Save(&book)
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("bookid")
	if err := config.DB.Delete(&models.Book{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.NoContent(http.StatusNoContent)
}
