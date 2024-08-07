package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Rute untuk Book
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookByID)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Rute untuk Kategori
	r.GET("/kategoris", controllers.GetKategoris)
	r.GET("/kategoris/:id", controllers.GetKategoriByID)
	r.POST("/kategoris", controllers.CreateKategori)
	r.PUT("/kategoris/:id", controllers.UpdateKategori)
	r.DELETE("/kategoris/:id", controllers.DeleteKategori)
}
