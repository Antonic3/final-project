package util

import "final-project/model"

func CreateResponse(isSuccess bool, data any, errorMessage string, message string) model.Response {
	return model.Response{
		Success: isSuccess,
		Data:    data,
		Error:   errorMessage,
		Message: message,
	}
}
