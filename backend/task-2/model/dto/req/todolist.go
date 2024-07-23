package req

import (
	"EkoEdyPurwanto/task-2/model"
	"time"
)

type (
	CreateRequest struct {
		Id          string       `json:"id"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedAt   time.Time    `json:"created_at"`
		UpdatedAt   time.Time    `json:"updated_at"`
		DeletedAt   time.Time    `json:"deleted_at"`
	}

	UpdateRequest struct {
		Id        string       `json:"id"`
		Status    model.Status `json:"status"`
		UpdatedAt time.Time    `json:"updated_at"`
	}
)
