package routes

import (
	"github.com/gin-gonic/gin"
	"book_sell/controllers"
)

func BookRoutes(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)  
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:bookId", controllers.GetBookByID) 
	r.PUT("/books/:bookId", controllers.UpdateBook)  
	r.DELETE("/books/:bookId", controllers.DeleteBook)
}
