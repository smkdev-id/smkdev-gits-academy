package main

import (
    "github.com/HerlyRyan/be-task-3/pkg/config"
    "github.com/HerlyRyan/be-task-3/pkg/routes"
    "github.com/HerlyRyan/be-task-3/pkg/models"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Book{})
    routes.RegisterBookStoreRoutes(router)

    router.Run(":8080")
}
