package main

import (
	"net/http"

	"golang-bookstore/app"
	"golang-bookstore/controller"
	"golang-bookstore/exception"
	"golang-bookstore/helper"
	"golang-bookstore/repository"
	"golang-bookstore/service"

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
