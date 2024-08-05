package helper


type Response struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    any `json:"data"`
}

func ResponseJSON(code int, message string, data any) Response {
	responseJSON := Response{
		Code: code,
		Message: message,
		Data: data,
	}

	return responseJSON
}

