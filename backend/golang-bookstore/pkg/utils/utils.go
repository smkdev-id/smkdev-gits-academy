package utils

import (
	"github.com/gin-gonic/gin"
	"golang-bookstore/pkg/config"
	"golang-bookstore/pkg/models"
)

func CreateBook(book *models.Book) (*models.Book, error) {
	result := config.DB.Create(book)
	return book, result.Error
}

func GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	result := config.DB.First(&book, id)
	return &book, result.Error
}

func UpdateBook(book *models.Book) (*models.Book, error) {
	result := config.DB.Save(book)
	return book, result.Error
}

func DeleteBook(id uint) error {
	result := config.DB.Delete(&models.Book{}, id)
	return result.Error
}

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := config.DB.Find(&books)
	return books, result.Error
}

func RespondSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.Response{
		Meta: models.Meta{
			Code:    code,
			Status:  "success",
			Message: message,
		},
		Data: data,
	})
}

func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, models.Response{
		Meta: models.Meta{
			Code:    code,
			Status:  "error",
			Message: message,
		},
	})
}
