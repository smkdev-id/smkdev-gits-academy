package utils

import (
	"go-bookstore/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Meta struct for metadata
type Meta struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
    Status  string `json:"status"`
}

// Response struct for API responses
type Response struct {
    Meta Meta        `json:"meta"`
    Data interface{} `json:"data"`
}

// APIResponse formats the API response
func APIResponse(message string, code int, status string, data interface{}) Response {
    return Response{
        Meta: Meta{
            Message: message,
            Code:    code,
            Status:  status,
        },
        Data: data,
    }
}

// BookFormatter struct to format book data
type BookFormatter struct {
	ID          uint   `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    PageCount   int    `json:"page_count"`
    Price       string `json:"price"`
    Author      string `json:"author"`
    Quantity    int    `json:"quantity"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}

// FormatBook formats a single book's data
func FormatBook(book models.Book) BookFormatter {
    return BookFormatter{
        ID:        book.ID,
        Title:     book.Title,
		Description: book.Description,
		PageCount: book.PageCount,
        Price:     book.FormatPrice(),
        Author:    book.Author,
        Quantity:  book.Quantity,
        CreatedAt: book.CreatedAt.Format("02 Januari 2006 15:04"),
        UpdatedAt: book.UpdatedAt.Format("02 Januari 2006 15:04"),
    }
}

// FormatBooks formats a list of books' data
func FormatBooks(books []models.Book) []BookFormatter {
    formattedBooks := make([]BookFormatter, len(books))
    for i, book := range books {
        formattedBooks[i] = FormatBook(book)
    }
    return formattedBooks
}

// ResponseSuccess formats a successful response
func ResponseSuccess(c echo.Context, message string, data interface{}) error {
    response := APIResponse(message, http.StatusOK, "success", data)
    return c.JSON(http.StatusOK, response)
}

// ResponseError formats an error response
func ResponseError(c echo.Context, statusCode int, message string) error {
    response := APIResponse(message, statusCode, "error", nil)
    return c.JSON(statusCode, response)
}

