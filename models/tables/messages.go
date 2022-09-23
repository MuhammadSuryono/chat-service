package tables

import (
	"chat-service/config/database"
	"chat-service/system"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id            int64     `json:"id"`
	GroupId       int64     `json:"conversation_id"`
	SenderId      int64     `json:"sender_id"`
	Sender        User      `json:"sender" gorm:"-"`
	MessageType   string    `json:"message_type"`
	Message       string    `json:"message"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	CreateMessage string    `json:"create_message"`
}

func (a *Message) AfterFind(tx *gorm.DB) error {
	a.CreateMessage = system.TimeClock(a.CreatedAt)
	return database.Connection.Model(&User{}).Where("id = ?", a.SenderId).First(&(a.Sender)).Error
}

func (a *Message) AfterCreate(tx *gorm.DB) error {
	return database.Connection.Model(&User{}).Where("id = ?", a.SenderId).First(&(a.Sender)).Error
}
