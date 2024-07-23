package repository

import (
	"todo-list/model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetByID(id string) (model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Update(todo model.Todo) (model.Todo, error)
	Delete(id string) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (t *todoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	err := t.db.Find(&todos).Error
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (t *todoRepository) GetByID(id string) (model.Todo, error) {
	var todo model.Todo

	err := t.db.Where("id = ?", id).First(&todo).Error

	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (t *todoRepository) Create(todo model.Todo) (model.Todo, error) {
	err := t.db.Create(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (t *todoRepository) Update(todo model.Todo) (model.Todo, error) {
	err := t.db.Save(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (t *todoRepository) Delete(id string) error {
	var todo model.Todo
	err := t.db.Where("id = ?", id).Delete(&todo).Error

	if err != nil {
		return err
	}

	return nil
}
