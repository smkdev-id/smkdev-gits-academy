package routes

import (
    "github.com/PorcoGalliard/BookStore/pkg/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine) {
    r.GET("/book", controllers.GetBooks)
    r.POST("/book", controllers.CreateBook)
    r.GET("/book/:bookid", controllers.GetBookByID)
    r.PUT("/book/:bookid", controllers.UpdateBook)
    r.DELETE("/book/:bookid", controllers.DeleteBook)
}
