package main

import (
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/config"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	config.InitDB()
	config.InitRedis()

	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}