package usecase

import (
	"EkoEdyPurwanto/task-2/model"
	"EkoEdyPurwanto/task-2/model/dto/req"
	"EkoEdyPurwanto/task-2/repository/postgres"
	"EkoEdyPurwanto/task-2/utility/common"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	TodoListUseCase interface {
		Create(payload req.CreateRequest) error
		GetAll() (model.TodoLists, error)
		GetByPayload(payload string) (model.TodoLists, error)
		UpdateStatus(payload req.UpdateRequest) error
		Delete(payload string) error
	}
	todoListUseCase struct {
		Repository postgres.TodoListRepository
		Validate   *validator.Validate
		Logger     *logrus.Logger
	}
)

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
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
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

// GetByPayload implements TodoListUseCase.
func (t *todoListUseCase) GetByPayload(payload string) (model.TodoLists, error) {
	// validation
	if payload == "" {
		return nil, fmt.Errorf("id is required")
	}

	todolists, err := t.Repository.FindByIdentifier(payload)
	if err != nil {
		return nil, err
	}

	return todolists, nil
}

// UpdateStatus implements TodoListUseCase.
func (t *todoListUseCase) UpdateStatus(payload req.UpdateRequest) error {
	// struct validation
	err := t.Validate.Struct(payload)
	if err != nil {
		t.Logger.Warnf("Invalid request body : %+v", err)
		return err
	}

	payload.UpdatedAt = time.Now()

	if err := t.Repository.Update(payload); err != nil {
		return err
	}
	return nil
}

// Delete implements TodoListUseCase.
func (t *todoListUseCase) Delete(payload string) error {
	// validation
	if payload == "" {
		return fmt.Errorf("id is required")
	}

	if err := t.Repository.Delete(payload); err != nil {
		return err
	}
	return nil
}

// NewTodoListUseCase Constructor
func NewTodoListUseCase(repository postgres.TodoListRepository, validate *validator.Validate, logger *logrus.Logger) TodoListUseCase {
	return &todoListUseCase{
		Repository: repository,
		Validate:   validate,
		Logger:     logger,
	}
}
