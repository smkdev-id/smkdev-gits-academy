package repository

import (
	"database/sql"
	"time"

	"github.com/alwi09/entity"
)

type (
	TodoRepository interface {
		FindAll() ([]entity.Todo, error)
		FindById(id string) (*entity.Todo, error)
		Create(todo *entity.Todo) (*entity.Todo, error)
		Update(id string, todo *entity.Todo) (*entity.Todo, error)
		Delete(id string) error
	}

	todoRepository struct {
		DB *sql.DB
	}
)

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db}
}

// FindAll implements TodoRepository
func (r *todoRepository) FindAll() ([]entity.Todo, error) {
	SQL := "SELECT id, title, description, status, created_at, updated_at FROM todo"
	rows, err := r.DB.Query(SQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		var createdAtStr, updatedAtStr sql.NullString

		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &createdAtStr, &updatedAtStr); err != nil {
			return nil, err
		}

		// Parse string to time.Time only if it's not null
		if createdAtStr.Valid {
			if createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				todo.CreatedAt = createdAt
			} else {
				return nil, err
			}
		} else {
			todo.CreatedAt = time.Time{}
		}

		if updatedAtStr.Valid {
			if updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				todo.UpdatedAt = updatedAt
			} else {
				return nil, err
			}
		} else {
			todo.UpdatedAt = time.Time{}
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// FindById implements TodoRepository
func (r *todoRepository) FindById(id string) (*entity.Todo, error) {
	SQL := "SELECT id, title, description, status, created_at, updated_at FROM todo WHERE id = ?"
	row := r.DB.QueryRow(SQL, id)
	var todo entity.Todo
	var createdAtStr, updatedAtStr sql.NullString

	if err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}

	// Parse string to time.Time only if it's not null
	if createdAtStr.Valid {
		if createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			todo.CreatedAt = createdAt
		} else {
			return nil, err
		}
	} else {
		todo.CreatedAt = time.Time{}
	}

	if updatedAtStr.Valid {
		if updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
			todo.UpdatedAt = updatedAt
		} else {
			return nil, err
		}
	} else {
		todo.UpdatedAt = time.Time{}
	}

	return &todo, nil
}

// Create implements TodoRepository
func (r *todoRepository) Create(todo *entity.Todo) (*entity.Todo, error) {
	SQL := "INSERT INTO todo (id, title, description, status, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(SQL, todo.Id, todo.Title, todo.Description, todo.Status, todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Update implements TodoRepository
func (r *todoRepository) Update(id string, todo *entity.Todo) (*entity.Todo, error) {
	SQL := "UPDATE todo SET title = ?, description = ?, status = ?, updated_at = ? WHERE id = ?"
	_, err := r.DB.Exec(SQL, todo.Title, todo.Description, todo.Status, todo.UpdatedAt, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Delete implements TodoRepository
func (r *todoRepository) Delete(id string) error {
	SQL := "DELETE FROM todo WHERE id = ?"
	_, err := r.DB.Exec(SQL, id)
	if err != nil {
		return err
	}
	return nil
}
