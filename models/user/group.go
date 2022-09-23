package tables

import (
	"github.com/gofrs/uuid"
)

type GroupChat struct {
	Id        uuid.Gen `json:"id" gorm:"primaryKey"`
	GroupName string   `json:"group_name"`
	LimitUser int      `json:"limit_user"`
	ChannelID string   `json:"channel_id"`
}
