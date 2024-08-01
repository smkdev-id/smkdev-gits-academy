package controllers

import (
	"net/http"

	"github.com/alwiirfan/internal/dto/request"
	"github.com/alwiirfan/internal/dto/response"
	"github.com/alwiirfan/internal/services"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

// Create New Book
func (c *BookController) Create(ctx echo.Context) error {
	var req request.CreateBookRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create book",
			Error:      err.Error(),
		})
	}

	if err := c.bookService.Create(ctx.Request().Context(), &req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create book",
			Error:      err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, response.CommonResponse{
		StatusCode: http.StatusCreated,
		Message:    "Create book successfully",
		Data:       nil,
	})
}
