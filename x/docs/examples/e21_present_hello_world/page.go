// @snippet_begin(PresetHelloWorldSample)
package e21_present_hello_world

import (
	"time"

	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gormop"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Customer struct {
	ID          int
	Name        string
	Email       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func PresetHelloWorld01() (r *presets.Builder) {
	db, err := gorm.Open("sqlite3", "/tmp/my.db")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	err = db.AutoMigrate(&Customer{}).Error
	if err != nil {
		panic(err)
	}

	r = presets.New().
		URIPrefix(PresetHelloWorldPath).
		DataOperator(gormop.DataOperator(db))
	r.Model(&Customer{})
	return
}

const PresetHelloWorldPath = "/samples/preset-hello-world"

// @snippet_end
