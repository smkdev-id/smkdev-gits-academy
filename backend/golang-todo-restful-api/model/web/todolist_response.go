package web

type TodoListResponse struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
