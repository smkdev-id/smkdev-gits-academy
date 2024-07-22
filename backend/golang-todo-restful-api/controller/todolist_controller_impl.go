package controller

import (
	"net/http"
	"strconv"

	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/helper"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/web"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/service"

	"github.com/julienschmidt/httprouter"
)

type TodoListControllerImpl struct {
	TodoListService service.TodoListService
}

func NewTodoListController(todoListService service.TodoListService) TodoListController {
	return &TodoListControllerImpl{TodoListService: todoListService}
}

func (controller *TodoListControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListCreateRequest := web.TodoListCreateRequest{}
	helper.ReadFromRequestBody(request, &todoListCreateRequest)

	todoListResponse := controller.TodoListService.Create(request.Context(), todoListCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponse,
	}

	helper.WriteToResonseBody(writer, webResponse)

}

func (controller *TodoListControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todoListUpdateRequest := web.TodoListUpdateRequest{}
	helper.ReadFromRequestBody(request, &todoListUpdateRequest)

	TodoListId := params.ByName("todolistId")
	id, err := strconv.Atoi(TodoListId)
	helper.PanicIfError(err)

	todoListUpdateRequest.Id = id

	todoListResponse := controller.TodoListService.Update(request.Context(), todoListUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponse,
	}

	helper.WriteToResonseBody(writer, webResponse)

}

func (controller *TodoListControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todoListId := params.ByName("todolistId")
	id, err := strconv.Atoi(todoListId)
	helper.PanicIfError(err)

	controller.TodoListService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResonseBody(writer, webResponse)

}

func (controller *TodoListControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListId := params.ByName("todolistId")
	id, err := strconv.Atoi(todoListId)
	helper.PanicIfError(err)

	todoListResponse := controller.TodoListService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponse,
	}

	helper.WriteToResonseBody(writer, webResponse)

}

func (controller *TodoListControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListResponses := controller.TodoListService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponses,
	}

	helper.WriteToResonseBody(writer, webResponse)

}
