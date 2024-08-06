package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct{}

func NewBookController() *BookController {
	return &BookController{}
}

func (bc *BookController) GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
