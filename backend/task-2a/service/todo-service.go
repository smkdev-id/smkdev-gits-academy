package service

import (
	"time"

	"github.com/alwi09/dto/request"
	"github.com/alwi09/dto/response"
	"github.com/alwi09/entity"
	"github.com/alwi09/repository"
	"github.com/alwi09/util"
	"github.com/go-playground/validator/v10"
)

type (
	TodoService interface {
		FindAll() (*[]response.TodoResponse, error)
		FindById(id string) (*response.TodoResponse, error)
		Create(todoReq *request.TodoCreateRequest) (*response.TodoResponse, error)
		Update(id string, todoReq *request.TodoUpdateRequest) (*response.TodoResponse, error)
		Delete(id string) error
	}

	todoService struct {
		todoRepository repository.TodoRepository
		validate       *validator.Validate
	}
)

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo, validator.New()}
}

func (s *todoService) FindAll() (*[]response.TodoResponse, error) {
	todos, err := s.todoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var dataList []response.TodoResponse
	for _, todo := range *todos {
		var updatedAt string
		if todo.UpdatedAt != nil {
			updatedAt = todo.UpdatedAt.Local().String()
		} else {
			updatedAt = ""
		}

		dataList = append(dataList, response.TodoResponse{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
			IsCompleted: todo.IsCompleted,
			CreatedAt:   todo.CreatedAt.Local().String(),
			UpdatedAt:   updatedAt,
		})
	}

	return &dataList, nil
}

func (s *todoService) FindById(id string) (*response.TodoResponse, error) {
	todo, err := s.todoRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	var updatedAt string
	if todo.UpdatedAt != nil {
		updatedAt = todo.UpdatedAt.Local().String()
	} else {
		updatedAt = ""
	}

	data := response.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt.Local().String(),
		UpdatedAt:   updatedAt,
	}

	return &data, nil
}

func (s *todoService) Create(todoReq *request.TodoCreateRequest) (*response.TodoResponse, error) {
	if err := s.validate.Struct(todoReq); err != nil {
		return nil, err
	}

	newTodo := entity.Todo{
		Id:          util.GenerateUUID(),
		Title:       todoReq.Title,
		Description: todoReq.Description,
		IsCompleted: false,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	todo, err := s.todoRepository.Create(&newTodo)
	if err != nil {
		return nil, err
	}

	// Convert time.Time to string
	var updatedAt string
	if todo.UpdatedAt != nil {
		updatedAt = todo.UpdatedAt.Local().String()
	} else {
		updatedAt = ""
	}

	data := response.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt.Local().String(),
		UpdatedAt:   updatedAt,
	}

	return &data, nil
}

func (s *todoService) Update(id string, todoReq *request.TodoUpdateRequest) (*response.TodoResponse, error) {
	if err := s.validate.Struct(todoReq); err != nil {
		return nil, err
	}

	todo, err := s.todoRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	todo.Title = todoReq.Title
	todo.Description = todoReq.Description
	todo.IsCompleted = todoReq.IsCompleted
	updatedAt := time.Now().Local()
	todo.UpdatedAt = &updatedAt

	todoRes, err := s.todoRepository.Update(id, todo)
	if err != nil {
		return nil, err
	}

	data := response.TodoResponse{
		Id:          todoRes.Id,
		Title:       todoRes.Title,
		Description: todoRes.Description,
		IsCompleted: todoRes.IsCompleted,
		CreatedAt:   todoRes.CreatedAt.Local().String(),
		UpdatedAt:   todoRes.UpdatedAt.Local().String(),
	}

	return &data, nil
}

func (s *todoService) Delete(id string) error {
	existingTodo, err := s.todoRepository.FindById(id)
	if err != nil {
		return err
	}

	err = s.todoRepository.Delete(existingTodo.Id)
	if err != nil {
		return err
	}

	return nil
}
