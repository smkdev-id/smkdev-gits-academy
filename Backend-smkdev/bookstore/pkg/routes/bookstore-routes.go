package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookByID)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
