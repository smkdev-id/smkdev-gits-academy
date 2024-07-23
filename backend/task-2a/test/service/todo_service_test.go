package service

import (
	"database/sql"
	"testing"
	"time"

	"github.com/alwi09/dto/request"
	"github.com/alwi09/entity"
	"github.com/alwi09/mocks"
	"github.com/alwi09/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRunTodoService(t *testing.T) {
	t.Run("FindAllTodoServiceWhenSuccess", TestFindAllTodoServiceWhenSuccess)
	t.Run("FindAllTodoServiceWhenFailedConnection", TestFindAllTodoServiceWhenFailedConnection)
	t.Run("FindByIdTodoServiceWhenSuccess", TestFindByIdTodoServiceWhenSuccess)
	t.Run("FindByIdTodoServiceWhenFailedNotFound", TestFindByIdTodoServiceWhenFailedNotFound)
	t.Run("FindByIdTodoServiceWhenFailedConnection", TestFindByIdTodoServiceWhenFailedConnection)
	t.Run("CreateTodoServiceWhenSuccess", TestCreateTodoServiceWhenSuccess)
	t.Run("CreateTodoServiceWhenFailedValidation", TestCreateTodoServiceWhenFailedValidation)
	t.Run("CreateTodoServiceWhenFailedConnection", TestCreateTodoServiceWhenFailedConnection)
	t.Run("UpdateTodoServiceWhenSuccess", TestUpdateTodoServiceWhenSuccess)
	t.Run("UpdateTodoServiceWhenFailedNotFound", TestUpdateTodoServiceWhenFailedNotFound)
	t.Run("UpdateTodoServiceWhenFailedValidation", TestUpdateTodoServiceWhenFailedValidation)
	t.Run("UpdateTodoServiceWhenFailedConnection", TestUpdateTodoServiceWhenFailedConnection)
	t.Run("DeleteTodoServiceWhenSuccess", TestDeleteTodoServiceWhenSuccess)
	t.Run("DeleteTodoServiceWhenFailedNotFound", TestDeleteTodoServiceWhenFailedNotFound)
	t.Run("DeleteTodoServiceWhenFailedConnection", TestDeleteTodoServiceWhenFailedConnection)
}

func TestFindAllTodoServiceWhenSuccess(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	updatedAt := time.Now().Local()

	dummyTodos := []entity.Todo{
		{
			Id:          "1",
			Title:       "Todo 1",
			Description: "Description 1",
			IsCompleted: false,
			CreatedAt:   time.Now().Local(),
			UpdatedAt:   nil,
		},
		{
			Id:          "2",
			Title:       "Todo 2",
			Description: "Description 2",
			IsCompleted: true,
			CreatedAt:   time.Now().Local(),
			UpdatedAt:   &updatedAt,
		},
	}

	mockRepo.On("FindAll").Return(&dummyTodos, nil)

	result, err := mockService.FindAll()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, *result, 2)
	assert.Equal(t, "Todo 1", (*result)[0].Title)
	assert.Equal(t, "Todo 2", (*result)[1].Title)
	assert.Equal(t, "Description 1", (*result)[0].Description)
	assert.Equal(t, "Description 2", (*result)[1].Description)
	assert.Equal(t, false, (*result)[0].IsCompleted)
	assert.Equal(t, true, (*result)[1].IsCompleted)
}

func TestFindAllTodoServiceWhenFailedConnection(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	mockRepo.On("FindAll").Return(nil, sql.ErrConnDone)

	_, err := mockService.FindAll()
	assert.Error(t, err)
}

func TestFindByIdTodoServiceWhenSuccess(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	dummyTodo := entity.Todo{
		Id:          "1",
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	mockRepo.On("FindById", "1").Return(&dummyTodo, nil)

	result, err := mockService.FindById("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Todo 1", result.Title)
	assert.Equal(t, "Description 1", result.Description)
	assert.Equal(t, false, result.IsCompleted)
	assert.Equal(t, dummyTodo.CreatedAt.Local().String(), result.CreatedAt)
	assert.Equal(t, "", result.UpdatedAt)
}

func TestFindByIdTodoServiceWhenFailedNotFound(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	mockRepo.On("FindById", "1").Return(nil, sql.ErrNoRows)

	_, err := mockService.FindById("1")
	assert.Error(t, err)
}

func TestFindByIdTodoServiceWhenFailedConnection(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	mockRepo.On("FindById", "1").Return(nil, sql.ErrConnDone)

	_, err := mockService.FindById("1")
	assert.Error(t, err)
}

func TestCreateTodoServiceWhenSuccess(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	dummyRequest := request.TodoCreateRequest{
		Title:       "Todo 1",
		Description: "Description 1",
	}

	dummyTodo := entity.Todo{
		Id:          "1",
		Title:       dummyRequest.Title,
		Description: dummyRequest.Description,
		IsCompleted: false,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(&dummyTodo, nil)

	result, err := mockService.Create(&dummyRequest)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Todo 1", result.Title)
	assert.Equal(t, "Description 1", result.Description)
	assert.Equal(t, false, result.IsCompleted)
	assert.Equal(t, dummyTodo.CreatedAt.Local().String(), result.CreatedAt)
	assert.Equal(t, "", result.UpdatedAt)
}

func TestCreateTodoServiceWhenFailedValidation(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)

	mockService := service.NewTodoService(mockRepo)

	dummyRequest := request.TodoCreateRequest{
		Title:       "",
		Description: "Description 1",
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(nil, nil)

	result, err := mockService.Create(&dummyRequest)
	assert.Error(t, err)
	assert.Nil(t, result)

	assert.Equal(t, err.Error(), err.Error())
}

func TestCreateTodoServiceWhenFailedConnection(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	dummyRequest := request.TodoCreateRequest{
		Title:       "Todo 1",
		Description: "Description 1",
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(nil, sql.ErrConnDone)

	_, err := mockService.Create(&dummyRequest)
	assert.Error(t, err)
}

func TestUpdateTodoServiceWhenSuccess(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "Todo 1 updated",
		Description: "Description 1 updated",
		IsCompleted: true,
	}

	// dummy existing todo
	dummyExistingTodo := entity.Todo{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	updatedAt := time.Now().Local()

	// dummy todo updated
	dummyTodo := entity.Todo{
		Id:          id,
		Title:       dummyRequest.Title,
		Description: dummyRequest.Description,
		IsCompleted: dummyRequest.IsCompleted,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   &updatedAt,
	}

	mockRepo.On("FindById", id).Return(&dummyExistingTodo, nil)
	mockRepo.On("Update", id, mock.AnythingOfType("*entity.Todo")).Return(&dummyTodo, nil)

	result, err := mockService.Update(id, &dummyRequest)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Todo 1 updated", result.Title)
	assert.Equal(t, "Description 1 updated", result.Description)
	assert.Equal(t, true, result.IsCompleted)
	assert.Equal(t, dummyTodo.CreatedAt.Local().String(), result.CreatedAt)
	assert.Equal(t, dummyTodo.UpdatedAt.Local().String(), result.UpdatedAt)
}

func TestUpdateTodoServiceWhenFailedNotFound(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "Todo 1 updated",
		Description: "Description 1 updated",
		IsCompleted: true,
	}

	mockRepo.On("FindById", id).Return(nil, sql.ErrNoRows)

	_, err := mockService.Update(id, &dummyRequest)
	assert.Error(t, err)
}

func TestUpdateTodoServiceWhenFailedValidation(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "",
		Description: "Description 1 updated",
		IsCompleted: true,
	}

	mockRepo.On("FindById", id).Return(nil, nil)
	mockRepo.On("Update", id, mock.AnythingOfType("*entity.Todo")).Return(nil, nil)

	_, err := mockService.Update(id, &dummyRequest)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), err.Error())
}

func TestUpdateTodoServiceWhenFailedConnection(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "Todo 1 updated",
		Description: "Description 1 updated",
		IsCompleted: true,
	}

	// dummy existing todo
	dummyExistingTodo := entity.Todo{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	mockRepo.On("FindById", id).Return(&dummyExistingTodo, nil)
	mockRepo.On("Update", id, mock.AnythingOfType("*entity.Todo")).Return(nil, sql.ErrConnDone)

	_, err := mockService.Update(id, &dummyRequest)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), err.Error())
}

func TestDeleteTodoServiceWhenSuccess(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	// dummy existing todo
	dummyExistingTodo := entity.Todo{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	mockRepo.On("FindById", id).Return(&dummyExistingTodo, nil)
	mockRepo.On("Delete", id).Return(nil)

	err := mockService.Delete(id)
	assert.NoError(t, err)
}

func TestDeleteTodoServiceWhenFailedNotFound(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	mockRepo.On("FindById", id).Return(nil, sql.ErrNoRows)
	mockRepo.On("Delete", id).Return(sql.ErrNoRows)

	err := mockService.Delete(id)
	assert.Error(t, err)
}

func TestDeleteTodoServiceWhenFailedConnection(t *testing.T) {
	mockRepo := new(mocks.TodoRepository)
	mockService := service.NewTodoService(mockRepo)

	id := "1"

	mockRepo.On("FindById", id).Return(nil, sql.ErrConnDone)
	mockRepo.On("Delete", id).Return(sql.ErrConnDone)

	err := mockService.Delete(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), err.Error())
}
