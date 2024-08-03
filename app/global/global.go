package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Gin *gin.Engine
	DB  *gorm.DB
)
