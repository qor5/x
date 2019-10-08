// @snippet_begin(PresetHelloWorldSample)
package e21_presents

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/actions"
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
	TermAgreedAt    *time.Time
	ApprovalComment string
}

var DB *gorm.DB

func init() {
	DB = setupDB()
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

func PresetsHelloWorld(b *presets.Builder) (m *presets.ModelBuilder, db *gorm.DB) {
	db = DB
	b.URIPrefix(PresetsHelloWorldPath).
		DataOperator(gormop.DataOperator(db))
	m = b.Model(&Customer{})

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
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, db = PresetsHelloWorld(b)
	b.URIPrefix(PresetsListingCustomizationFieldsPath)

	cl = cust.Listing("ID", "Name", "Company", "Email").
		SearchColumns("name", "email")
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
				EventFunc(actions.DrawerEdit, fmt.Sprint(comp.ID)),
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
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationFields(b)
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

// @snippet_begin(PresetsListingCustomizationTabsSample)

func PresetsListingCustomizationTabs(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationFilters(b)
	b.URIPrefix(PresetsListingCustomizationTabsPath)

	cl.FilterTabsFunc(func(ctx *web.EventContext) []*presets.FilterTab {
		var c Company
		db.First(&c)
		return []*presets.FilterTab{
			{
				Label: "Felix",
				Query: url.Values{"name.ilike": []string{"felix"}},
			},
			{
				Label: "The Plant",
				Query: url.Values{"company": []string{fmt.Sprint(c.ID)}},
			},
			{
				Label: "Approved",
				Query: url.Values{"approved.gt": []string{fmt.Sprint(1)}},
			},
			{
				Label: "All",
				Query: url.Values{"all": []string{"1"}},
			},
		}
	})
	return
}

const PresetsListingCustomizationTabsPath = "/samples/presets-listing-customization-tabs"

// @snippet_end

// @snippet_begin(PresetsListingCustomizationBulkActionsSample)

func PresetsListingCustomizationBulkActions(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, db = PresetsListingCustomizationTabs(b)
	b.URIPrefix(PresetsListingCustomizationBulkActionsPath)

	cl.BulkAction("Approve").Label("Approve").
		UpdateFunc(func(selectedIds []string, ctx *web.EventContext) (err error) {
			comment := ctx.R.FormValue("ApprovalComment")
			if len(comment) < 10 {
				ctx.Flash = "comment should larger than 10"
				return
			}
			err = db.Model(&Customer{}).
				Where("id IN (?)", selectedIds).
				Updates(map[string]interface{}{"approved_at": time.Now(), "approval_comment": comment}).Error
			if err != nil {
				ctx.Flash = err.Error()
			}
			return
		}).
		ComponentFunc(func(selectedIds []string, ctx *web.EventContext) HTMLComponent {
			comment := ctx.R.FormValue("ApprovalComment")
			errorMessage := ""
			if ctx.Flash != nil {
				errorMessage = ctx.Flash.(string)
			}
			return v.VTextField().
				FieldName("ApprovalComment").
				Value(comment).
				Label("Comment").
				ErrorMessages(errorMessage)
		})

	cl.BulkAction("Delete").Label("Delete").
		UpdateFunc(func(selectedIds []string, ctx *web.EventContext) (err error) {
			err = db.Where("id IN (?)", selectedIds).Delete(&Customer{}).Error
			return
		}).
		ComponentFunc(func(selectedIds []string, ctx *web.EventContext) HTMLComponent {
			return Div().Text(fmt.Sprintf("Are you sure you want to delete %s ?", selectedIds)).Class("title deep-orange--text")
		})

	return
}

const PresetsListingCustomizationBulkActionsPath = "/samples/presets-listing-customization-bulk-actions"

// @snippet_end
