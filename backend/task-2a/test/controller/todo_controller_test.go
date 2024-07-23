package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alwi09/controller"
	"github.com/alwi09/dto/request"
	"github.com/alwi09/dto/response"
	"github.com/alwi09/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRunTodoController(t *testing.T) {
	t.Run("TestFindAllTodoControllerWhenSuccess", TestFindAllTodoControllerWhenSuccess)
	t.Run("TestFindAllTodoControllerWhenFailedConnection", TestFindAllTodoControllerWhenFailedConnection)
	t.Run("TestFindByIdTodoControllerWhenSuccess", TestFindByIdTodoControllerWhenSuccess)
	t.Run("TestFindByIdTodoControllerWhenFailedNotFound", TestFindByIdTodoControllerWhenFailedNotFound)
	t.Run("TestCreateTodoControllerWhenSuccess", TestCreateTodoControllerWhenSuccess)
	t.Run("TestCreateTodoControllerWhenFailedValidation", TestCreateTodoControllerWhenFailedValidation)
	t.Run("TestCreateTodoControllerWhenFailedBadRequest", TestCreateTodoControllerWhenFailedBadRequest)
	t.Run("TestCreateTodoControllerWhenFailedConnection", TestCreateTodoControllerWhenFailedConnection)
	t.Run("TestUpdateTodoControllerWhenSuccess", TestUpdateTodoControllerWhenSuccess)
	t.Run("TestUpdateTodoControllerWhenFailedNotFound", TestUpdateTodoControllerWhenFailedNotFound)
	t.Run("TestUpdateTodoControllerWhenFailedValidation", TestUpdateTodoControllerWhenFailedValidation)
	t.Run("TestUpdateTodoControllerWhenFailedBadRequest", TestUpdateTodoControllerWhenFailedBadRequest)
	t.Run("TestUpdateTodoControllerWhenFailedConnection", TestUpdateTodoControllerWhenFailedConnection)
	t.Run("TestDeleteTodoControllerWhenSuccess", TestDeleteTodoControllerWhenSuccess)
	t.Run("TestDeleteTodoControllerWhenFailedNotFound", TestDeleteTodoControllerWhenFailedNotFound)
	t.Run("TestDeleteTodoControllerWhenFailedConnection", TestDeleteTodoControllerWhenFailedConnection)
}

func TestFindAllTodoControllerWhenSuccess(t *testing.T) {
	mockService := new(mocks.TodoService)
	mockService.On("FindAll").Return(&[]response.TodoResponse{}, nil)

	mockController := controller.NewTodoController(mockService)

	req, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.FindAll(rw, req)

	res := rw.Result()

	var respTodos response.CommonResponse[[]response.TodoResponse]
	err = json.NewDecoder(res.Body).Decode(&respTodos)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, []response.TodoResponse{}, respTodos.Data)
	assert.Equal(t, rw.Header().Get("Content-Type"), "application/json")
	assert.Equal(t, "Find all todos successfully", respTodos.Message)
}

func TestFindAllTodoControllerWhenFailedConnection(t *testing.T) {
	mockService := new(mocks.TodoService)
	mockService.On("FindAll").Return(nil, sql.ErrConnDone)

	mockController := controller.NewTodoController(mockService)

	req, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.FindAll(rw, req)

	res := rw.Result()

	var respTodos response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodos)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Equal(t, "Failed to find all todos", respTodos.Message)
	assert.Equal(t, "Check internet connection", respTodos.Error)
}

func TestFindByIdTodoControllerWhenSuccess(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy todo response
	dummyTodoResponse := response.TodoResponse{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
	}

	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)

	mockController := controller.NewTodoController(mockService)

	req, err := http.NewRequest("GET", "/api/v1/todos/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.FindById(rw, req, id)

	res := rw.Result()

	var respTodo response.CommonResponse[response.TodoResponse]
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, respTodo.StatusCode)
	assert.Equal(t, "Find todo successfully", respTodo.Message)
	assert.Equal(t, dummyTodoResponse, respTodo.Data)
	assert.Equal(t, rw.Header().Get("Content-Type"), "application/json")
}

func TestFindByIdTodoControllerWhenFailedNotFound(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	mockService.On("FindById", id).Return(nil, sql.ErrNoRows)

	mockController := controller.NewTodoController(mockService)

	req, err := http.NewRequest("GET", "/api/v1/todos/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.FindById(rw, req, id)

	res := rw.Result()

	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, respTodo.StatusCode)
	assert.Equal(t, "Failed to find todo", respTodo.Message)
	assert.Equal(t, "Todo not found", respTodo.Error)
}

func TestCreateTodoControllerWhenSuccess(t *testing.T) {
	mockService := new(mocks.TodoService)

	dummyRequest := request.TodoCreateRequest{
		Title:       "Todo 1",
		Description: "Description 1",
	}

	// dummy todo response
	dummyTodoResponse := response.TodoResponse{
		Id:          "1",
		Title:       dummyRequest.Title,
		Description: dummyRequest.Description,
		IsCompleted: false,
		CreatedAt:   time.Now().Local().String(),
	}

	mockService.On("Create", &dummyRequest).Return(&dummyTodoResponse, nil)

	mockController := controller.NewTodoController(mockService)

	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/todos", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.Create(rw, req)

	res := rw.Result()

	var respTodo response.CommonResponse[response.TodoResponse]
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, respTodo.StatusCode)
	assert.Equal(t, "Create todo successfully", respTodo.Message)
	assert.Equal(t, dummyTodoResponse, respTodo.Data)
	assert.Equal(t, rw.Header().Get("Content-Type"), "application/json")
}

func TestCreateTodoControllerWhenFailedValidation(t *testing.T) {
	mockService := new(mocks.TodoService)

	dummyRequest := request.TodoCreateRequest{
		Title:       "",
		Description: "Description 1",
	}

	var errValidate validator.ValidationErrors
	mockService.On("Create", &dummyRequest).Return(nil, errValidate)

	mockController := controller.NewTodoController(mockService)

	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/todos", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.Create(rw, req)

	res := rw.Result()

	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, respTodo.StatusCode)
	assert.Equal(t, "Failed to create todo", respTodo.Message)
	assert.Equal(t, "Validation error", respTodo.Error)
}

func TestCreateTodoControllerWhenFailedBadRequest(t *testing.T) {
	mockService := new(mocks.TodoService)

	dummyRequest := map[string]interface{}{
		"title":       true,
		"description": "Description 1",
	}

	mockService.On("Create", &dummyRequest).Return(nil, nil)

	mockController := controller.NewTodoController(mockService)

	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/todos", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.Create(rw, req)

	res := rw.Result()

	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, respTodo.StatusCode)
	assert.Equal(t, "Failed to create todo", respTodo.Message)
	assert.Equal(t, "Invalid request body", respTodo.Error)
}

func TestCreateTodoControllerWhenFailedConnection(t *testing.T) {
	mockService := new(mocks.TodoService)

	dummyRequest := request.TodoCreateRequest{
		Title:       "Todo 1",
		Description: "Description 1",
	}

	mockService.On("Create", &dummyRequest).Return(nil, sql.ErrConnDone)

	mockController := controller.NewTodoController(mockService)

	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/todos", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.Create(rw, req)

	res := rw.Result()

	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusInternalServerError, respTodo.StatusCode)
	assert.Equal(t, "Failed to create todo", respTodo.Message)
	assert.Equal(t, "Check internet connection", respTodo.Error)
}

func TestUpdateTodoControllerWhenSuccess(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: true,
	}

	// dummy response
	dummyTodoResponse := response.TodoResponse{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	// mock service
	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)
	mockService.On("Update", id, &dummyRequest).Return(&dummyTodoResponse, nil)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request body
	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	// request with body and id path variable
	req, err := http.NewRequest("PUT", "/api/v1/todos/1", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Update(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.CommonResponse[response.TodoResponse]
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, respTodo.StatusCode)
	assert.Equal(t, "Update todo successfully", respTodo.Message)
	assert.Equal(t, dummyTodoResponse, respTodo.Data)
}

func TestUpdateTodoControllerWhenFailedNotFound(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	mockService.On("FindById", id).Return(nil, sql.ErrNoRows)
	mockService.On("Update", id, nil).Return(nil, sql.ErrNoRows)

	mockController := controller.NewTodoController(mockService)

	req, err := http.NewRequest("PUT", "/api/v1/todos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	mockController.Update(rw, req, id)

	res := rw.Result()

	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, respTodo.StatusCode)
	assert.Equal(t, "Failed to update todo", respTodo.Message)
	assert.Equal(t, "Todo not found", respTodo.Error)
}

func TestUpdateTodoControllerWhenFailedValidation(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "",
		Description: "Description 1",
		IsCompleted: true,
	}

	// dummy response
	dummyTodoResponse := response.TodoResponse{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	var errValidate validator.ValidationErrors

	// mock service
	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)
	mockService.On("Update", id, &dummyRequest).Return(nil, errValidate)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request body
	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	// request with body and id path variable
	req, err := http.NewRequest("PUT", "/api/v1/todos/"+id, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Update(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, respTodo.StatusCode)
	assert.Equal(t, "Failed to update todo", respTodo.Message)
	assert.Equal(t, "Validation error", respTodo.Error)
}

func TestUpdateTodoControllerWhenFailedBadRequest(t *testing.T) {
	mockService := new(mocks.TodoService)

	dummyRequest := map[string]interface{}{
		"title":        true,
		"description":  "Description 1",
		"is_completed": true,
	}

	// dummy response
	dummyTodoResponse := response.TodoResponse{
		Id:          "1",
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	id := "1"

	// mock service
	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)
	mockService.On("Update", id, dummyRequest).Return(nil, nil)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request body
	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	// request with body and id path variable
	req, err := http.NewRequest("PUT", "/api/v1/todos/"+id, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Update(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, respTodo.StatusCode)
	assert.Equal(t, "Failed to update todo", respTodo.Message)
	assert.Equal(t, "Invalid request body", respTodo.Error)
}

func TestUpdateTodoControllerWhenFailedConnection(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy request
	dummyRequest := request.TodoUpdateRequest{
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: true,
	}

	// dummy response
	dummyTodoResponse := response.TodoResponse{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	// mock service
	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)
	mockService.On("Update", id, &dummyRequest).Return(nil, sql.ErrConnDone)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request body
	reqBody, err := json.Marshal(dummyRequest)
	if err != nil {
		t.Fatal(err)
	}

	// request with body and id path variable
	req, err := http.NewRequest("PUT", "/api/v1/todos/"+id, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Update(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusInternalServerError, respTodo.StatusCode)
	assert.Equal(t, "Failed to update todo", respTodo.Message)
	assert.Equal(t, "Check internet connection", respTodo.Error)
}

func TestDeleteTodoControllerWhenSuccess(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy response
	dummyTodoResponse := response.TodoResponse{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	// mock service
	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)
	mockService.On("Delete", id).Return(nil)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request with id path variable
	req, err := http.NewRequest("DELETE", "/api/v1/todos/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Delete(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.CommonResponse[response.TodoResponse]
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, respTodo.StatusCode)
	assert.Equal(t, "Delete todo successfully", respTodo.Message)
	assert.Equal(t, response.TodoResponse{}, respTodo.Data)
}

func TestDeleteTodoControllerWhenFailedNotFound(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy service
	mockService.On("FindById", id).Return(nil, sql.ErrNoRows)
	mockService.On("Delete", id).Return(nil)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request with id path variable
	req, err := http.NewRequest("DELETE", "/api/v1/todos/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Delete(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, respTodo.StatusCode)
	assert.Equal(t, "Failed to delete todo", respTodo.Message)
	assert.Equal(t, "Todo not found", respTodo.Error)
}

func TestDeleteTodoControllerWhenFailedConnection(t *testing.T) {
	mockService := new(mocks.TodoService)

	id := "1"

	// dummy response
	dummyTodoResponse := response.TodoResponse{
		Id:          id,
		Title:       "Todo 1",
		Description: "Description 1",
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	// dummy service
	mockService.On("FindById", id).Return(&dummyTodoResponse, nil)
	mockService.On("Delete", id).Return(sql.ErrConnDone)

	// mock controller
	mockController := controller.NewTodoController(mockService)

	// request with id path variable
	req, err := http.NewRequest("DELETE", "/api/v1/todos/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rw := httptest.NewRecorder()

	// call controller
	mockController.Delete(rw, req, id)

	// decode response
	res := rw.Result()

	// assert response
	var respTodo response.ErrorResponse
	err = json.NewDecoder(res.Body).Decode(&respTodo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusInternalServerError, respTodo.StatusCode)
	assert.Equal(t, "Failed to delete todo", respTodo.Message)
	assert.Equal(t, "Check internet connection", respTodo.Error)
}
