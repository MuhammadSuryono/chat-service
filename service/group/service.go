package group

import (
	"chat-service/models/dto"
	"chat-service/models/response"
	"chat-service/repository/group"
	"chat-service/repository/message"
	"chat-service/system"
)

type service struct {
}

var Service service

func (s service) GetAllGroupChat() {
	resp := group.Repository.FindAll()

	var groups []dto.GroupChatDto
	for _, chat := range resp {
		msg := message.Repository.LastMessageGroup(chat.Id)

		lastMessage := msg.Message
		if lastMessage == "" {
			lastMessage = "Tidak ada pesan"
		}

		groups = append(groups, dto.GroupChatDto{
			Id:        chat.Id,
			GroupName: chat.GroupName,
			ChannelID: chat.ChannelID,
			LastMessage: dto.LastMessageDto{
				Message:   lastMessage,
				CreatedAt: system.TimeClock(msg.CreatedAt),
			},
		})
	}
	system.Context.JSON(response.SUCCESS_CODE, response.SuccessResponse(false, "Success data", groups))
}

func (s service) ReadGroupChat(id int64) {
	resp := group.Repository.ReadGroup(id)
	system.Context.JSON(response.SUCCESS_CODE, response.SuccessResponse(false, "Success data", resp))
}

func (s service) GroupChatActive() {
	resp := group.Repository.GroupActive()
	system.Context.JSON(response.SUCCESS_CODE, response.SuccessResponse(false, "Success data", resp))
}
