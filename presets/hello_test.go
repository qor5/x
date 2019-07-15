package presets_test

import (
	"fmt"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/theplant/htmltestingutils"

	"github.com/sunfmin/bran/presets"

	"github.com/jinzhu/gorm"

	"github.com/sunfmin/bran/ui"
	"github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"

	. "github.com/sunfmin/bran/vuetify"
)

type User struct {
	ID        int
	Name      string
	Bool1     bool
	Date1     *time.Time
	Int1      int
	Float1    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TestHello(t *testing.T) {
	p := presets.New().URIPrefix("/admin")

	var db *gorm.DB

	m := p.Model(&User{}).URIName("user")
	m.Labels(
		"Name", "名字",
		"Bool1", "性别",
		"Float1", "体重",
	).Placeholders(
		"Name", "请输入你的名字",
	)

	m.SearchColumns("name", "bool1")

	l := m.Listing("Name", "Bool1", "Float1", "Int1")
	l.Field("Name").Label("列表的名字").ComponentFunc(func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		return h.A().Href(fmt.Sprintf("/users/%d", u.ID)).Text(u.Name)
	})

	l.BulkAction("ApproveAll").UpdateFunc(func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error) {
		err = db.Model(&User{}).Where("id IN (?)", selectedIds).UpdateColumn("approved_at = ?", time.Now()).Error
		return
	}).ComponentFunc(func(ctx *ui.EventContext) h.HTMLComponent {
		return VBtn("ApproveAll")
	})

	fl := l.Filtering("Name", "Int1", "Date1")
	fl.Filter("Name").ComponentFunc(func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		return vuetify.VAutocomplete().FieldName("Name").Value(u.Name)
	})

	ef := m.Editing("Name", "Bool1")
	ef.Field("Name").Label("名字").ComponentFunc(func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent {
		u := obj.(*User)
		return vuetify.VAutocomplete().FieldName("Name").Value(u.Name)
	}).SetterFunc(func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) {
		u := obj.(*User)
		ns := form.Value["Name"]
		if len(ns) > 0 {
			u.Name = ns[0]
		}
	})

	dp := m.Detailing("Name", "Bool1", "Float1", "Int1", "Date1", "CreatedAt", "UpdatedAt")
	ie := dp.Field("Bool1").InplaceEdit()
	ie.ComponentFunc(func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent {
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
	}).ComponentFunc(func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent {
		return VBtn("Approve")
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/admin/users", nil)
	p.ServeHTTP(w, r)
	diff := htmltestingutils.PrettyHtmlDiff(r.Body, "*", "abc")
	if len(diff) > 0 {
		t.Error(diff)
	}
}
