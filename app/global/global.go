package global

import (
	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/config"
	"gorm.io/gorm"
)

var (
	Gin    *gin.Engine
	DB     *gorm.DB
	Config config.Config
)
