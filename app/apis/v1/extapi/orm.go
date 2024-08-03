package extapi

import (
	"fmt"

	"github.com/noir017/goapis/app/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 储存格式化后的数据
type Hotword struct {
	gorm.Model
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

//	func AddHotword(hw Hotword) {
//		global.DB.Create(hw)
//	}
func SaveHotwords(hws []Hotword) {
	global.DB.Transaction(func(tx *gorm.DB) error {
		for _, hw := range hws {
			if err := tx.Create(&hw).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
func GetRandomHotwords() []Hotword {
	var hws []Hotword
	// 使用 SQL 的 RANDOM 函数和 LIMIT 子句获取随机的 50 条记录
	result := global.DB.Order("RANDOM()").Limit(50).Find(&hws)
	if result.Error != nil {
		panic("GetRandomHotwords：获取数据错误")
	}
	return hws
}

func Duolicate() {

	var hotwords []Hotword
	global.DB.Find(&hotwords)

	// 使用 map 来记录已经存在的 title
	titleMap := make(map[string]uint)

	for _, hotword := range hotwords {
		if _, exists := titleMap[hotword.Title]; exists {
			// 如果 title 已经存在，则删除当前记录
			global.DB.Delete(&Hotword{}, hotword.ID)
			logrus.Info(fmt.Sprintf("删除重复Hotword， ID %d 标题 %s\n", hotword.ID, hotword.Title))
		} else {
			// 如果 title 不存在，则添加到 map 中
			titleMap[hotword.Title] = hotword.ID
		}
	}
}
