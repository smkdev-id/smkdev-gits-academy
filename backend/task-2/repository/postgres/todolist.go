package postgres

import (
	"EkoEdyPurwanto/task-2/model"
	"database/sql"
)

type (
	TodoListRepository interface {
		Save(todoList model.TodoList) error
		FindAll() error
		Update(id string) error
		Delete(id string) error
	}

	todoListRepository struct {
		DB *sql.DB
	}
)

// Delete implements TodoListRepository.
func (t *todoListRepository) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements TodoListRepository.
func (t *todoListRepository) FindAll() error {
	panic("unimplemented")
}

// Update implements TodoListRepository.
func (t *todoListRepository) Update(id string) error {
	panic("unimplemented")
}

func (t *todoListRepository) Save(todoList model.TodoList) error {
	SQL := `INSERT INTO "todo"(id, title, description, status, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := t.DB.Exec(SQL,
		&todoList.Id,
		&todoList.Title,
		&todoList.Description,
		&todoList.Status,
		&todoList.CreatedAt,
		&todoList.UpdatedAt,
		&todoList.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// NewTodoListRepository Constructor
func NewTodoListRepository(db *sql.DB) TodoListRepository {
	return &todoListRepository{
		DB: db,
	}
}
