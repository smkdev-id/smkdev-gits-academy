package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"todos/entity"
)

// TodoRepositoryInterface mendefinisikan kontrak untuk operasi CRUD pada entitas Todo
type TodoRepositoryInterface interface {
	GetAll(c context.Context) ([]*entity.Todo, error)
	GetByID(c context.Context, id int) (*entity.Todo, error)
	Create(c context.Context, todo *entity.Todo) error
	Update(c context.Context, todo *entity.Todo) error
	Delete(c context.Context, id int) error
}

// TodoRepository adalah implementasi dari TodoRepositoryInterface
type TodoRepository struct {
	db *sql.DB // Database connection
}

// NewTodoRepository membuat instance baru dari TodoRepository
func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

// GetAll mengambil semua entitas Todo dari database
func (r *TodoRepository) GetAll(c context.Context) ([]*entity.Todo, error) {
	todos := make([]*entity.Todo, 0)

	// Buat query SQL untuk mengambil semua todo
	rows, err := r.db.Query("SELECT id, title, description, completed, created_at, updated_at FROM todos")
	if err != nil {
		return nil, errors.New("todo not found")
	}
	defer rows.Close()

	// Iterasi melalui hasil query dan tambahkan ke slice todos
	for rows.Next() {
		todo := &entity.Todo{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, errors.New("todo not found")
		}
		todos = append(todos, todo)
	}

	// Periksa error dari iterasi rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByID mengambil satu entitas Todo berdasarkan ID dari database
func (r *TodoRepository) GetByID(c context.Context, id int) (*entity.Todo, error) {
	todo := &entity.Todo{}

	// Buat query SQL untuk mengambil todo berdasarkan ID
	row := r.db.QueryRow("SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = ?", id)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

// Create menambahkan entitas Todo baru ke dalam database
func (r *TodoRepository) Create(c context.Context, todo *entity.Todo) error {
	// Buat query SQL untuk menyimpan todo baru
	_, err := r.db.Exec("INSERT INTO todos (title, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", todo.Title, todo.Description, todo.Completed, time.Now(), time.Now())
	return err
}

// Update memperbarui entitas Todo yang sudah ada di dalam database
func (r *TodoRepository) Update(c context.Context, todo *entity.Todo) error {
	// Buat query SQL untuk memperbarui todo
	_, err := r.db.Exec("UPDATE todos SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?", todo.Title, todo.Description, todo.Completed, time.Now(), todo.ID)
	return err
}

// Delete menghapus entitas Todo dari database berdasarkan ID
func (r *TodoRepository) Delete(c context.Context, id int) error {
	// Buat query SQL untuk menghapus todo
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
