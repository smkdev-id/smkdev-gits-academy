package web

type TodoListUpdateRequest struct {
	Id          int    `validate:"required"`
	Description string `validate:"required,min=1,max=200" json:"description"`
	Completed   bool   `json:"completed"`
}
