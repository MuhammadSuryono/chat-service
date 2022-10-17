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
	fmt.Println("Time now without location", time.Now(), time.Now().Location())
	loc, err := time.LoadLocation("Asia/Jakarta")
	fmt.Println("Timezone start", loc)
	if err != nil {
		_ = fmt.Sprintf("Error timezone %v", err)
	}

	fmt.Println("Time now With location", time.Now().In(loc), time.Now().Location())
	diff := time.Now().In(loc).Sub(t)
	hours, minutes, _ := t.Clock()
	year, month, day := t.Date()
	if (diff.Hours() > 24 && diff.Minutes() > 0) && diff.Hours() <= 48 {
		return fmt.Sprintf("Yesterday %02d:%02d", hours, minutes)
	} else if diff.Hours() > 48 && diff.Minutes() > 0 {
		return fmt.Sprintf("%02d/%02d/%02d %02d:%02d", year, month, day, hours, minutes)
	}
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func TimeNowString() string {
	t := time.Now()
	return fmt.Sprintf("%d-%2d-%2d", t.Year(), t.Month(), t.Day())
}
