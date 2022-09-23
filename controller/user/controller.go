package user

import (
	"chat-service/Exception"
	"chat-service/requests"
	"chat-service/service/user"
	"github.com/gin-gonic/gin"
)

type controller struct {
}

var Controller controller

func (c controller) saveUser(ctx *gin.Context) {
	defer Exception.GetErrorJson(ctx)
	dataBind, isExists := ctx.Get("data")
	if !isExists {
		panic("Error binding data request.")
	}

	request := dataBind.(requests.AddUserRequest)
	user.Service.CreateUser(request)
}
