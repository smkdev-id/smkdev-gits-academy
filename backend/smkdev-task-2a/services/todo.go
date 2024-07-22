package services

import (
	"go-todo/models"

	"gorm.io/gorm"
)

type TodoService struct {
    DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
    return &TodoService{DB: db}
}

func (s *TodoService) FetchTodos() ([]models.Todo, error) {
    var todos []models.Todo
    if err := s.DB.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

func (s *TodoService) CreateTodos(todo *models.Todo) error {
    return s.DB.Create(todo).Error
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
    var todo models.Todo
    if err := s.DB.First(&todo, id).Error; err != nil {
        return nil, err
    }
    return &todo, nil
}

func (s *TodoService) UpdateTodos(todo *models.Todo) error {
    return s.DB.Save(todo).Error
}

func (s *TodoService) DeleteTodos(todo *models.Todo) error {
    return s.DB.Delete(todo).Error
}
