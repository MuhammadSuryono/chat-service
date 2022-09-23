package message

import (
	"chat-service/config/database"
	"chat-service/models/tables"
	"fmt"
)

type message struct {
}

var Repository message

func (g message) LastMessageGroup(groupId int64) (msg tables.Message) {
	database.Connection.Where("group_id = ?", groupId).Last(&msg)
	return
}

func (g message) MessageGroup(groupId int64, limit int, page int) (msg []tables.Message) {
	offset := (page * limit) - limit
	fmt.Println("Offset", offset, page, limit)
	database.Connection.Where("group_id = ?", groupId).Debug().Limit(limit).Offset(offset).Order("created_at ASC").Find(&msg)
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
