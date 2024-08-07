package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"
	"bookstore/pkg/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	logger := utils.InitializeLogger()

	e := echo.New()


	logger.Info("Initializing database")
	config.InitDB()

	logger.Info("Setting up routes")
	routes.BookStoreRoutes(e)

	logger.Info("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
