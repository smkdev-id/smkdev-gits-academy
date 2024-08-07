package utils

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "smkdev-task-3/pkg/models"
)

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func ResponseError(c echo.Context, code int, message string) error {
    return c.JSON(code, Response{
        Status:  "error",
        Message: message,
    })
}

func ResponseSuccess(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusOK, Response{
        Status:  "success",
        Message: message,
        Data:    data,
    })
}

func FormatBook(book models.Book) map[string]interface{} {
    return map[string]interface{}{
        "id":          book.ID,
        "title":       book.Title,
        "description": book.Description,
        "page_count":  book.PageCount,
        "author":      book.Author,
        "price":       book.FormatPrice(),
        "quantity":    book.Quantity,
        "created_at":  models.FormatDate(book.CreatedAt),
        "updated_at":  models.FormatDate(book.UpdatedAt),
    }
}

func FormatBooks(books []models.Book) []map[string]interface{} {
    var formattedBooks []map[string]interface{}
    for _, book := range books {
        formattedBooks = append(formattedBooks, FormatBook(book))
    }
    return formattedBooks
}
