package routes

import (
	"books_store/PKG/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/book", controllers.GetBooks)
	e.POST("/book", controllers.CreateBook)
	e.GET("/book/:id", controllers.GetBookById)
	e.PUT("/book/:id", controllers.UpdateBook)
	e.DELETE("/book/:id", controllers.DeleteBook)
}
