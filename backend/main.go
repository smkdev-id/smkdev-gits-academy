package main

import (
	"github.com/firyalhfz/backend-go-crud-api/controllers/productcontroller"
	"github.com/firyalhfz/backend-go-crud-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default();
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()

}