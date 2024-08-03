package app

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/noir017/goapis/pkg/tools"
	"github.com/sirupsen/logrus"
)

func Init() {
	// 创建或打开日志文件
	// file, err := os.Create("logs/run.log")
	// if err != nil {
	// 	panic(err)
	// }
	tools.Makedirs("logs")
	tools.Makedirs("config")

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
