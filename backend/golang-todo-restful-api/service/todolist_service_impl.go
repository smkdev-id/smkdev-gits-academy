package service

import (
	"context"
	"database/sql"

	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/exception"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/helper"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/domain"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/web"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type TodoListServiceImpl struct {
	TodoListRepository repository.TodoListRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewTodoListService(todoListRepository repository.TodoListRepository, DB *sql.DB, validate *validator.Validate) TodoListService {
	return &TodoListServiceImpl{
		TodoListRepository: todoListRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *TodoListServiceImpl) Create(ctx context.Context, request web.TodoListCreateRequest) web.TodoListResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommiOrRollBack(tx)

	TodoList := domain.TodoList{
		Description: request.Description,
		Completed:   request.Completed,
	}

	TodoList = service.TodoListRepository.Save(ctx, tx, TodoList)

	return helper.TodoListResponse(TodoList)
}

func (service *TodoListServiceImpl) Update(ctx context.Context, request web.TodoListUpdateRequest) web.TodoListResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommiOrRollBack(tx)

	// check if TodoList is already exists
	todoList, err := service.TodoListRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todoList.Description = request.Description
	todoList.Completed = request.Completed

	todoList = service.TodoListRepository.Update(ctx, tx, todoList)

	return helper.TodoListResponse(todoList)
}

func (service *TodoListServiceImpl) Delete(ctx context.Context, TodoListId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommiOrRollBack(tx)

	// check if TodoList is already exists
	TodoList, err := service.TodoListRepository.FindById(ctx, tx, TodoListId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoListRepository.Delete(ctx, tx, TodoList)
}

func (service *TodoListServiceImpl) FindById(ctx context.Context, todoListId int) web.TodoListResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommiOrRollBack(tx)

	todoList, err := service.TodoListRepository.FindById(ctx, tx, todoListId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.TodoListResponse(todoList)
}

func (service *TodoListServiceImpl) FindAll(ctx context.Context) []web.TodoListResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommiOrRollBack(tx)

	todoList := service.TodoListRepository.FindAll(ctx, tx)

	return helper.TodoListResponses(todoList)

}
