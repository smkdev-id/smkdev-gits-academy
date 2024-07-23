package controller

import (
	"EkoEdyPurwanto/task-2/helper"
	"EkoEdyPurwanto/task-2/model/dto/req"
	"EkoEdyPurwanto/task-2/usecase"
	"encoding/json"
	"net/http"
)

type TodoListController struct {
	UseCase usecase.TodoListUseCase
	Engine  *http.ServeMux
}

// Route todolist
func (t *TodoListController) Route() {
	apiMux := http.NewServeMux()

	apiMux.Handle("/create/todos", helper.Post(http.HandlerFunc(t.createHandler)))
	apiMux.Handle("/fetch/todos", helper.Get(http.HandlerFunc(t.getAllHandler)))
	apiMux.Handle("/fetch/todos/byIdentifier", helper.Get(http.HandlerFunc(t.getByPayloadHandler)))
	apiMux.Handle("/update/todos", helper.Put(http.HandlerFunc(t.updateHandler)))
	apiMux.Handle("/delete/todos", helper.Delete(http.HandlerFunc(t.deleteHandler)))

	t.Engine.Handle("/api/v1/", http.StripPrefix("/api/v1", apiMux))
}

// createHandler implements TodoListController.
func (t *TodoListController) createHandler(w http.ResponseWriter, r *http.Request) {

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

// getAllHandler implements TodoListController.
func (t *TodoListController) getAllHandler(w http.ResponseWriter, r *http.Request) {

	todolists, err := t.UseCase.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve todo lists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todolists); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// getByPayloadHandler implements TodoListController.
func (t *TodoListController) getByPayloadHandler(w http.ResponseWriter, r *http.Request) {
	payload := r.URL.Query().Get("status")
	if payload == "" {
		http.Error(w, "Missing identifier", http.StatusBadRequest)
		return
	}

	todolists, err := t.UseCase.GetByPayload(payload)
	if err != nil {
		http.Error(w, "Failed to retrieve todo lists", http.StatusInternalServerError)
		return
	}

	if len(todolists) == 0 {
		http.Error(w, "No todo lists found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todolists); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// updateHandler implements TodoListController.
func (t *TodoListController) updateHandler(w http.ResponseWriter, r *http.Request) {

	var payload req.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := t.UseCase.UpdateStatus(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status update successfully"))
}

// deleteHandler implements TodoListController.
func (t *TodoListController) deleteHandler(w http.ResponseWriter, r *http.Request) {

	var payload string
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := t.UseCase.Delete(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TodoList deleted successfully"))
}

// NewTodoListController Constructor
func NewTodoListController(useCase usecase.TodoListUseCase, engine *http.ServeMux) *TodoListController {
	return &TodoListController{
		UseCase: useCase,
		Engine:  engine,
	}
}
