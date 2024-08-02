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

// Create New Book Controller
func (c *BookController) Create(ctx echo.Context) error {
	// check request url, if it has query params, return error
	if len(ctx.QueryParams()) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create book",
			Error:      "Invalid request url",
		})
	}

	var req request.CreateBookRequest

	// bind request body & validate request body
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create book",
			Error:      err.Error(),
		})
	}

	// create new book, if error, return error
	if err := c.bookService.Create(ctx.Request().Context(), &req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create book",
			Error:      err.Error(),
		})
	}

	// return created book success response
	return ctx.JSON(http.StatusCreated, response.CommonResponse{
		StatusCode: http.StatusCreated,
		Message:    "Create book successfully",
		Data:       nil,
	})
}

// Find All Books Controller with pagination
func (c *BookController) FindAll(ctx echo.Context) error {

	// init current_page and page_size for query params
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

	// check is query params valid
	for param := range ctx.QueryParams() {
		if !allowedParams[param] {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid query params: " + param,
				Error:      "Invalid url query params",
			})
		}
	}

	// find all books, if error, return error
	books, totalRecords, err := c.bookService.FindAll(ctx.Request().Context(), page, pageSize)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to find all books",
			Error:      err.Error(),
		})
	}

	// math total pages based on total records
	totalPages := totalRecords / pageSize
	if totalRecords%pageSize != 0 {
		totalPages++
	}

	// init pagination response
	paginationResp := response.PaginationResponse{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalData:   totalRecords,
		TotalPages:  totalPages,
	}

	// return find all books success response
	commonResponse := response.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "Find all books successfully",
		Data:       books,
		Pagination: paginationResp,
	}

	return ctx.JSON(http.StatusOK, commonResponse)
}

// Find Book By ID Controller
func (c *BookController) FindByID(ctx echo.Context) error {
	var bookId = ctx.Param("id")

	// check request url, if it has query params, return error
	if len(ctx.QueryParams()) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to find book by id",
			Error:      "Invalid request url",
		})
	}

	// find book by id, if error, return error
	book, err := c.bookService.FindByID(ctx.Request().Context(), bookId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Failed to find book by id, book not found",
			Error:      err.Error(),
		})
	}

	// return find book by id success response
	return ctx.JSON(http.StatusOK, response.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "Find book by id successfully",
		Data:       book,
	})

}

// Update Book Controller
func (c *BookController) Update(ctx echo.Context) error {
	var bookId = ctx.Param("id")
	var req request.UpdateBookRequest

	// check request url, if it has query params, return error
	if len(ctx.QueryParams()) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update book",
			Error:      "Invalid request url",
		})
	}

	// parse request body, if error, return error
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update book",
			Error:      err.Error(),
		})
	}

	// find book by id, if error, return error
	existingBook, err := c.bookService.FindByID(ctx.Request().Context(), bookId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Failed to update book, book not found",
			Error:      err.Error(),
		})
	}

	// update book, if error, return error
	if err := c.bookService.Update(ctx.Request().Context(), existingBook.Id, &req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update book",
			Error:      err.Error(),
		})
	}

	// return update book success response
	return ctx.JSON(http.StatusOK, response.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "Update book successfully",
	})
}

// Delete Book Controller
func (c *BookController) DeleteByID(ctx echo.Context) error {
	var bookId = ctx.Param("id")

	// check request url, if it has query params, return error
	if len(ctx.QueryParams()) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to delete book",
			Error:      "Invalid request url",
		})
	}

	// find book by id, if error, return error
	existingBook, err := c.bookService.FindByID(ctx.Request().Context(), bookId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Failed to delete book, book not found",
			Error:      err.Error(),
		})
	}

	// delete book, if error, return error
	if err := c.bookService.DeleteByID(ctx.Request().Context(), existingBook.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete book",
			Error:      err.Error(),
		})
	}

	// return delete book success response
	return ctx.JSON(http.StatusOK, response.CommonResponse{
		StatusCode: http.StatusOK,
		Message:    "Delete book successfully",
	})
}
