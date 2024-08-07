package main

import (
	"todo-crud/todos"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	todos.InitDB("todos.db")

	e.GET("/", todos.GetTodos)
	e.POST("/", todos.CreateTodo)
	e.PUT("/:id", todos.UpdateTodo)
	e.DELETE("/:id", todos.DeleteTodo)

	e.Logger.Fatal(e.Start(":8000")) // Menjalankan server pada port 8000
}
