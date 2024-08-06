package main

import (
	"bookstore-api/pkg/config"
	"bookstore-api/pkg/controllers"
	"bookstore-api/pkg/models"
	"bookstore-api/pkg/routes"
	"bookstore-api/pkg/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Database
	db, err := config.ConnectDB()
	utils.CheckError(err)
	db.AutoMigrate(&models.Book{})

	// Initialize Controller
	bookController := controllers.NewBookController()

	// Initialize Routes
	bookRoutes := routes.BookRoutes(bookController)

	// Initialize Server
	e := echo.New()
	for _, v := range bookRoutes {
		e.Add(v.Method, v.Path, v.Handler)
	}

	e.Logger.Fatal(e.Start(":8080"))

}
