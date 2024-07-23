package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
}

func TestFindAllTodoRepositoryWhenSuccess(t *testing.T) {
	db, mock, closeFunc := newMockDB(t)
	defer closeFunc()

	repo := repository.NewTodoRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "is_completed", "created_at", "updated_at"}).
		AddRow("1", "Todo 1", "Description 1", false, "2022-01-01 00:00:00", nil).
		AddRow("2", "Todo 2", "Description 2", true, "2024-07-22 18:04:28", "2024-07-22 18:04:28")

	mock.ExpectQuery("SELECT id, title, description, is_completed, created_at, updated_at FROM todo").WillReturnRows(rows)

	todos, err := repo.FindAll()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, *todos, 2)
	assert.Equal(t, "Todo 1", (*todos)[0].Title)
	assert.Equal(t, "Todo 2", (*todos)[1].Title)
}
