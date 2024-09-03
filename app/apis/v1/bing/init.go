package bing

import (
	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/app/global"
)

func Init() {
	// 自动创建orm
	global.DB.AutoMigrate(&Rewards{})
}

func RegisterApis(r *gin.Engine) {
	Init()

	bapi := BingApi{}
	apisRouter := r.Group("/bing")
	// apisRouter.GET("fetch", eapi.GetNewsEveryDay)
	apisRouter.GET("upload", bapi.RecordReward)
	apisRouter.GET("/", bapi.Hello)
}
