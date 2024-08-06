package app

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/noir017/goapis/app/global"
	"github.com/noir017/goapis/config"
	"github.com/noir017/goapis/pkg/dbo"
	"github.com/noir017/goapis/pkg/tools"
	"github.com/sirupsen/logrus"
)

func Init() {
	// 创建或打开日志文件
	// file, err := os.Create("logs/run.log")
	// if err != nil {
	// 	panic(err)
	// }
	gin.SetMode(gin.DebugMode)

	tools.Makedirs("logs")
	tools.Makedirs("config")
	tools.Makedirs("static")
	// 初始化全局变量
	global.DB = dbo.InitDB()
	global.Gin = gin.Default()
	global.Config = config.ReadYml()

	rotator, err := rotatelogs.New(
		"logs/run_%Y%W.log",
		rotatelogs.WithRotationTime(7*24*time.Hour),
		rotatelogs.WithMaxAge(6*7*24*time.Hour),
		rotatelogs.WithRotationCount(6),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	// 创建一个同时写入文件和标准输出的多写器
	multiWriter := io.MultiWriter(os.Stdout, rotator)

	// 配置 logrus
	logrus.SetOutput(multiWriter)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}
