package routes

import (
	"bookstore/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/book", controllers.GetBooks)
	e.POST("/book", controllers.CreateBook)
	e.GET("/book/:bookid", controllers.GetBookByID)
	e.PUT("/book/:bookid", controllers.UpdateBook)
	e.DELETE("/book/:bookid", controllers.DeleteBook)
}
