package routes

import (
	"github.com/gin-gonic/gin"
	"golang-bookstore/pkg/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
		bookRoutes.GET("/", controllers.GetAllBooks)
	}
}
