package routes

import (
    "bookstore/pkg/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, bookController *controllers.BookController) {
    router.GET("/book", bookController.GetBooks)
    router.POST("/book", bookController.CreateBook)
    router.GET("/book/:id", bookController.GetBookByID)
    router.PUT("/book/:id", bookController.UpdateBook)
    router.DELETE("/book/:id", bookController.DeleteBook)
}