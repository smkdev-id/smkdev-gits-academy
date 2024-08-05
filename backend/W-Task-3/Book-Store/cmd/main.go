package main

import (
	"log"
	"net/http"

	"Book-Store/pkg/config"
	"Book-Store/pkg/routes"

	"github.com/gorilla/handlers"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Initialize the router
	r := routes.RegisterBookStoreRoutes()

	   // CORS middleware
	   corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )(r)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
