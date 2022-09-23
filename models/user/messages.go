package tables

import "time"

type Message struct {
	Id             int64     `json:"id"`
	ConversationId int64     `json:"conversation_id"`
	SenderId       int64     `json:"sender_id"`
	MessageType    string    `json:"message_type"`
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
