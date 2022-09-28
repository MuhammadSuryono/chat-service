package tables

import "time"

type User struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(32)"`
	Email       string    `json:"email" gorm:"type:varchar(256);UNIQUE"`
	Password    string    `json:"-" gorm:"type:varchar(256)"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(256)"`
	Pesantren   string    `json:"pesantren" gorm:"type:varchar(256)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
