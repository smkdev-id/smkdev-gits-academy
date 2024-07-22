package helper

import (
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/domain"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/web"
)

func TodoListResponse(todolist domain.TodoList) web.TodoListResponse {
	return web.TodoListResponse{
		Id:          todolist.Id,
		Description: todolist.Description,
		Completed:   todolist.Completed,
	}
}

func TodoListResponses(todolist []domain.TodoList) []web.TodoListResponse {
	var TodoListResponses []web.TodoListResponse
	for _, TodoList := range todolist {
		TodoListResponses = append(TodoListResponses, TodoListResponse(TodoList))
	}

	return TodoListResponses
}
