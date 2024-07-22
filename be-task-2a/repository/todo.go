package repository

import (
	"context"
	"github.com/HerlyRyan/be-task-2a/model"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
	GetTodoByID(ctx context.Context, id int) (*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}
