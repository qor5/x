package gormx

import (
	"testing"

	"github.com/theplant/testenv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	env, err := testenv.New().DBEnable(true).SetUp()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := env.TearDown(); err != nil {
			panic(err)
		}
	}()

	db = env.DB
	db.Logger = db.Logger.LogMode(logger.Info)
	db.Config.DisableForeignKeyConstraintWhenMigrating = true
	db.Config.TranslateError = true

	m.Run()
}
