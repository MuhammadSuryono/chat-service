package message

import (
	"chat-service/models/dto"
	"chat-service/models/response"
	"chat-service/repository/message"
	"chat-service/system"
	"fmt"
	"math"
)

type msg struct {
}

var Service msg

func (receiver msg) GetMessageGroup(groupId int64) {
	var paginationRequest dto.PaginationDto
	system.Context.BindQuery(&paginationRequest)

	if paginationRequest.Page == 0 {
		paginationRequest.Page = 1
	}

	if paginationRequest.Limit == 0 {
		paginationRequest.Limit = 50
	}

	msgs := message.Repository.MessageGroup(groupId, paginationRequest.Limit, paginationRequest.Page)

	var dtoMessages []dto.MessageDto
	for _, t := range msgs {
		dtoMessages = append(dtoMessages, dto.MessageDto{
			Id:             t.Id,
			Message:        t.Message,
			CreatedAt:      system.TimeClock(t.CreatedAt),
			SenderName:     fmt.Sprintf("%s %s", t.Sender.FirstName, t.Sender.LastName),
			SenderUsername: t.Sender.Username,
			SenderId:       t.SenderId,
		})
	}

	totalMessage := message.Repository.TotalMessageGroup(groupId)
	totalPage := math.Round(float64(totalMessage / int64(paginationRequest.Limit)))
	if totalPage <= 1 {
		totalPage = 1
	}

	system.Context.JSON(response.SUCCESS_CODE, response.SuccessResponse(false, "Success", dto.ResponseMessageDto{
		TotalMessage:         message.Repository.TotalMessageGroup(groupId),
		CurrentTotalResponse: int64(len(msgs)),
		LimitMessage:         paginationRequest.Limit,
		Page:                 paginationRequest.Page,
		TotalPage:            int(totalPage),
		Data:                 dtoMessages,
	}))
}
