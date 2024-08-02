package response

type (
	CommonResponse struct {
		StatusCode int                `json:"status_code"`
		Message    string             `json:"message"`
		Data       interface{}        `json:"data"`
		Pagination PaginationResponse `json:"pagination"`
	}

	PaginationResponse struct {
		CurrentPage int `json:"current_page"`
		PageSize    int `json:"page_size"`
		TotalData   int `json:"total_data"`
		TotalPages  int `json:"total_pages"`
	}
)
