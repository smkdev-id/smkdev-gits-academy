package controllers

import (
	"github.com/hasban-fardani/go-api-book-store/pkg/database"
	models "github.com/hasban-fardani/go-api-book-store/pkg/model"
	"github.com/hasban-fardani/go-api-book-store/pkg/utils"
	"github.com/labstack/echo/v4"
)

type BookController struct{}

func NewBookController() *BookController {
	return &BookController{}
}

func (controller *BookController) Index(c echo.Context) error {
	db := database.OpenDB()

	defer database.CloseDB(db)

	books := []models.Book{}

	db.Find(&books)

	return c.JSON(200, books)
}

func (controller *BookController) Show(c echo.Context) error {
	id := c.Param("id")

	db := database.OpenDB()

	defer database.CloseDB(db)

	book := models.Book{}

	db.First(&book, id)

	return c.JSON(200, book)
}

func (controller *BookController) Create(c echo.Context) error {
	var bookRequest models.Book

	err := c.Bind(&bookRequest)

	utils.HandleError(err)

	db := database.OpenDB()

	defer database.CloseDB(db)

	db.Create(&bookRequest)

	return c.JSON(200, bookRequest)
}

func (controller *BookController) Update(c echo.Context) error {
	var bookRequest models.Book

	err := c.Bind(&bookRequest)

	utils.HandleError(err)

	bookID := c.Param("id")

	db := database.OpenDB()

	defer database.CloseDB(db)

	db.Model(&bookRequest).Where("id = ?", bookID).Updates(bookRequest)

	//
	bookRequest.ID = utils.StringToInt(bookID)

	return c.JSON(200, bookRequest)
}

func (controller *BookController) Delete(c echo.Context) error {
	bookID := c.Param("id")

	db := database.OpenDB()

	defer database.CloseDB(db)

	db.Delete(&models.Book{}, bookID)

	return c.JSON(200, nil)
}
