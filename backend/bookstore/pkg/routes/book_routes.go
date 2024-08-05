package routes

import (
	"bookstore/pkg/controllers"

	"github.com/labstack/echo"
)

func Routes(controller *controllers.BookController, e *echo.Group) {
	e.GET("/", controller.GetAllBooks)
	e.GET("/:id", controller.GetBookByID)
	e.POST("/", controller.CreateBook)
	e.PUT("/:id", controller.UpdateBook)
	e.DELETE("/:id", controller.DeleteBook)
}
