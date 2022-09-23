package controller

import (
	"chat-service/controller/group"
	"chat-service/controller/message"
	"chat-service/controller/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	group.Controller.Routes(r)
	message.Controller.Routes(r)
	user.Controller.Routes(r)
}
