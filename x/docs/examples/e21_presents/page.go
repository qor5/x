// @snippet_begin(PresetHelloWorldSample)
package e21_presents

import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gormop"
	v "github.com/goplaid/x/vuetify"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/theplant/htmlgo"
)

type Customer struct {
	ID              int
	Name            string
	Email           string
	Description     string
	CompanyID       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ApprovedAt      *time.Time
	ApprovalComment string
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

// @snippet_begin(PresetsListingCustomizationFieldsSample)

type Company struct {
	ID   int
	Name string
}

func PresetsListingCustomizationFields(b *presets.Builder) (
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	db = setupDB()

	b.URIPrefix(PresetsListingCustomizationFieldsPath).
		DataOperator(gormop.DataOperator(db))

	cust := b.Model(&Customer{})

	cl = cust.Listing("ID", "Name", "Company", "Email")
	cl.Field("Company").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		c := obj.(*Customer)
		var comp Company
		if c.CompanyID > 0 {
			db.First(&comp, "id = ?", c.CompanyID)
		}
		return Td(
			web.Bind(
				A().Text(comp.Name)).
				URL(PresetsListingCustomizationFieldsPath+"/companies").
				EventFunc(presets.DrawerEdit, fmt.Sprint(comp.ID)),
		)
	})

	ce = cust.Editing("Name", "CompanyID")
	ce.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) HTMLComponent {
		c := obj.(*Customer)
		var comps []Company
		db.Find(&comps)
		return v.VSelect().
			Label("Company").
			Items(comps).
			ItemText("Name").
			ItemValue("ID").
			Value(c.CompanyID).
			FieldName("CompanyID")
	})

	b.Model(&Company{})

	return
}

const PresetsListingCustomizationFieldsPath = "/samples/presets-listing-customization-fields"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationFiltersSample)

func PresetsListingCustomizationFilters(b *presets.Builder) (
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cl, ce, db = PresetsListingCustomizationFields(b)
	b.URIPrefix(PresetsListingCustomizationFiltersPath)

	cl.FilterDataFunc(func(ctx *web.EventContext) v.FilterData {
		var companyOptions []*v.SelectItem
		err := db.Model(&Company{}).Select("name as text, id as value").Scan(&companyOptions).Error
		if err != nil {
			panic(err)
		}

		return []*v.FilterItem{
			{
				Key:          "created",
				Label:        "Created",
				ItemType:     v.ItemTypeDate,
				SQLCondition: `cast(strftime('%%s', created_at) as INTEGER) %s ?`,
			},
			{
				Key:          "approved",
				Label:        "Approved",
				ItemType:     v.ItemTypeDate,
				SQLCondition: `cast(strftime('%%s', approved_at) as INTEGER) %s ?`,
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     v.ItemTypeString,
				SQLCondition: `name %s ?`,
			},
			{
				Key:          "company",
				Label:        "Company",
				ItemType:     v.ItemTypeSelect,
				SQLCondition: `company_id %s ?`,
				Options:      companyOptions,
			},
		}
	})
	return
}

const PresetsListingCustomizationFiltersPath = "/samples/presets-listing-customization-filters"

// @snippet_end
