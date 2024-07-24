package main

import (
    "github.com/labstack/echo/v4"
    "github.com/PorcoGalliard/Todo-List/api"
    "github.com/PorcoGalliard/Todo-List/db"
)

func main() {
    db.InitDB("todos.db")

    e := echo.New()

    e.GET("/", api.GetTodos)
    e.POST("/", api.CreateTodo)
    e.PUT("/:id", api.UpdateTodo)
    e.DELETE("/:id", api.DeleteTodo)

    e.Start(":8080")
}
