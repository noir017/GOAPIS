package extapi

import (
	"time"

	"github.com/noir017/goapis/app/global"
	"github.com/noir017/goapis/pkg/tools"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Name    string
	RunTime time.Time
}

func CheckExecuted(funcName string) bool {
	// today := time.Now().Format("2006-01-02")
	today := tools.GetTimeToday()
	var record Config
	result := global.DB.Where("name =? AND DATE(run_time) =?", funcName, today).First(&record)

	return result.RowsAffected > 0
	// if result.RowsAffected > 0 {
	// 	// fmt.Println("Function1 已经在今天运行过，不再执行")
	// 	return true
	// }
	// return false
}

func RecordExecuted(funcName string) {
	newRecord := Config{
		Name:    funcName,
		RunTime: time.Now(),
	}
	global.DB.Create(&newRecord)
}
