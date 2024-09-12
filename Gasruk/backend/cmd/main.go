package main

import (
	"gasruk/pkg/config"
	"gasruk/pkg/controllers"
	"gasruk/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi database
	db := config.ConnectDatabase()

	// Inisialisasi controller dengan database
	controllers.Initialize(db)

	// Buat instance Echo
	e := echo.New()
	routes.SetupRoutes(e)

	// Serve static files
	e.Static("/static", "pkg/static")

	// Run the server
	e.Start(":8080")
}
