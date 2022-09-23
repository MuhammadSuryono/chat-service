package tables

import (
	"chat-service/Exception"
	"chat-service/config/database"
)

func Migrate() {
	defer Exception.GetError()
	database.Connection.AutoMigrate(&Message{})
	database.Connection.AutoMigrate(&GroupChat{})
	database.Connection.AutoMigrate(&Participant{})
	database.Connection.AutoMigrate(&User{})

}
