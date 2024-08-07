package main

import (
	"github.com/PorcoGalliard/BookStore/pkg/models"
	"github.com/PorcoGalliard/BookStore/pkg/config"
	"github.com/PorcoGalliard/BookStore/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.Book{})

    routes.RegisterBookRoutes(r)

    r.Run(":8080")
}
