package main

import (
	"todo-list/config"
	"todo-list/controller"
	"todo-list/repository"
	"todo-list/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.ConnectDB()

	repository := repository.NewTodoRepository(db)
	service := service.NewTodoService(repository)
	controller := controller.NewTodoController(service)

	e.GET("/", controller.GetAllTodo)
	e.GET("/:id", controller.GetTodoByID)
	e.POST("/add", controller.CreateTodo)
	e.PUT("/edit/:id", controller.UpdateTodo)
	e.DELETE("/delete/:id", controller.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
