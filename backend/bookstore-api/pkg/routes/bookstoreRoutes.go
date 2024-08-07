package routes

import (
	"bookstore-api/pkg/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func BookRoutes(bookcontroller *controllers.BookController) []*Route {
	return []*Route{
		{
			Method:  "GET",
			Path:    "/",
			Handler: bookcontroller.GetAllBooks,
		},
		{
			Method:  "GET",
			Path:    "/book/:id",
			Handler: bookcontroller.GetBookByID,
		},
		{
			Method:  "POST",
			Path:    "/book",
			Handler: bookcontroller.CreateBook,
		},
		{
			Method:  "PUT",
			Path:    "/book/:id",
			Handler: bookcontroller.UpdateBook,
		},
		{
			Method:  "DELETE",
			Path:    "/book/:id",
			Handler: bookcontroller.DeleteBook,
		},
	}
}
