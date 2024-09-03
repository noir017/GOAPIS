package bing

import (
	"time"

	"github.com/noir017/goapis/app/global"
	"gorm.io/gorm"
)

// 储存格式化后的数据
type Rewards struct {
	gorm.Model
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"CreatedAt"`
}

//	func AddHotword(hw Hotword) {
//		global.DB.Create(hw)
//	}
func SaveRewards(rw Rewards) {
	global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&rw).Error
		if err != nil {
			return err
		}
		// 进行其他操作，例如查询刚刚创建的记录
		// createdRewards := Rewards{}
		// tx.First(&createdRewards, rw.ID)
		return nil
	})
}
