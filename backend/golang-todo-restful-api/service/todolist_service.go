package service

import (
	"context"

	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/web"
)

type TodoListService interface {
	Create(ctx context.Context, request web.TodoListCreateRequest) web.TodoListResponse
	Update(ctx context.Context, request web.TodoListUpdateRequest) web.TodoListResponse
	Delete(ctx context.Context, todolistId int)
	FindById(ctx context.Context, todolistId int) web.TodoListResponse
	FindAll(ctx context.Context) []web.TodoListResponse
}
