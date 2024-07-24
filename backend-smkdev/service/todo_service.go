package service

import (
	"context"
	"todos/entity"
	"todos/repository"
)

// TodoServiceInteface mendefinisikan kontrak untuk layanan Todo
type TodoServiceInteface interface {
	GetAll(c context.Context) ([]*entity.Todo, error)
	GetByID(c context.Context, id int) (*entity.Todo, error)
	Create(c context.Context, todo *entity.Todo) error
	Update(c context.Context, todo *entity.Todo, id int) error
	Delete(c context.Context, id int) error
}

// TodoService adalah implementasi dari TodoServiceInteface
type TodoService struct {
	repo repository.TodoRepositoryInterface
}

// NewTodoService menginisialisasi instance baru dari TodoService
func NewTodoService(repo repository.TodoRepositoryInterface) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

// GetAll memanggil metode GetAll dari repository untuk mendapatkan semua todo
func (s *TodoService) GetAll(c context.Context) ([]*entity.Todo, error) {
	return s.repo.GetAll(c)
}

// GetByID memanggil metode GetByID dari repository untuk mendapatkan todo berdasarkan ID
func (s *TodoService) GetByID(c context.Context, id int) (*entity.Todo, error) {
	return s.repo.GetByID(c, id)
}

// Create memanggil metode Create dari repository untuk membuat todo baru
func (s *TodoService) Create(c context.Context, todo *entity.Todo) error {
	return s.repo.Create(c, todo)
}

// Update memanggil metode Update dari repository untuk memperbarui todo berdasarkan ID
func (s *TodoService) Update(c context.Context, todo *entity.Todo, id int) error {
	todo.ID = id // Set ID todo sebelum memperbaruinya
	return s.repo.Update(c, todo)
}

// Delete memanggil metode Delete dari repository untuk menghapus todo berdasarkan ID
func (s *TodoService) Delete(c context.Context, id int) error {
	return s.repo.Delete(c, id)
}
