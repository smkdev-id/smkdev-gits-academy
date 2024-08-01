package dto

import (
	"todos/entity"
)

//Struct Untuk request Todo
type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

//Struct Untuk response Todo
type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}


//Method Untuk mapping request ke entity
func (t *TodoRequest) ToEntity() *entity.Todo {
	return &entity.Todo{
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

//Function Untuk mapping entity ke response
func EntityToResponse(todo *entity.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   todo.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
