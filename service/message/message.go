package message

import (
	"chat-service/models/dto"
	"chat-service/models/response"
	"chat-service/repository/message"
	"chat-service/system"
	"math"
)

type msg struct {
}

var Service msg

func (receiver msg) GetMessageGroup(groupId int64) {
	var paginationRequest dto.PaginationDto
	_ = system.Context.BindQuery(&paginationRequest)

	if paginationRequest.Page == 0 {
		paginationRequest.Page = 1
	}

	if paginationRequest.Limit == 0 {
		paginationRequest.Limit = 50
	}

	msgs, totalMessage := message.Repository.MessageGroup(groupId, paginationRequest.Limit, paginationRequest.Page, paginationRequest.LastDataDate)

	var dtoMessages []dto.MessageDto
	for _, t := range msgs {
		dtoMessages = append(dtoMessages, dto.MessageDto{
			Id:              t.Id,
			Message:         t.Message,
			CreatedAt:       system.TimeClock(t.CreatedAt),
			SenderName:      t.Sender.Name,
			SenderUsername:  t.Sender.Email,
			SenderId:        t.SenderId,
			SenderPesantren: t.Sender.Pesantren,
			LastDateData:    t.LastDateData,
		})
	}

	system.Context.JSON(response.SUCCESS_CODE, response.SuccessResponse(false, "Success", dto.ResponseMessageDto{
		TotalMessage:         message.Repository.TotalMessageGroup(groupId),
		CurrentTotalResponse: int64(len(msgs)),
		LimitMessage:         paginationRequest.Limit,
		Page:                 paginationRequest.Page,
		TotalPage:            int(math.Ceil(float64(totalMessage) / float64(paginationRequest.Limit))),
		Data:                 dtoMessages,
	}))
}
