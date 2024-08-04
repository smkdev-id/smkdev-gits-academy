package routes

import (
	"book-sales-management/controllers"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes the routes for the API
func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Book routes
	router.GET("/book/", controllers.GetAllBooks)
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/:bookId", controllers.GetBookByID)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)

	return router
}