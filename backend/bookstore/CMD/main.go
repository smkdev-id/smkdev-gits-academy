package main

import (
	"bookstore/PKG/config"
	"bookstore/PKG/routes"
	"net/http"
)

func main() {
	config.Connect()
	routes.HandleRequests()
	http.ListenAndServe(":8080", nil)
}
