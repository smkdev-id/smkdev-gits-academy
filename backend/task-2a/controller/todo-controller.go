package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alwi09/dto/request"
	"github.com/alwi09/dto/response"
	"github.com/alwi09/service"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{
		todoService,
	}
}

func (c *TodoController) GetAll(rw http.ResponseWriter, r *http.Request) {
	todos, err := c.todoService.FindAll()

	rw.Header().Set("Content-Type", "application/json")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to find all todos",
			Error:   "Check internet connection",
		})
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to find all todos: %v", err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.CommonResponse[[]response.TodoResponse]{
		Status:  http.StatusOK,
		Message: "Find all todos successfully",
		Data:    *todos,
	})
}

func (c *TodoController) FindById(rw http.ResponseWriter, r *http.Request, id string) {

	rw.Header().Set("Content-Type", "application/json")

	if id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to find todo",
			Error:   "id is required",
		})
		http.Error(rw, "id is required", http.StatusBadRequest)
		log.Printf("Failed to find todo: id is required")
		return
	}

	todo, err := c.todoService.FindById(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Failed to find todo",
			Error:   "Todo not found",
		})
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to find todo: %v", err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.CommonResponse[response.TodoResponse]{
		Status:  http.StatusOK,
		Message: "Find todo successfully",
		Data:    *todo,
	})
}

func (c *TodoController) Create(rw http.ResponseWriter, r *http.Request) {
	var todoReq request.TodoCreateRequest

	rw.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create todo",
			Error:   "Invalid request body",
		})
		http.Error(rw, err.Error(), http.StatusBadRequest)
		log.Printf("Failed to create todo: %v", err.Error())
		return
	}

	todo, err := c.todoService.Create(&todoReq)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create todo",
			Error:   "Check internet connection",
		})
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to create todo: %v", err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.CommonResponse[response.TodoResponse]{
		Status:  http.StatusOK,
		Message: "Create todo successfully",
		Data:    *todo,
	})
}

func (c *TodoController) Update(rw http.ResponseWriter, r *http.Request, id string) {
	var todoReq request.TodoUpdateRequest

	rw.Header().Set("Content-Type", "application/json")

	if id == "" {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Failed to update todo",
			Error:   "id is required",
		})
		http.Error(rw, "id is required", http.StatusNotFound)
		log.Printf("Failed to update todo: id is required")
		return
	}

	existingTodo, err := c.todoService.FindById(id)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Failed to update todo",
			Error:   "Todo not found",
		})
		http.Error(rw, err.Error(), http.StatusNotFound)
		log.Printf("Failed to update todo: %v", err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update todo",
			Error:   "Invalid request body",
		})
		http.Error(rw, err.Error(), http.StatusBadRequest)
		log.Printf("Failed to update todo: %v", err.Error())
		return
	}

	updatedTodo, err := c.todoService.Update(existingTodo.Id, &todoReq)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update todo",
			Error:   "Check internet connection",
		})
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to update todo: %v", err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.CommonResponse[response.TodoResponse]{
		Status:  http.StatusOK,
		Message: "Update todo successfully",
		Data:    *updatedTodo,
	})
}

func (c *TodoController) Delete(rw http.ResponseWriter, r *http.Request, id string) {
	rw.Header().Set("Content-Type", "application/json")

	if id == "" {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Failed to delete todo",
			Error:   "id is required",
		})
		http.Error(rw, "id is required", http.StatusNotFound)
		log.Printf("Failed to delete todo: id is required")
		return
	}

	existingTodo, err := c.todoService.FindById(id)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Failed to delete todo",
			Error:   "Todo not found",
		})
		http.Error(rw, err.Error(), http.StatusNotFound)
		log.Printf("Failed to delete todo: %v", err.Error())
		return
	}

	err = c.todoService.Delete(existingTodo.Id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete todo",
			Error:   "Check internet connection.",
		})
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to delete todo: %v", err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.CommonResponse[interface{}]{
		Status:  http.StatusOK,
		Message: "Delete todo successfully",
	})
}
