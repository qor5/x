package admin

import (
	"os"

	"github.com/goplaid/x/goplaid/admintemplate/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DB_PARAMS"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	err = db.AutoMigrate(
		&models.Post{},
	).Error
	if err != nil {
		panic(err)
	}
	return
}
