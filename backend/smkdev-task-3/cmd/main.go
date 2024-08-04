package main

import (
	"github.com/FuadGrimaldi/bookstore-api/pkg/config"
	"github.com/FuadGrimaldi/bookstore-api/pkg/controllers"
	"github.com/FuadGrimaldi/bookstore-api/pkg/database"
	"github.com/FuadGrimaldi/bookstore-api/pkg/routes"
	"github.com/FuadGrimaldi/bookstore-api/pkg/server"
)

func main() {
	// Load application configuration from the .env file
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// Connect to the SQLite database
	db, err := database.ConnectToSqlite(cfg)
	checkError(err)

	// Initialize controllers
	bookController := controllers.NewBookController(db)

	// Define routes
	bookRoutes := routes.BookRoutes(bookController)

	// Create a new server
	s := server.NewServer(cfg, bookRoutes)

	// Start the server
	s.Run()

	// Handle graceful shutdown
	s.GracefulShutdown()
}

// checkError checks
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}