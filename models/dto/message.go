package dto

import "time"

type LastMessageDto struct {
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type MessageDto struct {
	Id              int64     `json:"id"`
	Message         string    `json:"message"`
	CreatedAt       time.Time `json:"created_at"`
	SenderId        int64     `json:"sender_id"`
	SenderName      string    `json:"sender_name"`
	SenderUsername  string    `json:"sender_username"`
	SenderPesantren string    `json:"sender_pesantren"`
	LastDateData    string    `json:"last_data_date"`
}

type ResponseMessageDto struct {
	TotalMessage         int64       `json:"total_message"`
	CurrentTotalResponse int64       `json:"current_total_response"`
	LimitMessage         int         `json:"limit_message"`
	Page                 int         `json:"page"`
	TotalPage            int         `json:"total_page"`
	Data                 interface{} `json:"data"`
}
