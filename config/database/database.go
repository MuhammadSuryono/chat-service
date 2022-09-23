package database

import (
	"chat-service/system"
	"gorm.io/gorm"
)

var Connection *gorm.DB

const (
	MysqlConnectionPetra = "mysql_connection_petra"
)

type configOtherDatabase struct {
	config map[string]map[string]interface{}
}

var configMapDatabase = map[string]map[string]interface{}{
	MysqlConnectionPetra: {
		"driver":        system.InitialDefaultValue("DB_DRIVER_MYSQL", ""),
		"host":          system.InitialDefaultValue("DB_HOST_MYSQL", ""),
		"password":      system.InitialDefaultValue("DB_PASS_MYSQL", ""),
		"user":          system.InitialDefaultValue("DB_USER_MYSQL", ""),
		"database_name": system.InitialDefaultValue("DB_NAME_MYSQL", ""),
		"port":          system.InitialDefaultValue("DB_PORT_MYSQL", ""),
		"tables":        MysqlConnectionPetra,
	},
}

type ConnectionHandler struct {
	DB_HOST   string
	DB_PORT   string
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	TIMEZONE  string
	DB_DRIVER string
}
