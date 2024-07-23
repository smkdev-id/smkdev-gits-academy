package dto

import (
	"todos/entity"
)

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *TodoRequest) ToEntity() *entity.Todo {
	return &entity.Todo{
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

func EntityToResponse(todo *entity.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}
}
