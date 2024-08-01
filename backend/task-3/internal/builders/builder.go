package builders

import (
	"github.com/alwiirfan/configs"
	"github.com/alwiirfan/internal/http/controllers"
	"github.com/alwiirfan/internal/http/routes"
	"github.com/alwiirfan/internal/repositories"
	"github.com/alwiirfan/internal/services"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *configs.Config, db *gorm.DB) []*routes.Route {
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(cfg, bookRepository)
	bookController := controllers.NewBookController(bookService)

	return routes.PublicRoutes(bookController)
}

func BuildPrivateRoutes(cfg *configs.Config, db *gorm.DB) []*routes.Route {
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(cfg, bookRepository)
	bookController := controllers.NewBookController(bookService)

	return routes.PrivateRoutes(bookController)
}
