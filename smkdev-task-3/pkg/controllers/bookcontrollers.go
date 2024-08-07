package controllers

import (
    "net/http"
    "smkdev-task-3/pkg/config"
    "smkdev-task-3/pkg/models"
    "smkdev-task-3/pkg/utils"

    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
)

var validate = validator.New()

func GetBooks(c echo.Context) error {
    var books []models.Book
    result := config.DB.Find(&books)
    if result.Error != nil || len(books) == 0 {
        return utils.ResponseError(c, http.StatusNotFound, "Data Not Found")
    }
    return utils.ResponseSuccess(c, "Get books", utils.FormatBooks(books))
}

func CreateBook(c echo.Context) error {
    book := new(models.Book)
    if err := c.Bind(book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err.Error())
    }
    if err := validate.Struct(book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err.Error())
    }
    config.DB.Create(book)
    return utils.ResponseSuccess(c, "Success Create Book", utils.FormatBook(*book))
}

func GetBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book
    if result := config.DB.First(&book, id); result.Error != nil {
        return utils.ResponseError(c, http.StatusNotFound, "Data Not Found")
    }
    return utils.ResponseSuccess(c, "Get Book by Id", utils.FormatBook(book))
}

func UpdateBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book
    if result := config.DB.First(&book, id); result.Error != nil {
        return utils.ResponseError(c, http.StatusNotFound, "Data Not Found")
    }
    if err := c.Bind(&book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err.Error())
    }
    if err := validate.Struct(book); err != nil {
        return utils.ResponseError(c, http.StatusBadRequest, err.Error())
    }
    config.DB.Save(&book)
    return utils.ResponseSuccess(c, "Success Update Book", utils.FormatBook(book))
}

func DeleteBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book
    if result := config.DB.First(&book, id); result.Error != nil {
        return utils.ResponseError(c, http.StatusNotFound, "Data Not Found")
    }
    config.DB.Delete(&book, id)
    return utils.ResponseSuccess(c, "Success Delete Book", nil)
}
