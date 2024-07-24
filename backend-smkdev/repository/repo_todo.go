package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"todos/entity"
)

type TodoRepositoryInterface interface {
	GetAll(c context.Context) ([]*entity.Todo, error)
	GetByID(c context.Context, id int) (*entity.Todo, error)
	Create(c context.Context, todo *entity.Todo) error
	Update(c context.Context, todo *entity.Todo) error
	Delete(c context.Context, id int) error
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) GetAll(c context.Context) ([]*entity.Todo, error) {
	todos := make([]*entity.Todo, 0)
	// Buat query SQL untuk mengambil semua todo
	rows, err := r.db.Query("SELECT id, title, description, completed, created_at, updated_at FROM todos")
	if err != nil {
		return nil, errors.New("todo not found")
	}
	defer rows.Close()

	// Baca setiap baris dan tambahkan ke slice todos
	for rows.Next() {
		todo := &entity.Todo{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, errors.New("todo not found")
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

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

func (r *TodoRepository) Create(c context.Context, todo *entity.Todo) error {
	// Buat query SQL untuk menyimpan todo baru
	_, err := r.db.Exec("INSERT INTO todos (title, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", todo.Title, todo.Description, todo.Completed, time.Now(), time.Now())
	return err
}

func (r *TodoRepository) Update(c context.Context, todo *entity.Todo) error {
	// Buat query SQL untuk memperbarui todo
	_, err := r.db.Exec("UPDATE todos SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?", todo.Title, todo.Description, todo.Completed, time.Now(), todo.ID)
	return err
}

func (r *TodoRepository) Delete(c context.Context, id int) error {
	// Buat query SQL untuk menghapus todo
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

// func (r *TodoRepository) GetAll(c context.Context) ([]*entity.Todo, error) {
// 	todos := make([]*entity.Todo, 0)
// 	err := r.db.WithContext(c).Find(&todos).Error
// 	return todos, err
// }

// func (r *TodoRepository) GetByID(c context.Context, id int) (*entity.Todo, error) {
// 	todo := new(entity.Todo)
// 	err := r.db.WithContext(c).First(&todo, id).Error
// 	return todo, err
// }

// func (r *TodoRepository) Create(c context.Context, todo *entity.Todo) error {
// 	err := r.db.WithContext(c).Create(&todo).Error
// 	return err
// }

// func (r *TodoRepository) Update(c context.Context, todo *entity.Todo) error {
// 	err := r.db.WithContext(c).Model(&entity.Todo{}).Where("id = ?", todo.ID).Updates(&todo).Error
// 	return err
// }

// func (r *TodoRepository) Delete(c context.Context, id int) error {
// 	return r.db.WithContext(c).Delete(&entity.Todo{}, id).Error
// }
