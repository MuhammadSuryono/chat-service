package message

import (
	"chat-service/config/database"
	"chat-service/models/tables"
	"chat-service/system"
	"fmt"
)

type message struct {
}

var Repository message

func (g message) LastMessageGroup(groupId int64) (msg tables.Message) {
	database.Connection.Where("group_id = ?", groupId).Last(&msg)
	return
}

func (g message) MessageGroup(groupId int64, limit int, page int, lastDate string) (msg []tables.Message, total int64) {
	offset := (page * limit) - limit
	fmt.Println("Offset", offset, page, limit)
	query := database.Connection.Where("group_id = ?", groupId).Debug().Offset(offset).Limit(limit).Order("created_at DESC")
	if lastDate == "" {
		lastDate = system.TimeNowString()
	}
	if lastDate != "" {
		query = query.Where("created_at <= ?", lastDate)
	} else {
		query = query.Where("DATE(created_at) <= ?", system.TimeNowString())
	}

	query.Find(&msg)
	query.Count(&total)

	for i, j := 0, len(msg)-1; i < j; i, j = i+1, j-1 {
		msg[i], msg[j] = msg[j], msg[i]
	}
	return
}

func (g message) TotalMessageGroup(groupId int64) (total int64) {
	database.Connection.Where("group_id = ?", groupId).Model(&tables.Message{}).Count(&total)
	return
}

func (g message) SaveNewMessageChat(message string, groupId, sender int64) *tables.Message {
	msg := &tables.Message{
		GroupId:     groupId,
		SenderId:    sender,
		MessageType: "chat",
		Message:     message,
	}

	database.Connection.Create(&msg)
	return msg
}
