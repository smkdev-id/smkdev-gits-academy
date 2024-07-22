package main

import (
	"go-todo/db"
	"go-todo/handlers"
	"go-todo/services"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    db.ConnectDatabase()
    todoService := services.NewTodoService(db.DB)
    todoHandler := handlers.NewTodoHandler(todoService)

    e.GET("/", todoHandler.FetchTodos)
    e.POST("/", todoHandler.CreateTodos)
    e.PUT("/:id", todoHandler.UpdateTodos)
    e.DELETE("/:id", todoHandler.DeleteTodos)

    e.Logger.Fatal(e.Start(":8080"))
}
