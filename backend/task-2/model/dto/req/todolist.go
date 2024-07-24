package req

import (
	"EkoEdyPurwanto/task-2/model"
	"time"
)

type (
	CreateRequest struct {
		Id          string       `json:"id"`
		Title       string       `json:"title" validate:"required"`
		Description string       `json:"description" validate:"max=100"`
		Status      model.Status `json:"status"`
		CreatedAt   time.Time    `json:"created_at"`
		UpdatedAt   time.Time    `json:"updated_at"`
		DeletedAt   time.Time    `json:"deleted_at"`
	}

	UpdateRequest struct {
		Id        string       `json:"id" validate:"required"`
		Status    model.Status `json:"status" validate:"required"`
		UpdatedAt time.Time    `json:"updated_at"`
	}

	DeleteRequest struct {
		Id string `json:"id" validate:"required"`
	}
)
