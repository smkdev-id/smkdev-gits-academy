package routers

import (
	"task3/PKG/controllers"

	"github.com/gin-gonic/gin"
)

// router setup
func SetupRouter() *gin.Engine {
	r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("/", controllers.CREATEBOOK)
		bookRoutes.GET("/", controllers.GETBOOK)
		bookRoutes.GET("/:id", controllers.GETBOOKBYID)
		bookRoutes.PUT("/:id", controllers.UPDATEBOOK)
		bookRoutes.DELETE("/:id", controllers.DELETEBOOK)
	}

	

	return r
}
