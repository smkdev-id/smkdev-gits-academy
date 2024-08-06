package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"
	"log"
)

func main() {
	// Initialize the database connection
	config.ConnectDatabase()

	// Set up the router
	r := routes.SetupRouter()

	// Start the server
	log.Println("Starting the server on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
