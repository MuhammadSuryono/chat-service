package tables

import (
	"chat-service/config/database"
	"chat-service/system"
	"fmt"
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
	LastDateData  string    `json:"last_data_date" gorm:"-"`
}

func (a *Message) AfterFind(tx *gorm.DB) error {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		_ = fmt.Sprintf("Error timezone %v", err)
	}
	_ = fmt.Sprintf("Location %v", loc)
	a.CreateMessage = system.TimeClock(a.CreatedAt.In(loc))

	a.LastDateData = fmt.Sprintf("%d-%2d-%2d %2d:%2d:%2d", a.CreatedAt.Year(), a.CreatedAt.Month(), a.CreatedAt.Day(), a.CreatedAt.Hour(), a.CreatedAt.Minute(), a.CreatedAt.Second())
	return database.Connection.Model(&User{}).Where("id = ?", a.SenderId).First(&(a.Sender)).Error
}

func (a *Message) AfterCreate(tx *gorm.DB) error {
	return database.Connection.Model(&User{}).Where("id = ?", a.SenderId).First(&(a.Sender)).Error
}
