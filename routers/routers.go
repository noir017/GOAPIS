package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/app/apis/v1/extapi"
	"github.com/noir017/goapis/app/global"
)

func Cors() gin.HandlerFunc {
	// 跨域处理
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token,AccessToken,Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()

	}
}

func StartService() {
	global.Gin.Use(gin.Logger())
	global.Gin.Use(Cors())
	//设置静态文件目录
	global.Gin.Static("/static", "./static")

	// 注册api路由
	extapi.RegisterApis(global.Gin)

	// 定义一个默认的路由处理函数
	global.Gin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello,gin",
		})
	})

	// 启动服务器，监听 8080 端口
	// r.Run(":8080")
	if global.Config.TLS.Enable {
		global.Gin.RunTLS(":8080", global.Config.TLS.CA, global.Config.TLS.Key)
	} else {
		global.Gin.Run(":8080")
	}

}
