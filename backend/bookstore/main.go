package main

import (
    "log"
    "bookstore/pkg/config"
    "bookstore/pkg/controllers"
    "bookstore/pkg/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    db, err := config.SetupDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    bookController := controllers.NewBookController(db)

    router := gin.Default()
    routes.SetupRoutes(router, bookController)

    router.Run(":8080")
}