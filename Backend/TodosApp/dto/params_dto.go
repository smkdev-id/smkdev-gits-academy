package dto

//Struct Untuk response params
type ResponseParams struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
