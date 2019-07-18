package examples

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/sunfmin/bran/presets/gormop"

	"github.com/sunfmin/reflectutils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sunfmin/bran/presets"
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type Thumb struct {
	Name string
}

type User struct {
	ID        int
	Name      string
	JobTitle  string
	Bool1     bool
	Date1     *time.Time
	Int1      int
	Float1    float64
	Thumb1    *Thumb
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Preset1() (r *presets.Builder) {
	p := presets.New().URIPrefix("/admin")

	p.FieldType(&Thumb{}).
		ListingComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		}).
		DetailingComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		}).
		EditingComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
			i, err := reflectutils.Get(obj, field.Name)
			if err != nil {
				panic(err)
			}
			return h.Text(i.(*Thumb).Name)
		})

	db, err := gorm.Open("postgres", os.Getenv("TEST_DB"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	err = db.AutoMigrate(&User{}).Error
	if err != nil {
		panic(err)
	}
	p.DataOperator(gormop.DataOperator(db))

	m := p.Model(&User{}).URIName("user")
	m.Labels(
		"Name", "名字",
		"Bool1", "性别",
		"Float1", "体重",
	).Placeholders(
		"Name", "请输入你的名字",
	)

	l := m.Listing("Name", "Bool1", "Float1", "Int1").SearchColumns("name", "job_title")
	l.Field("Name").Label("列表的名字").ComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		return h.Td(h.A().Href(fmt.Sprintf("/admin/users/%d/edit", u.ID)).Text(u.Name))
	})

	l.BulkAction("ApproveAll").UpdateFunc(func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error) {
		err = db.Model(&User{}).Where("id IN (?)", selectedIds).UpdateColumn("approved_at = ?", time.Now()).Error
		return
	}).ComponentFunc(func(ctx *ui.EventContext) h.HTMLComponent {
		return VBtn("ApproveAll")
	})

	fl := l.Filtering("Name", "Int1", "Date1")
	fl.Filter("Name").ComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		return VAutocomplete().FieldName("Name").Value(u.Name).Label(field.Label).Items([]string{"1111", "2222"})
	})

	ef := m.Editing("Name", "Bool1", "Int1")
	ef.Field("Name").Label("名字").ComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
		//u := obj.(*User)
		return VAutocomplete().FieldName("Name").Label(field.Label).Items([]string{"Felix", "Hello"})
	}).SetterFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) {
		u := obj.(*User)
		ns := form.Value["Name"]
		if len(ns) > 0 {
			u.Name = ns[0]
		}
	})

	dp := m.Detailing("Name", "Bool1", "Float1", "Int1", "Date1", "CreatedAt", "UpdatedAt")
	ie := dp.Field("Bool1").InplaceEdit()
	ie.ComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
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
	}).ComponentFunc(func(obj interface{}, field *presets.Field, ctx *ui.EventContext) h.HTMLComponent {
		return VBtn("Approve")
	})
	return p
}
