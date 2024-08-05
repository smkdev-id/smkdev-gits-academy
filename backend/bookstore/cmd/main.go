package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/controllers"
	"bookstore/pkg/repository"
	"bookstore/pkg/routes"
	"bookstore/pkg/services"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := config.ConnectDB()

	repository := repository.NewBookRepository(db)
	service := services.NewBookService(repository)
	controller := controllers.NewBookController(service)

	route := e.Group("/book")
	routes.Routes(controller, route)

	e.Logger.Fatal(e.Start(":8080"))
}
