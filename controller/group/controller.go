package group

import (
	"chat-service/service/group"
	"github.com/gin-gonic/gin"
	"strconv"
)

type controller struct {
}

var Controller controller

func (c controller) getAllGroup(ctx *gin.Context) {
	group.Service.GetAllGroupChat()
}

func (c controller) readGroup(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	group.Service.ReadGroupChat(int64(id))
}

func (c controller) readGroupActive(ctx *gin.Context) {
	group.Service.GroupChatActive()
}
