package message

import (
	"chat-service/Exception"
	"chat-service/helpers/ws"
	"chat-service/helpers/ws/wsmultiple"
	"chat-service/service/message"
	"chat-service/system"
	"github.com/gin-gonic/gin"
	"strconv"
)

type controller struct {
}

var Controller controller

func (c controller) getMessageGroup(ctx *gin.Context) {
	defer Exception.GetErrorJson(system.Context)
	groupIdString, _ := ctx.GetQuery("groupId")
	groupId, _ := strconv.Atoi(groupIdString)
	message.Service.GetMessageGroup(int64(groupId))
}

func (c controller) wsChatConnectionRoom(ctx *gin.Context) {
	channelID := ctx.Param("roomId")
	ws.ServeWs(ctx.Writer, ctx.Request, channelID, 0)
}

func (c controller) wsChatConnectionRoomGroup(ctx *gin.Context) {
	channelID := ctx.Param("roomId")
	wsmultiple.ServeWs(ctx.Writer, ctx.Request, channelID)
}
