package main

import (
	"net/http"
	"task-2-golang-todo-crud-api/database"
	"task-2-golang-todo-crud-api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitDB("database/todo_database.db")

	e := echo.New()
	e.GET("/tes", func(c echo.Context) error {
		return c.String(http.StatusOK, "testing berhasil")
	})
	e.GET("/items", handlers.GetItems(db))
	e.GET("/items/:id", handlers.GetItem(db))
	e.POST("/items", handlers.CreateItem(db))
	e.PUT("/items/:id", handlers.UpdateItem(db))
	e.DELETE("/items/:id", handlers.DeleteItem(db))

	e.Logger.Fatal(e.Start(":8080"))
}
