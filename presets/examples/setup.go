package examples

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sunfmin/bran/presets"
	"github.com/sunfmin/bran/presets/gormop"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type Thumb struct {
	Name string
}

type User struct {
	ID              int
	Name            string
	JobTitle        string
	Bool1           bool
	Date1           *time.Time
	Int1            int
	Float1          float64
	Thumb1          *Thumb
	CompanyID       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ApprovalComment string
}

type Company struct {
	ID   int
	Name string
}

type Product struct {
	ID        int
	Name      string
	OwnerName string
}

func Preset1(db *gorm.DB) (r *presets.Builder) {

	err := db.AutoMigrate(&User{}, &Product{}, &Company{}).Error
	if err != nil {
		panic(err)
	}

	p := presets.New().URIPrefix("/admin").PrimaryColor("cyan darken-3")

	p.BrandFunc(func(ctx *ui.EventContext) h.HTMLComponent {
		return h.Components(
			//h.Img("https://material.io/tools/icons/static/ic_material_192px_light.svg").Style("height: 32px"),
			VIcon("directions_boat").Class("pr-2"),
			VToolbarTitle("My Admin"),
		)
	}).BrandTitle("The Plant")

	p.FieldType(&Thumb{}).
		ComponentFunc(presets.LISTING, func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		}).
		ComponentFunc(presets.DETAILING, func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		}).
		ComponentFunc(presets.EDITING, func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		})

	p.DataOperator(gormop.DataOperator(db))

	p.MenuGroup("User Management").Icon("group")
	mp := p.Model(&Product{}).MenuIcon("laptop")
	mp.Listing().PerPage(3)

	m := p.Model(&User{}).URIName("user").MenuGroup("User Management")
	p.Model(&Company{}).MenuGroup("User Management")
	m.Labels(
		"Name", "名字",
		"Bool1", "性别",
		"Float1", "体重",
		"CompanyID", "公司",
	).Placeholders(
		"Name", "请输入你的名字",
	)

	l := m.Listing("Name", "CompanyID", "Bool1", "Float1", "Int1", "ApprovalComment").SearchColumns("name", "job_title").PerPage(5)
	l.Field("Name").Label("列表的名字").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		return h.Td(ui.Bind(h.A().Text(u.Name)).PushStateLink(fmt.Sprintf("/admin/users/%d/edit", u.ID)))
	})

	l.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		var comp Company
		err := db.Find(&comp, u.CompanyID).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return h.Td(ui.Bind(h.A().Text(comp.Name)).
			PageURL("/admin/companies").EventFunc("formDrawerEdit", fmt.Sprint(comp.ID)))
	})

	l.BulkAction("Approve").Label("Approve").UpdateFunc(func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error) {
		comment := ctx.R.FormValue("ApprovalComment")
		if len(comment) < 10 {
			ctx.Flash = "comment should larger than 10"
			return
		}
		err = db.Model(&User{}).
			Where("id IN (?)", selectedIds).
			Updates(map[string]interface{}{"approved_at": time.Now(), "approval_comment": comment}).Error
		if err != nil {
			ctx.Flash = err.Error()
		}
		return
	}).ComponentFunc(func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent {
		comment := ctx.R.FormValue("ApprovalComment")
		errorMessage := ""
		if ctx.Flash != nil {
			errorMessage = ctx.Flash.(string)
		}
		return VTextField().
			FieldName("ApprovalComment").
			Value(comment).
			Label("Comment").
			Error(len(errorMessage) > 0).
			ErrorMessages(errorMessage)
	})

	l.BulkAction("Delete").Label("Delete").UpdateFunc(func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error) {
		err = db.Where("id IN (?)", selectedIds).Delete(&User{}).Error
		return
	}).ComponentFunc(func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent {
		return h.Div().Text(fmt.Sprintf("Are you sure you want to delete %s ?", selectedIds)).Class("title deep-orange--text")
	})

	l.Filter([]*FilterItem{
		{
			Key:          "created",
			Label:        "Created",
			ItemType:     ItemTypeDate,
			SQLCondition: `extract(epoch from created_at) %s ?`,
		},
		{
			Key:          "name",
			Label:        "Name",
			ItemType:     ItemTypeString,
			SQLCondition: `name %s ?`,
		},
	})

	ef := m.Editing("Name", "CompanyID", "Bool1", "Int1")
	ef.Field("Name").Label("名字").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		//u := obj.(*User)
		return VAutocomplete().
			FieldName("Name").
			Label(field.Label).
			Items([]string{"Felix", "Hello"}).
			Multiple(false).
			Value(reflectutils.MustGet(obj, field.Name))
	}).SetterFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) {
		u := obj.(*User)
		ns := form.Value["Name"]
		if len(ns) > 0 {
			u.Name = ns[0]
		}
	})

	ef.Field("CompanyID").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		var companies []*Company
		err := db.Find(&companies).Error
		if err != nil {
			panic(err)
		}
		return VSelect().
			FieldName("CompanyID").
			Label(field.Label).
			Items(companies).
			ItemText("Name").
			ItemValue("ID").
			Multiple(false).
			Value(u.CompanyID)
	})

	dp := m.Detailing("Name", "Bool1", "Float1", "Int1", "Date1", "CreatedAt", "UpdatedAt")
	ie := dp.Field("Bool1").InplaceEdit()
	ie.ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		//u := obj.(*User)
		return VCheckbox().FieldName("Bool1")
	}).UpdateFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) (err error) {
		u := obj.(*User)
		err = db.Model(&User{}).UpdateColumn("Name", u.Name).Error
		return
	})

	dp.Action("Approve").UpdateFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) (err error) {
		u := obj.(*User)
		err = db.Model(&User{}).Where("id = ?", u.ID).UpdateColumn("approved_at = ?", time.Now()).Error
		return
	}).ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *ui.EventContext) h.HTMLComponent {
		return VBtn("Approve")
	})
	return p
}
