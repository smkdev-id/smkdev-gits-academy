package routes

import (
    "net/http"
    "task2_crud_golang/database"
    "task2_crud_golang/models"
    "github.com/gorilla/mux"
    "encoding/json"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/todos", GetTodos).Methods("GET")
    router.HandleFunc("/todos", CreateTodo).Methods("POST")
    router.HandleFunc("/todos/{id:[0-9]+}", UpdateTodo).Methods("PUT")
    router.HandleFunc("/todos/{id:[0-9]+}", DeleteTodo).Methods("DELETE")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
    var todos []models.Todo
    database.DB.Find(&todos)
    json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)
    database.DB.Create(&todo)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)
    database.DB.Model(&todo).Where("id = ?", id).Updates(todo)
    json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    database.DB.Delete(&models.Todo{}, id)
    w.WriteHeader(http.StatusNoContent)
}
