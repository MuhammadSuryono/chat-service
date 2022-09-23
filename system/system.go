package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitialDefaultValue(envKey interface{}, defaultValue interface{}) interface{} {
	if envKey == "" || envKey == nil {
		return defaultValue
	}
	return envKey
}

func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

var Context *gin.Context

func TimeClock(t time.Time) string {
	diff := time.Now().Sub(t)
	hours, minutes, _ := t.Clock()
	year, month, day := t.Date()
	if (diff.Hours() > 24 && diff.Minutes() > 0) && diff.Hours() <= 48 {
		return fmt.Sprintf("Yesterday %02d:%02d", hours, minutes)
	} else if diff.Hours() > 48 && diff.Minutes() > 0 {
		return fmt.Sprintf("%02d/%02d/%02d %02d:%02d", year, month, day, hours, minutes)
	}
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
