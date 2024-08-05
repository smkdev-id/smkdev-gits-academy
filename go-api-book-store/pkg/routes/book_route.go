package routes

import (
	"github.com/hasban-fardani/go-api-book-store/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func BookRoute(e *echo.Echo) {

	bookController := controllers.NewBookController()

	e.GET("/book", bookController.Index)
	e.GET("/book/:id", bookController.Show)
	e.POST("/book", bookController.Create)
	e.PUT("/book/:id", bookController.Update)
	e.DELETE("/book/:id", bookController.Delete)
}
