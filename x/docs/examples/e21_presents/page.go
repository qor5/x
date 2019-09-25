// @snippet_begin(PresetHelloWorldSample)
package e21_presents

import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gormop"
	"github.com/goplaid/x/vuetify"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/theplant/htmlgo"
)

type Customer struct {
	ID          int
	Name        string
	Email       string
	Description string
	CompanyID   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func setupDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/my.db")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	err = db.AutoMigrate(
		&Customer{},
		&Company{},
	).Error
	if err != nil {
		panic(err)
	}
	return
}

func PresetsHelloWorld() (r *presets.Builder) {
	db := setupDB()
	r = presets.New().
		URIPrefix(PresetsHelloWorldPath).
		DataOperator(gormop.DataOperator(db))
	r.Model(&Customer{})
	return
}

const PresetsHelloWorldPath = "/samples/presets-hello-world"

// @snippet_end

// @snippet_begin(PresetListingCustomization01Sample)

type Company struct {
	ID   int
	Name string
}

func PresetsListingCustomization01() (r *presets.Builder) {
	db := setupDB()
	r = presets.New().
		URIPrefix(PresetsListingCustomization01PATH).
		DataOperator(gormop.DataOperator(db))
	cust := r.Model(&Customer{})

	cl := cust.Listing("ID", "Name", "Company", "Email")
	cl.Field("Company").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		c := obj.(*Customer)
		var comp Company
		if c.CompanyID > 0 {
			db.First(&comp, "id = ?", c.CompanyID)
		}
		return Td(
			web.Bind(
				A().Text(comp.Name)).
				URL(PresetsListingCustomization01PATH+"/companies").
				EventFunc(presets.DrawerEdit, fmt.Sprint(comp.ID)),
		)
	})

	ce := cust.Editing("Name", "CompanyID")
	ce.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		c := obj.(*Customer)
		var comps []Company
		db.Find(&comps)
		return vuetify.VSelect().
			Label("Company").
			Items(comps).
			ItemText("Name").
			ItemValue("ID").
			Value(c.CompanyID).
			FieldName("CompanyID")
	})

	r.Model(&Company{})

	return
}

const PresetsListingCustomization01PATH = "/samples/presets-listing-customization01"

// @snippet_end
