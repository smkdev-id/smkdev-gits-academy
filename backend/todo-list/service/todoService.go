package service

import (
	"time"
	"todo-list/helper"
	"todo-list/model"
	"todo-list/repository"

	"github.com/gofrs/uuid"
)

type TodoService interface {
	GetAllTodo() ([]model.Todo, error)
	GetTodoByID(id string) (model.Todo, error)
	CreateTodo(input helper.Input) (model.Todo, error)
	UpdateTodo(id string, input helper.Input) (model.Todo, error)
	DeleteTodo(id string) error
}

type todoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) *todoService {
	return &todoService{repository}
}

func (s todoService) GetAllTodo() ([]model.Todo, error) {
	todos, err := s.repository.GetAll()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s todoService) GetTodoByID(id string) (model.Todo, error) {
	todo, err := s.repository.GetByID(id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s todoService) CreateTodo(input helper.Input) (model.Todo, error) {
	var todo model.Todo
	id, _ := uuid.NewV4()
	todo.ID = id.String()
	todo.Title = input.Title
	todo.Description = input.Description
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	todo, err := s.repository.Create(todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s todoService) UpdateTodo(id string, input helper.Input) (model.Todo, error) {
	todo, err := s.repository.GetByID(id)
	if err != nil {
		return todo, err
	}

	todo.Title = input.Title
	todo.Description = input.Description
	todo.UpdatedAt = time.Now()

	newTodo, err := s.repository.Update(todo)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil
}

func (s todoService) DeleteTodo(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
