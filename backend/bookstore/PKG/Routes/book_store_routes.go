package routes

import (
	"net/http"

	"bookstore/PKG/controllers"
)

func HandleRequests() {
	http.HandleFunc("/books", controllers.CreateBook)
	// ... tambahkan endpoint lain
}
