package user

import (
	"chat-service/requests"
	"github.com/gin-gonic/gin"
)

func (c controller) Routes(r *gin.RouterGroup) {
	g := r.Group("user")
	{
		g.POST("create", requests.ValidateRequest, c.saveUser)
	}
}
