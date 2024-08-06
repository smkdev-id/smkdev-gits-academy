package controllers

import (
	"github.com/labstack/echo"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/database"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/models"
)

type BookController struct{}

func NewBookController() *BookController {
	return &BookController{}
}

func (controller *BookController) Index(c echo.Context) error {
	db := database.ConnectToSQLite()
	defer database.CloseDB(db)

	books := []models.Book{}

	db.Find(&books)

	return c.JSON(200, books)
}

func (controller *BookController) Show(c echo.Context) error {
	db := database.ConnectToSQLite()
	defer database.CloseDB(db)

	id := c.Param("id")
	book := models.Book{}
	db.Find(&book, id)

	return c.JSON(200, book)
}

func (controller *BookController) Create(c echo.Context) error {
	db := database.ConnectToSQLite()
	defer database.CloseDB(db)

	book := new(models.Book)
	c.Bind(book)
	db.Create(&book)

	return c.JSON(200, book)
}

func (controller *BookController) Update(c echo.Context) error {
	db := database.ConnectToSQLite()
	defer database.CloseDB(db)

	id := c.Param("id")
	book := new(models.Book)
	c.Bind(book)
	db.Model(&book).Where("id = ?", id).Updates(&book)

	return c.JSON(200, book)
}

func (controller *BookController) Delete(c echo.Context) error {
	db := database.ConnectToSQLite()
	defer database.CloseDB(db)

	id := c.Param("id")
	db.Delete(&models.Book{}, id)

	return c.NoContent(204)
}
