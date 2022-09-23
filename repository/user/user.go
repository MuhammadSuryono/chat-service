package user

import (
	"chat-service/config/database"
	"chat-service/models/tables"
	"fmt"
)

type user struct {
}

var Repository user

func (receiver user) FindAll() (users []tables.User) {
	database.Connection.Find(&users)
	return
}

func (receiver user) Save(value tables.User) tables.User {
	if err := database.Connection.Create(&value).Error; err != nil {
		panic(fmt.Sprintf("Can't save new data user. With error: %v", err.Error()))
	}
	return value
}

func (receiver user) FindByUsername(username string) (u tables.User) {
	if err := database.Connection.Where("username = ?", username).First(&u).Error; err != nil {
		panic("Failed get data user with username: " + username)
	}
	return
}

func (receiver user) FindByUsernameIsExist(username string) bool {
	var u tables.User
	database.Connection.Where("username = ?", username).First(&u)

	return username == u.Username
}
