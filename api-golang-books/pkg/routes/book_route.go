package routes

import (
	"github.com/labstack/echo"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/controllers"
)

func BookRoute(e *echo.Echo) {

	bookController := controllers.NewBookController()

	e.GET("/books", bookController.Index)
	e.GET("/books/:id", bookController.Show)
	e.POST("/books", bookController.Create)
	e.PUT("/books/:id", bookController.Update)
	e.DELETE("/books/:id", bookController.Delete)
}
