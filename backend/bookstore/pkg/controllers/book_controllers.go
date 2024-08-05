package controllers

import (
	"bookstore/pkg/helper"
	"bookstore/pkg/services"
	"net/http"

	"github.com/labstack/echo"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{service}
}

func (s *BookController) GetAllBooks(c echo.Context) error {
	books, err := s.service.GetAllBooks()

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to get all book", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to get all book", books)
	c.JSON(http.StatusOK, response)
	return nil
}

func (s *BookController) GetBookByID(c echo.Context) error {
	id := c.Param("id")

	book, err := s.service.GetBookByID(id)

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to get book", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to get book", book)
	c.JSON(http.StatusOK, response)
	return nil
}

func (s *BookController) CreateBook(c echo.Context) error {
	var request helper.Request

	err := c.Bind(&request)

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to parse JSON data", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	book, err := s.service.CreateBook(request)
	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to create book", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to create book", book)
	c.JSON(http.StatusOK, response)
	return nil
}

func (s *BookController) UpdateBook(c echo.Context) error {
	var request helper.Request

	err := c.Bind(&request)
	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to parse JSON data", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	id := c.Param("id")

	book, err := s.service.UpdateBook(id, request)

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to update book", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to update book", book)
	c.JSON(http.StatusOK, response)
	return nil
}

func (s *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	err := s.service.DeleteBook(id)
	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to delete book", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to delete book", nil)
	c.JSON(http.StatusOK, response)
	return nil
}
