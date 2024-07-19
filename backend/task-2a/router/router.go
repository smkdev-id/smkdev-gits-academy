package router

import (
	"net/http"
	"strings"

	"github.com/alwi09/controller"
)

func NewRouter(todoController *controller.TodoController) *http.ServeMux {
	mux := http.NewServeMux()

	// set routes for create todo and find all todos
	mux.HandleFunc("/todos", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoController.GetAll(rw, r)
		case http.MethodPost:
			todoController.Create(rw, r)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//  set routes for find todo by id, update todo and delete todo
	mux.HandleFunc("/todos/", func(rw http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 || parts[2] == "" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		id := parts[2]

		switch r.Method {
		case http.MethodGet:
			todoController.FindById(rw, r, id)
		case http.MethodPut:
			todoController.Update(rw, r, id)
		case http.MethodDelete:
			todoController.Delete(rw, r, id)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
