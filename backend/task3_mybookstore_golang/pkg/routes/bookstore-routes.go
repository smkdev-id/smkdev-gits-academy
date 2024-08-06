package routes

import (
    "task3_mybookstore_golang/pkg/controllers"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    RegisterBookStoreRoutes(router)
    return router
}



var RegisterBookStoreRoutes = func(router *mux.Router) {
    router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
    router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
