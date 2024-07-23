package service

import (
	"context"
	"todos/entity"
	"todos/repository"
)

type TodoServiceInteface interface {
	GetAll(c context.Context) ([]*entity.Todo, error)
	GetByID(c context.Context, id int) (*entity.Todo, error)
	Create(c context.Context, todo *entity.Todo) error
	Update(c context.Context, todo *entity.Todo, id int) error
	Delete(c context.Context, id int) error
}

type TodoService struct{
	repo repository.TodoRepositoryInterface
}

func NewTodoService(repo repository.TodoRepositoryInterface) *TodoService{
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetAll(c context.Context) ([]*entity.Todo, error) {
	return s.repo.GetAll(c)
}	

func (s *TodoService) GetByID(c context.Context, id int) (*entity.Todo, error) {
	return s.repo.GetByID(c, id)
}

func (s *TodoService) Create(c context.Context, todo *entity.Todo) error {
	return s.repo.Create(c, todo)
}	

func (s *TodoService) Update(c context.Context, todo *entity.Todo, id int) error {
	todo.ID = id
	return s.repo.Update(c, todo)
}

func (s *TodoService) Delete(c context.Context, id int) error {
	return s.repo.Delete(c, id)
}