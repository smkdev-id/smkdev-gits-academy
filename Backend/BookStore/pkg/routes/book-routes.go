package routes

import (
	"BookStore/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	g := e.Group("/api")
	g.GET("/book", controllers.GetBooks)
	g.GET("/book/:id", controllers.GetBookByID)
	g.POST("/book", controllers.CreateBook)
	g.PUT("/book/:id", controllers.UpdateBook)
	g.DELETE("/book/:id", controllers.DeleteBook)
}
