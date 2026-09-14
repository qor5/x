package exchange_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/qor5/x/v3/gormx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var tables = []interface{}{
	&TestExchangeModel{},
	&TestExchangeCompositePrimaryKeyModel{},
	&Phone{},
	&ShoppingSite{},
	&Camera{},
	&Intro{},
	&ExtraIntro{},
}

var dropTableStructs = []interface{}{
	&TestExchangeModel{},
	&TestExchangeCompositePrimaryKeyModel{},
	&Intro{},
	&ExtraIntro{},
	&ShoppingSite{},
	&Camera{},
	&Phone{},
}

type TestExchangeModel struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Age      *int
	Birth    *time.Time
	Appender string
}

type TestExchangeCompositePrimaryKeyModel struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"primarykey"`
	Age      *int
	Appender string
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	testSuite := gormx.MustStartTestSuite(ctx)
	defer func() {
		if err := testSuite.Stop(context.Background()); err != nil {
			fmt.Printf("Error during teardown: %v\n", err)
		}
	}()

	// Open a plain connection without OmitAssociationsPlugin, which MustStartTestSuite installs
	// via SetupDatabase. Exchange's importer relies on GORM creating nested associations for new
	// records during CreateInBatches, so the plugin must not be active on this connection.
	var err error
	db, err = gorm.Open(postgres.Open(testSuite.DSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger = db.Logger.LogMode(logger.Info)

	migrateTables()

	os.Exit(m.Run())
}

func migrateTables() {
	if err := db.AutoMigrate(tables...); err != nil {
		panic(err)
	}
}

func dropTables() {
	var err error
	err = db.Exec("drop table phone_selling_shopping_site").Error
	if err != nil {
		panic(err)
	}

	for _, m := range dropTableStructs {
		stmt := &gorm.Statement{DB: db}
		stmt.Parse(m)
		err = db.Exec(fmt.Sprintf("drop table %s", stmt.Schema.Table)).Error
		if err != nil {
			panic(err)
		}
	}
}

func initTables() {
	dropTables()
	migrateTables()
}

func ptrInt(v int) *int {
	return &v
}

func ptrTime(v time.Time) *time.Time {
	return &v
}
