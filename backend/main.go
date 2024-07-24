package main

import (
    "log"
    "net/http"
    "todolist/db"
    "todolist/routes"
    "github.com/gorilla/mux"
)

func main() {
    database.ConnectDatabase()
    r := mux.NewRouter()
    routes.RegisterRoutes(r)
    log.Fatal(http.ListenAndServe(":8080", r))
}
