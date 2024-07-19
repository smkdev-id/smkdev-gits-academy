package request

type (
	TodoCreateRequest struct {
		Title       string `json:"title" validate:"required,min=3,max=100"`
		Description string `json:"description" validate:"required,min=5,max=500"`
	}

	TodoUpdateRequest struct {
		Title       string `json:"title" validate:"required,min=3,max=100"`
		Description string `json:"description" validate:"required,min=5,max=500"`
		Status      bool   `json:"status"`
	}
)
