package repository

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alwi09/entity"
	"github.com/alwi09/repository"
	"github.com/stretchr/testify/assert"
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	closeFunc := func() {
		db.Close()
	}

	return db, mock, closeFunc
}

func TestRunTodoRepository(t *testing.T) {
	t.Run("FindAllTodoRepositoryWhenSuccess", TestFindAllTodoRepositoryWhenSuccess)
	t.Run("FindAllTodoRepositoryWhenFailedConnection", TestFindAllTodoRepositoryWhenFailedConnection)
	t.Run("FindByIdTodoRepositoryWhenSuccess", TestFindByIdTodoRepositoryWhenSuccess)
	t.Run("FindByIdTodoRepositoryWhenFailedNotFound", TestFindByIdTodoRepositoryWhenFailedNotFound)
	t.Run("FindByIdTodoRepositoryWhenFailedFailedConnection", TestFindByIdTodoRepositoryWhenFailedConnection)
	t.Run("CreateTodoRepositoryWhenSuccess", TestCreateTodoRepositoryWhenSuccess)
	t.Run("CreateTodoRepositoryWhenFailedConnection", TestCreateTodoRepositoryWhenFailedConnection)
	t.Run("UpdateTodoRepositoryWhenSuccess", TestUpdateTodoRepositoryWhenSuccess)
	t.Run("UpdateTodoRepositoryWhenFailedNotFound", TestUpdateTodoRepositoryWhenFailedNotFound)
	t.Run("UpdateTodoRepositoryWhenFailedFailedConnection", TestUpdateTodoRepositoryWhenFailedConnection)
	t.Run("DeleteTodoRepositoryWhenSuccess", TestDeleteTodoRepositoryWhenSuccess)
	t.Run("DeleteTodoRepositoryWhenFailedNotFound", TestDeleteTodoRepositoryWhenFailedNotFound)
	t.Run("DeleteTodoRepositoryWhenFailedFailed", TestDeleteTodoRepositoryWhenFailedConnection)
}

func TestFindAllTodoRepositoryWhenSuccess(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	createdAt := time.Now().Local().Format(time.DateTime)
	updatedAt := time.Now().Local().Format(time.DateTime)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "is_completed", "created_at", "updated_at"}).
		AddRow("1", "Todo 1", "Description 1", false, createdAt, nil).
		AddRow("2", "Todo 2", "Description 2", true, createdAt, updatedAt)

	mock.ExpectQuery("SELECT id, title, description, is_completed, created_at, updated_at FROM todo").WillReturnRows(rows)

	todos, err := repo.FindAll()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, *todos, 2)
	assert.Equal(t, "Todo 1", (*todos)[0].Title)
	assert.Equal(t, "Todo 2", (*todos)[1].Title)
	assert.Equal(t, "Description 1", (*todos)[0].Description)
	assert.Equal(t, "Description 2", (*todos)[1].Description)
	assert.Equal(t, false, (*todos)[0].IsCompleted)
	assert.Equal(t, true, (*todos)[1].IsCompleted)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindAllTodoRepositoryWhenFailedConnection(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	mock.ExpectQuery("SELECT id, title, description, is_completed, created_at, updated_at FROM todo").WillReturnError(sql.ErrConnDone)

	todos, err := repo.FindAll()
	assert.Error(t, err)
	assert.Nil(t, todos)
}

func TestFindByIdTodoRepositoryWhenSuccess(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	createdAt := time.Now().Local().Format(time.DateTime)
	updatedAt := time.Now().Local().Format(time.DateTime)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "is_completed", "created_at", "updated_at"}).
		AddRow("1", "Todo 1", "Description 1", false, createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, is_completed, created_at, updated_at FROM todo WHERE id = ?")).
		WillReturnRows(rows)

	todo, err := repo.FindById("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Todo 1", todo.Title)
	assert.Equal(t, "Description 1", todo.Description)
	assert.Equal(t, false, todo.IsCompleted)
}

func TestFindByIdTodoRepositoryWhenFailedNotFound(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, is_completed, created_at, updated_at FROM todo WHERE id = ?")).
		WithArgs("1").WillReturnError(sql.ErrNoRows)

	todo, err := repo.FindById("1")
	assert.Error(t, err)
	assert.Nil(t, todo)
}

func TestFindByIdTodoRepositoryWhenFailedConnection(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, is_completed, created_at, updated_at FROM todo WHERE id = ?")).
		WithArgs("1").WillReturnError(sql.ErrConnDone)

	todo, err := repo.FindById("1")
	assert.Error(t, err)
	assert.Nil(t, todo)
}

func TestCreateTodoRepositoryWhenSuccess(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	todo := &entity.Todo{
		Id:          "1",
		Title:       "Todo 1",
		IsCompleted: false,
		Description: "Description 1",
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	mock.ExpectExec("INSERT INTO todo").
		WithArgs(todo.Id, todo.Title, todo.Description, todo.IsCompleted, todo.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdTodo, err := repo.Create(todo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, todo.Id, createdTodo.Id)
	assert.Equal(t, todo.Title, createdTodo.Title)
	assert.Equal(t, todo.Description, createdTodo.Description)
	assert.Equal(t, todo.IsCompleted, createdTodo.IsCompleted)
	assert.Equal(t, todo.CreatedAt, createdTodo.CreatedAt)
	assert.Equal(t, todo.UpdatedAt, createdTodo.UpdatedAt)
}

func TestCreateTodoRepositoryWhenFailedConnection(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	todo := &entity.Todo{
		Id:          "1",
		Title:       "Todo 1",
		IsCompleted: false,
		Description: "Description 1",
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   nil,
	}

	mock.ExpectExec("INSERT INTO todo").
		WithArgs(todo.Id, todo.Title, todo.Description, todo.IsCompleted, todo.CreatedAt).
		WillReturnError(sql.ErrConnDone)

	createdTodo, err := repo.Create(todo)
	assert.Error(t, err)
	assert.Nil(t, createdTodo)
}

func TestUpdateTodoRepositoryWhenSuccess(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	updatedAt := time.Now().Local()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todo SET title = ?, description = ?, is_completed = ?, updated_at = ? WHERE id = ?")).
		WithArgs("updated title", "updated description", true, &updatedAt, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	todo := &entity.Todo{
		Id:          "1",
		Title:       "updated title",
		Description: "updated description",
		IsCompleted: true,
		UpdatedAt:   &updatedAt,
	}

	updatedTodo, err := repo.Update(todo.Id, todo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, updatedTodo)
	assert.Equal(t, todo.Id, updatedTodo.Id)
	assert.Equal(t, todo.Title, updatedTodo.Title)
	assert.Equal(t, todo.Description, updatedTodo.Description)
	assert.Equal(t, todo.IsCompleted, updatedTodo.IsCompleted)
	assert.Equal(t, todo.UpdatedAt, updatedTodo.UpdatedAt)
}

func TestUpdateTodoRepositoryWhenFailedNotFound(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	updatedAt := time.Now().Local()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todo SET title = ?, description = ?, is_completed = ?, updated_at = ? WHERE id = ?")).
		WithArgs("updated title", "updated description", true, &updatedAt, "1").
		WillReturnError(sql.ErrNoRows)

	todo := &entity.Todo{
		Id:          "1",
		Title:       "updated title",
		Description: "updated description",
		IsCompleted: true,
		UpdatedAt:   &updatedAt,
	}

	updatedTodo, err := repo.Update(todo.Id, todo)
	assert.Error(t, err)
	assert.Nil(t, updatedTodo)
}

func TestUpdateTodoRepositoryWhenFailedConnection(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	updatedAt := time.Now().Local()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todo SET title = ?, description = ?, is_completed = ?, updated_at = ? WHERE id = ?")).
		WithArgs("updated title", "updated description", true, &updatedAt, "1").
		WillReturnError(sql.ErrConnDone)

	todo := &entity.Todo{
		Id:          "1",
		Title:       "updated title",
		Description: "updated description",
		IsCompleted: true,
		UpdatedAt:   &updatedAt,
	}

	updatedTodo, err := repo.Update(todo.Id, todo)
	assert.Error(t, err)
	assert.Nil(t, updatedTodo)
}

func TestDeleteTodoRepositoryWhenSuccess(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("1")
	assert.NoError(t, err)
}

func TestDeleteTodoRepositoryWhenFailedNotFound(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
		WithArgs("1").
		WillReturnError(sql.ErrNoRows)

	err := repo.Delete("1")
	assert.Error(t, err)
}

func TestDeleteTodoRepositoryWhenFailedConnection(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
		WithArgs("1").
		WillReturnError(sql.ErrConnDone)

	err := repo.Delete("1")
	assert.Error(t, err)
}
