package tables

type GroupChat struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	GroupName string `json:"group_name"`
	LimitUser int    `json:"limit_user"`
	ChannelID string `json:"channel_id"`
	IsActive  bool   `json:"is_active" gorm:"default:1"`
}

type GroupReadMessage struct {
}
