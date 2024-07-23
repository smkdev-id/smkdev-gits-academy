package usecase

import (
	"EkoEdyPurwanto/task-2/model"
	"EkoEdyPurwanto/task-2/model/dto/req"
	"EkoEdyPurwanto/task-2/repository/postgres"
	"EkoEdyPurwanto/task-2/utility/common"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	TodoListUseCase interface {
		Create(payload req.CreateRequest) error
		GetAll() (model.TodoLists, error)
		GetByPayload(payload string) (model.TodoLists, error)
	}
	todoListUseCase struct {
		Repository postgres.TodoListRepository
		Validate   *validator.Validate
		Logger     *logrus.Logger
	}
)

// GetByPayload implements TodoListUseCase.
func (t *todoListUseCase) GetByPayload(payload string) (model.TodoLists, error) {
	todolists, err := t.Repository.FindByIdentifier(payload)
	if err != nil {
		return nil, err
	}

	return todolists, nil
}

// Create implements TodoListUseCase.
func (t *todoListUseCase) Create(payload req.CreateRequest) error {
	// struct validation
	err := t.Validate.Struct(payload)
	if err != nil {
		t.Logger.Warnf("Invalid request body : %+v", err)
		return err
	}

	// fill value of todolist
	todolist := model.TodoList{
		Id:          common.GenerateID(),
		Title:       payload.Title,
		Description: payload.Description,
		Status:      model.StatusPending,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:   sql.NullTime{Time: time.Time{}, Valid: false},
		DeletedAt:   sql.NullTime{Time: time.Time{}, Valid: false},
	}

	// save todolist
	err = t.Repository.Save(todolist)
	if err != nil {
		t.Logger.Warnf("failed save todolist : %+v", err)
		return err
	}
	return nil
}

// GetAll implements TodoListUseCase.
func (t *todoListUseCase) GetAll() (model.TodoLists, error) {
	todolists, err := t.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return todolists, nil
}

// NewTodoListUseCase Constructor
func NewTodoListUseCase(repository postgres.TodoListRepository, validate *validator.Validate, logger *logrus.Logger) TodoListUseCase {
	return &todoListUseCase{
		Repository: repository,
		Validate:   validate,
		Logger:     logger,
	}
}
