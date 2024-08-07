package routes

import (
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/controllers"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/role", controllers.CreateRole)
	router.GET("/role", controllers.GetRoles)
	router.POST("/assign-role/:userID", controllers.AssignRole)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/book", controllers.GetBooks)
	protected.POST("/book", controllers.CreateBook)
	protected.GET("/book/:bookId", controllers.GetBookById)
	protected.PUT("/book/:bookId", controllers.UpdateBook)
	protected.DELETE("/book/:bookId", controllers.DeleteBook)
	protected.POST("/logout", controllers.Logout)
}	