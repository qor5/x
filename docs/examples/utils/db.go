package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(sqlite.Open("/tmp/my.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)

	return
}

type Page struct {
	ID   int
	Name string
}
