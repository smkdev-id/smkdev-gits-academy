package main

import (
	"log"
	"os"
	"strings"
	"task3/PKG/config"
	"task3/PKG/models"
	"task3/PKG/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    gin.SetMode(os.Getenv("GIN_MODE"))

	// check database connection
    log.Println("Connecting to database...")
    config.ConnectDatabase()
    log.Println("Database connection successful")

	// check migration database
    log.Println("Running migrations...")
    config.DB.AutoMigrate(&models.Book{})
    log.Println("Migrations completed")
	
	// check route
    log.Println("Setting up router...")
    r := routers.SetupRouter()

	//  allows to configure the trusted proxies ( Gin application through an env variable)
	trustedProxies := os.Getenv("TRUSTED_PROXIES")
    if trustedProxies != "" {
        proxyList := strings.Split(trustedProxies, ",")
        if err := r.SetTrustedProxies(proxyList); err != nil {
            log.Fatalf("Error setting trusted proxies: %v", err)
        }
        log.Printf("Trusted proxies set to: %v", proxyList)
    }

	// check to start 
    port := os.Getenv("SERVER_PORT")
    log.Printf("Starting server on port %s...", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
