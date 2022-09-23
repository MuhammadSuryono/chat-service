package requests

import (
	validate2 "chat-service/helpers/validate"
	"github.com/gin-gonic/gin"
)

type AddUserRequest struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

func ValidateRequest(c *gin.Context) {
	var request AddUserRequest
	_ = c.BindJSON(&request)
	validate := validate2.NewValidate()
	validate.ValidationStruct(request).JsonResponse(c)
}
