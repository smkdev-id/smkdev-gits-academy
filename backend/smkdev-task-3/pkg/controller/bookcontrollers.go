package controllers

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

func GetBooks(c echo.Context) error {
    var books []models.Book
    config.DB.Find(&books)
    return utils.ResponseSuccess(c, books)
}

func CreateBook(c echo.Context) error {
    book := new(models.Book)
    if err := c.Bind(book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err)
    }

    if err := validate.Struct(book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err)
    }

    config.DB.Create(book)
    return utils.ResponseSuccess(c, book)
}

func GetBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book
    if result := config.DB.First(&book, id); result.Error != nil {
        return utils.ResponseError(c, http.StatusNotFound, result.Error)
    }
    return utils.ResponseSuccess(c, book)
}

func UpdateBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book
    if result := config.DB.First(&book, id); result.Error != nil {
        return utils.ResponseError(c, http.StatusNotFound, result.Error)
    }

    if err := c.Bind(&book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err)
    }

    if err := validate.Struct(book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err)
    }

    config.DB.Save(&book)
    return utils.ResponseSuccess(c, book)
}

func DeleteBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book
    if result := config.DB.Delete(&book, id); result.Error != nil {
        return utils.ResponseError(c, http.StatusNotFound, result.Error)
    }
    return utils.ResponseNoContent(c)
}
