package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/app/apis/v1/extapi"
	"github.com/noir017/goapis/app/global"
	"github.com/noir017/goapis/pkg/dbo"
)

func StartService() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(gin.Logger())

	// 初始化全局变量
	global.DB = dbo.InitDB()
	global.Gin = r

	// 注册api路由
	extapi.RegisterApis(r)

	// 定义一个默认的路由处理函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello,gin",
		})
	})

	// 启动服务器，监听 8080 端口
	// r.Run(":8080")
	r.RunTLS(":8080", "/config/.config/code-server/code-serve.pem", "/config/.config/code-server/code-serve-key.pem")
}
