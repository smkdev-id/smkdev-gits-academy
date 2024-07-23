package helper

import (
	"todos/dto"
)

type ResponseWithData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(params dto.ResponseParams) any {
	var response any
	var status string

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "Success"
	} else {
		status = "Failed"
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}
	return response
}
