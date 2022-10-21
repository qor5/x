package admin

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"os"

	"github.com/goplaid/x/goplaid/admintemplate/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (db *gorm.DB) {
	var err error
	// Create database connection
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")))
	if err != nil {
		panic(err)
	}

	// Set db log level
	db.Logger = db.Logger.LogMode(logger.Info)

	// Create data table in the database
	err = db.AutoMigrate(models.Post{})
	if err != nil {
		panic(err)
	}

	return
}
