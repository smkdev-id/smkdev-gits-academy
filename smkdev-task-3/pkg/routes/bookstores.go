package routes

import (
    "smkdev-task-3/pkg/controllers"
    "github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {
    e := echo.New()

    e.GET("/books", controllers.GetBooks)
    e.POST("/books", controllers.CreateBook)
    e.GET("/books/:id", controllers.GetBook)
    e.PUT("/books/:id", controllers.UpdateBook)
    e.DELETE("/books/:id", controllers.DeleteBook)

    return e
}
