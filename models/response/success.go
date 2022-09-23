package response

import "chat-service/models"

func SuccessResponse(created bool, message string, data interface{}) models.CommonResponse {
	var code = 200
	if created {
		code = 201
	}
	return models.CommonResponse{
		Code:      code,
		IsSuccess: true,
		Message:   message,
		Data:      data,
	}
}
