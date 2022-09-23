package group

import (
	"chat-service/config/database"
	"chat-service/models/tables"
)

type group struct {
}

var Repository group

func (g group) FindAll() (groups []tables.GroupChat) {
	database.Connection.Find(&groups)
	return
}

func (g group) ReadGroup(id int64) (value tables.GroupChat) {
	database.Connection.Where("id = ?", id).First(&value)
	return
}

func (g group) GroupActive() (value tables.GroupChat) {
	database.Connection.Last(&value)
	return
}
