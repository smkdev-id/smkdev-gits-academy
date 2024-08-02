package controllers

import (
	"net/http"
	"strconv"

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
	if len(ctx.QueryParams()) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create book",
			Error:      "Invalid request url",
		})
	}
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

// Find All Books Controller with pagination
func (c *BookController) FindAll(ctx echo.Context) error {

	page, _ := strconv.Atoi(ctx.QueryParam("current_page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(ctx.QueryParam("page_size"))

	if pageSize < 1 {
		pageSize = 10
	}

	allowedParams := map[string]bool{
		"current_page": true,
		"page_size":    true,
	}

	// check is query param valid
	for param := range ctx.QueryParams() {
		if !allowedParams[param] {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid query params: " + param,
				Error:      "Invalid url query params",
			})
		}
	}

	books, totalRecords, err := c.bookService.FindAll(ctx.Request().Context(), page, pageSize)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to find all books",
			Error:      err.Error(),
		})
	}

	totalPages := totalRecords / pageSize
	if totalRecords%pageSize != 0 {
		totalPages++
	}

	paginationResp := response.PaginationResponse{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalData:   totalRecords,
		TotalPages:  totalPages,
	}

	commonResponse := response.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "Find all books successfully",
		Data:       books,
		Pagination: paginationResp,
	}

	return ctx.JSON(http.StatusOK, commonResponse)
}
