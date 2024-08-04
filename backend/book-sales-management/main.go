package main

import (
	"book-sales-management/config"
	"book-sales-management/routes"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize the database
	config.InitDB()

	// Initialize the router
	router := routes.InitRoutes()

	// Start the server
	fmt.Println("Server is running at port 8080")
	log.Fatal(router.Run(":8080"))
}