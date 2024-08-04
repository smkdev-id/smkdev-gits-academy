package routes

import (
	"net/http"

	"github.com/alwiirfan/internal/http/controllers"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func Routes(bookController *controllers.BookController) []*Route {
	return []*Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/v1/books",
			Handler: bookController.Create,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/v1/books",
			Handler: bookController.FindAll,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/api/v1/books",
			Handler: bookController.DeleteAll,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/v1/books/:id",
			Handler: bookController.FindByID,
		},
		{
			Method:  http.MethodPut,
			Path:    "/api/v1/books/:id",
			Handler: bookController.Update,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/api/v1/books/:id",
			Handler: bookController.DeleteByID,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/v1/books/isbn/:isbn",
			Handler: bookController.FindByISBN,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/v1/books/search",
			Handler: bookController.FindAllSearch,
		},
	}
}
