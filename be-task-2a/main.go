package main

import (
	"github.com/HerlyRyan/be-task-2a/controller"
	"github.com/HerlyRyan/be-task-2a/model"
	"github.com/HerlyRyan/be-task-2a/repository"
	"github.com/HerlyRyan/be-task-2a/route"
	"github.com/HerlyRyan/be-task-2a/db"
)

func main() {
	db := db.Connectdb()

	db.AutoMigrate(&model.Todo{})

	todoRepo := repository.NewTodoRepository(db)

	todoController := controller.NewTodoController(todoRepo)

	app := &route.App{TodoController: todoController}
	app.InitializeRoutes()

	app.Run()
}
