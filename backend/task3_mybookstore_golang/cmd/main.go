package main

import (
    "task3_mybookstore_golang/pkg/routes"
    "log"
    "net/http"
)

func main() {
    router := routes.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
