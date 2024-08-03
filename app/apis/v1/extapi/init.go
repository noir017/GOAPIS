package extapi

import (
	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/app/global"
)

func Init() {
	// 自动创建orm
	global.DB.AutoMigrate(&Hotword{})
	global.DB.AutoMigrate(&Config{})
}

func RegisterApis(r *gin.Engine) {
	Init()

	eapi := ExtApi{}
	apisRouter := r.Group("/keywords")
	// apisRouter.GET("fetch", eapi.GetNewsEveryDay)
	apisRouter.GET("random", eapi.GetRandomNews)
	apisRouter.GET("duolicate", eapi.DuolicateNews)
	apisRouter.GET("/", eapi.Hello)
}
