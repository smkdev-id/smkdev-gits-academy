package service

import (
	"context"

	"golang-bookstore/model/web"
)

type TodoListService interface {
	Create(ctx context.Context, request web.TodoListCreateRequest) web.TodoListResponse
	Update(ctx context.Context, request web.TodoListUpdateRequest) web.TodoListResponse
	Delete(ctx context.Context, todolistId int)
	FindById(ctx context.Context, todolistId int) web.TodoListResponse
	FindAll(ctx context.Context) []web.TodoListResponse
}
