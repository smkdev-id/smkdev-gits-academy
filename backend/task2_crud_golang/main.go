package main

import (
    "log"
    "net/http"
    "task2_crud_golang/database"
    "task2_crud_golang/routes"
    "github.com/gorilla/mux"
)

func main() {
    database.ConnectDatabase()
    r := mux.NewRouter()
    routes.RegisterRoutes(r)
    log.Fatal(http.ListenAndServe(":8080", r))
}
