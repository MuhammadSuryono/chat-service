package group

import "github.com/gin-gonic/gin"

func (c controller) Routes(r *gin.RouterGroup) {
	g := r.Group("group")
	{
		g.GET("all", c.getAllGroup)
		g.GET("read/:id", c.readGroup)
		g.GET("active", c.readGroupActive)
	}
}
