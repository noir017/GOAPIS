package dbo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("config/data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
