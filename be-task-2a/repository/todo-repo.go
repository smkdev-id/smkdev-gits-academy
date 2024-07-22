package repository

import (
	"context"
	"github.com/HerlyRyan/be-task-2a/model"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) CreateTodo(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *todoRepository) GetTodoByID(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.WithContext(ctx).First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := r.db.WithContext(ctx).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) UpdateTodo(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *todoRepository) DeleteTodo(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}
