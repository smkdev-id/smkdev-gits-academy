package controller

import (
	"EkoEdyPurwanto/task-2/model/dto/req"
	"EkoEdyPurwanto/task-2/usecase"
	"encoding/json"
	"net/http"
)

type TodoListController struct {
	UseCase usecase.TodoListUseCase
	Engine  *http.ServeMux
}

// Route auth
func (t *TodoListController) Route() {

	t.Engine.HandleFunc("/api/v1/todos", t.createHandler)
}

func (t *TodoListController) createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload req.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := t.UseCase.Create(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("TodoList created successfully"))
}

// NewTodoListController Constructor
func NewTodoListController(useCase usecase.TodoListUseCase, engine *http.ServeMux) *TodoListController {
	return &TodoListController{
		UseCase: useCase,
		Engine:  engine,
	}
}
