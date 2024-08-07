package routes

import (
	"bookstore/pkg/controller"

	"github.com/labstack/echo/v4"
)

func BookStoreRoutes(e *echo.Echo) {
	v := e.Group("/api/v1")
	v.GET("/books", controller.GetBooks)
	v.GET("/books/:bookId", controller.GetBookById)
	v.POST("/books", controller.CreateBook)
	v.PATCH("/books/:bookId", controller.UpdateBook)
	v.DELETE("/books/:bookId", controller.DeleteBook)
}
