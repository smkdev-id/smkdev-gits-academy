package main

import (
	"net/http"

	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/app"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/controller"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/exception"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/helper"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/repository"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	router := httprouter.New()

	todoListRepository := repository.NewTodoListRepository()
	todoListService := service.NewTodoListService(todoListRepository, db, validate)
	todoListController := controller.NewTodoListController(todoListService)

	router.GET("/api/todolist", todoListController.FindAll)
	router.GET("/api/todolist/:todolistId", todoListController.FindById)
	router.POST("/api/todolist", todoListController.Create)
	router.PUT("/api/todolist/:todolistId", todoListController.Update)
	router.DELETE("/api/todolist/:todolistId", todoListController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
