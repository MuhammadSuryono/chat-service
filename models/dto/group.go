package dto

type GroupChatDto struct {
	Id          int64          `json:"id"`
	GroupName   string         `json:"group_name"`
	ChannelID   string         `json:"channel_id"`
	LastMessage LastMessageDto `json:"last_message"`
}
