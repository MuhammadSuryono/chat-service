package tables

import (
	"time"
)

type Participant struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	GroupChatId int64     `json:"group_chat_id"`
	UserId      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
