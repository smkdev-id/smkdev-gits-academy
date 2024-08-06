package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run()
}
