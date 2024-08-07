package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/HerlyRyan/be-task-3/pkg/controllers"
)

func RegisterBookStoreRoutes(router *gin.Engine) {
    router.POST("/book", controllers.CreateBook)
    router.GET("/book", controllers.GetBooks)
    router.GET("/book/:bookid", controllers.GetBookByID)
    router.PUT("/book/:bookid", controllers.UpdateBook)
    router.DELETE("/book/:bookid", controllers.DeleteBook)
}
