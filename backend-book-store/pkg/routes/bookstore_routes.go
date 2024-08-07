package routes

import (
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/controllers"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	
	userProtected := router.Group("/")
	userProtected.Use(middleware.AuthMiddleware())
	
	userProtected.GET("/book", controllers.GetBooks)
	userProtected.GET("/book/:bookId", controllers.GetBookById)
	
	userProtected.POST("/logout", controllers.Logout)

	adminProtected := router.Group("/")
	adminProtected.Use(middleware.AuthMiddleware(), middleware.AdminOnly())

	adminProtected.POST("/book", controllers.CreateBook)
	adminProtected.PUT("/book/:bookId", controllers.UpdateBook)
	adminProtected.DELETE("/book/:bookId", controllers.DeleteBook)

	adminProtected.POST("/role", controllers.CreateRole)
	adminProtected.GET("/role", controllers.GetRoles)
	adminProtected.POST("/assign-role/:userID", controllers.AssignRole)
}	