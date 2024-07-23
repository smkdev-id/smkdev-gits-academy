package req

import (
	"EkoEdyPurwanto/task-2/model"
	"database/sql"
)

type (
	CreateRequest struct {
		Id          string       `json:"id"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedAt   sql.NullTime `json:"created_at"`
		UpdatedAt   sql.NullTime `json:"updated_at"`
		DeletedAt   sql.NullTime `json:"deleted_at"`
	}

	UpdateRequest struct {
		Id     string       `json:"id"`
		Status model.Status `json:"status"`
	}
)
