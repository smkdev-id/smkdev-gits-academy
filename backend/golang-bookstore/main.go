package main

import (
	"github.com/gin-gonic/gin"
	"golang-bookstore/pkg/config"
	"golang-bookstore/pkg/routes"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":3000")
}
