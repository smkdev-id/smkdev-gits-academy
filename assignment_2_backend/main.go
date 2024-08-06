package main

import (
	"book_sell/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/book", controllers.GetBooks)
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:bookId", controllers.GetBookByID)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)

	router.Run(":8080")
}
