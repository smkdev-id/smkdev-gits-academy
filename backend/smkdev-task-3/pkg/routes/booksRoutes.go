package routes

import (
	"github.com/FuadGrimaldi/bookstore-api/pkg/controllers"
	"github.com/labstack/echo/v4"
)

// Route defines, structure for HTTP routes
type Route struct {
	Method 	string
	Path 	string
	Handler	echo.HandlerFunc
}

// BookRoutes Function
func BookRoutes(bookController *controllers.BookController) []*Route {
	return []*Route{
		{
			Method: echo.GET,
			Path: "/book",
			Handler: bookController.GetAllBooks,
		},
		{
			Method: echo.GET,
			Path: "/book/:id",
			Handler: bookController.GetOneBook,
		},
		{
			Method: echo.POST,
			Path: "/book",
			Handler: bookController.CreateBook,
		},
		{
			Method: echo.PUT,
			Path: "/book/:id",
			Handler: bookController.UpdateBook,
		},
		{
			Method: echo.DELETE,
			Path: "/book/:id",
			Handler: bookController.DeleteBook,
		},
	}
}