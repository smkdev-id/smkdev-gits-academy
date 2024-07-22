package web

type TodoListCreateRequest struct {
	Description string `validate:"required,min=1,max=200" json:"description"`
	Completed   bool   `json:"completed"`
}
