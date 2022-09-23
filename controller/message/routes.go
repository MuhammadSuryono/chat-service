package message

import "github.com/gin-gonic/gin"

func (c controller) Routes(r *gin.RouterGroup) {
	g := r.Group("message")
	{
		g.GET("group", c.getMessageGroup)
		g.GET("chat/:roomId", c.wsChatConnectionRoom)
	}
}
