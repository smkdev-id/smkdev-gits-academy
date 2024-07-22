package repository

import (
	"context"
	"database/sql"

	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/domain"
)

type TodoListRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todolist domain.TodoList) domain.TodoList
	Update(ctx context.Context, tx *sql.Tx, todolist domain.TodoList) domain.TodoList
	Delete(ctx context.Context, tx *sql.Tx, todolist domain.TodoList)
	FindById(ctx context.Context, tx *sql.Tx, todolistId int) (domain.TodoList, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.TodoList
}
