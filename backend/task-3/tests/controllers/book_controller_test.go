package controllers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alwiirfan/internal/dto/request"
	"github.com/alwiirfan/internal/dto/response"
	"github.com/alwiirfan/internal/http/controllers"
	"github.com/alwiirfan/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BookControllerSuite struct {
	suite.Suite
	controller      *controllers.BookController
	mockBookService *mocks.BookService
	echoInstance    *echo.Echo
}

func (s *BookControllerSuite) SetupTest() {
	s.mockBookService = new(mocks.BookService)
	s.controller = controllers.NewBookController(s.mockBookService)
	s.echoInstance = echo.New()
}

func TestBookController(t *testing.T) {
	suite.Run(t, new(BookControllerSuite))
}

func (s *BookControllerSuite) TestCreateNewBookControllerWhenSuccess() {
	// Dummy request body
	dummyReq := request.CreateBookRequest{
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             1000,
	}

	s.mockBookService.On("Create", context.Background(), &dummyReq).Return(nil)

	body, _ := json.Marshal(dummyReq)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPost, "/api/v1/books", readReqBody),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Create(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusCreated, resp.StatusCode)
	require.Equal(s.T(), "Create book successfully", resp.Message)
}

func (s *BookControllerSuite) TestCreateNewBookControllerWhenFailedInvalidURL() {
	// Dummy request body
	dummyReq := request.CreateBookRequest{
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             1000,
	}

	s.mockBookService.On("Create", context.Background(), &dummyReq).Return(nil)

	body, _ := json.Marshal(dummyReq)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books?error=error", readReqBody),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Create(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)
	require.Equal(s.T(), "Failed to create book", resp.Message)
}

func (s *BookControllerSuite) TestCreateNewBookControllerWhenFailedValidation() {
	// Dummy request body
	dummyReq := request.CreateBookRequest{
		Title:             "",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             1000,
	}

	var errValidate validator.ValidationErrors
	s.mockBookService.On("Create", context.Background(), &dummyReq).Return(errValidate)

	body, _ := json.Marshal(dummyReq)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPost, "/api/v1/books", readReqBody),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Create(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)
	require.Equal(s.T(), "Validation failed", resp.Message)
}

func (s *BookControllerSuite) TestCreateNewBookControllerWhenFailedBadRequest() {
	// dummy request
	dummyReq := map[string]interface{}{
		"title":       true,
		"description": "Test Description",
		"author":      "Test Author",
		"year":        2023,
		"stock":       10,
		"price":       1000,
	}

	s.mockBookService.On("Create", context.Background(), &dummyReq).Return(nil)

	body, _ := json.Marshal(dummyReq)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPost, "/api/v1/books", readReqBody),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Create(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)
	require.Equal(s.T(), "Failed to create book", resp.Message)
}

func (s *BookControllerSuite) TestCreateNewBookControllerFailedConnection() {
	dummyReq := request.CreateBookRequest{
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             1000,
	}

	s.mockBookService.On("Create", context.Background(), &dummyReq).Return(sql.ErrConnDone)

	body, _ := json.Marshal(dummyReq)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPost, "/api/v1/books", readReqBody),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Create(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	require.Equal(s.T(), "Failed to create book", resp.Message)
}

func (s *BookControllerSuite) TestFindAllBookControllerWhenSuccess() {

	s.mockBookService.On("FindAll", context.Background(), 1, 10).Return([]*response.BookResponse{
		{
			Id:                "1",
			Isbn:              "1234567890",
			Title:             "Test Book",
			Description:       "Test Description",
			Author:            "Test Author",
			YearOfManufacture: 2023,
			IsDisplayed:       true,
			Stock:             10,
			Price:             1000,
			CreatedAt:         time.Now().Local().String(),
			UpdatedAt:         time.Now().Local().String(),
		},
		{
			Id:                "2",
			Isbn:              "1234567890",
			Title:             "Test Book",
			Description:       "Test Description",
			Author:            "Test Author",
			YearOfManufacture: 2023,
			IsDisplayed:       true,
			Stock:             10,
			Price:             1000,
			CreatedAt:         time.Now().Local().String(),
			UpdatedAt:         time.Now().Local().String(),
		},
	}, int(2), nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books", nil),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindAll(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Find all books successfully", resp.Message)
}

func (s *BookControllerSuite) TestFindAllBookControllerWhenFailedConnection() {

	s.mockBookService.On("FindAll", context.Background(), 1, 10).Return(nil, 0, sql.ErrConnDone)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books", nil),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindAll(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	require.Equal(s.T(), "Failed to find all books", resp.Message)
}

func (s *BookControllerSuite) TestFindAllSearchBooksControllerWhenSuccess() {
	dummyreq := request.SearchBookRequest{}
	s.mockBookService.On("FindAllSearch", context.Background(), 1, 10, &dummyreq).Return([]*response.BookResponse{
		{
			Id:                "1",
			Isbn:              "1234567890",
			Title:             "Test Book",
			Description:       "Test Description",
			Author:            "Test Author",
			YearOfManufacture: 2023,
			IsDisplayed:       true,
			Stock:             10,
			Price:             1000,
			CreatedAt:         time.Now().Local().String(),
			UpdatedAt:         time.Now().Local().String(),
		},
		{
			Id:                "2",
			Isbn:              "1234567890",
			Title:             "Test Book",
			Description:       "Test Description",
			Author:            "Test Author",
			YearOfManufacture: 2023,
			IsDisplayed:       true,
			Stock:             10,
			Price:             1000,
			CreatedAt:         time.Now().Local().String(),
			UpdatedAt:         time.Now().Local().String(),
		},
	}, int(2), nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/search", nil),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindAllSearch(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Find all books successfully", resp.Message)
}

func (s *BookControllerSuite) TestFindAllSearchBooksControllerWhenFailedQueryParams() {

	dummyreq := request.SearchBookRequest{}

	s.mockBookService.On("FindAllSearch", context.Background(), 1, 10, &dummyreq).Return(nil, 0, nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/search?error=error", nil),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindAllSearch(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)
	require.Equal(s.T(), "Invalid query params: error", resp.Message)
}

func (s *BookControllerSuite) TestFindAllSearchBooksControllerWhenFailedConnection() {

	dummyreq := request.SearchBookRequest{}

	s.mockBookService.On("FindAllSearch", context.Background(), 1, 10, &dummyreq).Return(nil, 0, sql.ErrConnDone)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/search", nil),
		rec,
	)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindAllSearch(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	require.Equal(s.T(), "Failed to find all books", resp.Message)
}

func (s *BookControllerSuite) TestFindByIDBookControllerWhenSuccess() {

	s.mockBookService.On("FindByID", context.Background(), "1").Return(&response.BookResponse{
		Id:                "1",
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}, nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/1", nil),
		rec,
	)
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindByID(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Find book by id successfully", resp.Message)
}

func (s *BookControllerSuite) TestFindByIDBookControllerWhenFailedNotFound() {

	s.mockBookService.On("FindByID", context.Background(), "1").Return(nil, sql.ErrNoRows)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/1", nil),
		rec,
	)
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindByID(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusNotFound, resp.StatusCode)
	require.Equal(s.T(), "Failed to find book by id, book not found", resp.Message)
}

func (s *BookControllerSuite) TestFindByISBNBookControllerWhenSuccess() {

	s.mockBookService.On("FindByISBN", context.Background(), "1234567890").Return(&response.BookResponse{
		Id:                "1",
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}, nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/1234567890", nil),
		rec,
	)
	c.SetParamNames("isbn")
	c.SetParamValues("1234567890")

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindByISBN(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Find book by isbn successfully", resp.Message)
}

func (s *BookControllerSuite) TestFindByISBNBookControllerWhenFailedNotFound() {

	s.mockBookService.On("FindByISBN", context.Background(), "1234567890").Return(nil, sql.ErrNoRows)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodGet, "/api/v1/books/1234567890", nil),
		rec,
	)
	c.SetParamNames("isbn")
	c.SetParamValues("1234567890")

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.FindByISBN(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusNotFound, resp.StatusCode)
	require.Equal(s.T(), "Failed to find book by isbn, book not found", resp.Message)
}

func (s *BookControllerSuite) TestUpdateBookControllerWhenSuccess() {
	// dommy book
	id := "1"

	// dummy request
	dummyRequest := request.UpdateBookRequest{
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       false,
		Stock:             10,
		Price:             1000,
	}

	// dummy response
	dummyResponse := response.BookResponse{
		Id:                id,
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}
	s.mockBookService.On("FindByID", context.Background(), id).Return(&dummyResponse, nil)
	s.mockBookService.On("Update", context.Background(), id, &dummyRequest).Return(nil)

	body, _ := json.Marshal(dummyRequest)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPut, "/api/v1/books/1", readReqBody),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Update(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Update book successfully", resp.Message)
}

func (s *BookControllerSuite) TestUpdateBookControllerWhenFailedNotFound() {

	id := "1"

	s.mockBookService.On("FindByID", context.Background(), id).Return(nil, sql.ErrNoRows)
	s.mockBookService.On("Update", context.Background(), id, &request.UpdateBookRequest{}).Return(sql.ErrNoRows)

	body, _ := json.Marshal(request.UpdateBookRequest{})
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPut, "/api/v1/books/1", readReqBody),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Update(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusNotFound, resp.StatusCode)
	require.Equal(s.T(), "Failed to update book, book not found", resp.Message)
}

func (s *BookControllerSuite) TestUpdateBookControllerWhenFailedValidationError() {
	// dommy book
	id := "1"

	// dummy request
	dummyRequest := request.UpdateBookRequest{
		Title:             "",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       false,
		Stock:             10,
		Price:             1000,
	}

	// dummy response
	dummyResponse := response.BookResponse{
		Id:                id,
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}

	var errValidate validator.ValidationErrors

	s.mockBookService.On("FindByID", context.Background(), id).Return(&dummyResponse, nil)
	s.mockBookService.On("Update", context.Background(), id, &dummyRequest).Return(errValidate)

	body, _ := json.Marshal(dummyRequest)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPut, "/api/v1/books/1", readReqBody),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Update(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusBadRequest, resp.StatusCode)
	require.Equal(s.T(), "Failed to update book", resp.Message)
}

func (s *BookControllerSuite) TestUpdateBookControllerWhenFailedConnection() {

	id := "1"

	// dummy request
	dummyRequest := request.UpdateBookRequest{
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       false,
		Stock:             10,
		Price:             1000,
	}

	// dummy response
	dummyResponse := response.BookResponse{
		Id:                id,
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}

	s.mockBookService.On("FindByID", context.Background(), id).Return(&dummyResponse, nil)
	s.mockBookService.On("Update", context.Background(), id, &dummyRequest).Return(sql.ErrConnDone)

	body, _ := json.Marshal(dummyRequest)
	readReqBody := bytes.NewBuffer(body)
	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodPut, "/api/v1/books/1", readReqBody),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	c.Request().Header.Set("Content-Type", "application/json")

	err := s.controller.Update(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	require.Equal(s.T(), "Failed to update book", resp.Message)
}

func (s *BookControllerSuite) TestDeletByIDControllerWhenSuccess() {

	id := "1"

	// dummy response
	dummyResponse := response.BookResponse{
		Id:                id,
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}

	s.mockBookService.On("FindByID", context.Background(), id).Return(&dummyResponse, nil)
	s.mockBookService.On("DeleteByID", context.Background(), id).Return(nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodDelete, "/api/v1/books/1", nil),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	err := s.controller.DeleteByID(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Delete book successfully", resp.Message)

}

func (s *BookControllerSuite) TestDeletByIDControllerWhenFailedNotFound() {

	id := "1"

	s.mockBookService.On("FindByID", context.Background(), id).Return(nil, sql.ErrNoRows)
	s.mockBookService.On("DeleteByID", context.Background(), id).Return(sql.ErrNoRows)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodDelete, "/api/v1/books/1", nil),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	err := s.controller.DeleteByID(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusNotFound, resp.StatusCode)
	require.Equal(s.T(), "Failed to delete book, book not found", resp.Message)
}

func (s *BookControllerSuite) TestDeletByIDControllerWhenFailedConnnection() {

	id := "1"

	// dummy response
	dummyResponse := response.BookResponse{
		Id:                id,
		Isbn:              "1234567890",
		Title:             "Test Book",
		Description:       "Test Description",
		Author:            "Test Author",
		YearOfManufacture: 2023,
		IsDisplayed:       true,
		Stock:             10,
		Price:             1000,
		CreatedAt:         time.Now().Local().String(),
		UpdatedAt:         time.Now().Local().String(),
	}

	s.mockBookService.On("FindByID", context.Background(), id).Return(&dummyResponse, nil)
	s.mockBookService.On("DeleteByID", context.Background(), id).Return(sql.ErrConnDone)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodDelete, "/api/v1/books/1", nil),
		rec,
	)

	c.SetParamNames("id")
	c.SetParamValues(id)

	err := s.controller.DeleteByID(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	require.Equal(s.T(), "Failed to delete book", resp.Message)
}

func (s *BookControllerSuite) TestDeleteAllControllerWhenSuccess() {

	s.mockBookService.On("DeleteAll", context.Background()).Return(nil)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodDelete, "/api/v1/books", nil),
		rec,
	)

	err := s.controller.DeleteAll(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.CommonResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusOK, resp.StatusCode)
	require.Equal(s.T(), "Delete all books successfully", resp.Message)
}

func (s *BookControllerSuite) TestDeleteAllControllerWhenFailedConnnection() {

	s.mockBookService.On("DeleteAll", context.Background()).Return(sql.ErrConnDone)

	rec := httptest.NewRecorder()
	c := s.echoInstance.NewContext(
		httptest.NewRequest(http.MethodDelete, "/api/v1/books", nil),
		rec,
	)

	err := s.controller.DeleteAll(c)
	if err != nil {
		s.T().Fatal(err)
	}

	require.NoError(s.T(), err)

	var resp response.ErrorResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		s.T().Fatal(err)
	}

	require.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	require.Equal(s.T(), "Failed to delete all books", resp.Message)
}
