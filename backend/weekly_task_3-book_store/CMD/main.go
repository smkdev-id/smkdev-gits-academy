package main

import (
	"books_store/PKG/config"
	"books_store/PKG/models"
	"books_store/PKG/routes"
	"books_store/PKG/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize DB connection
	config.InitDB()

	// Set custom error handler
	e.HTTPErrorHandler = utils.HTTPErrorHandler

	// Migrate the schema
	models.Migrate(config.DB)

	// Initialize routes
	routes.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
