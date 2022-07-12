package e10_vuetify_autocomplete

// @snippet_begin(VuetifyAutoCompleteSample)

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	"github.com/goplaid/x/vuetifyx"
	h "github.com/theplant/htmlgo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type myFormValue struct {
	Values1 []string
	Values2 []string
	Value3  string
}

type User struct {
	Login string
	Name  string
}

var selectedItems1 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "charles", Name: "Charles"},
}

var options1 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

var selectedItems2 = []*User{
	{Login: "charles", Name: "Charles"},
}

var selectedItems3 = []*User{
	{Login: "charles", Name: "Charles"},
}

var options2 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

var globalState = &myFormValue{
	Values1: []string{
		"sam",
		"charles",
	},
	Values2: []string{
		"charles",
	},
	Value3: "charles",
}

type Product struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

var remoteRes *vuetifyx.RemoteResource

func init() {
	db, err := gorm.Open(sqlite.Open("/tmp/my.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	db.Where("1=1").Delete(&Product{})

	for i := 1; i < 300; i++ {
		db.Create(&Product{Name: fmt.Sprintf("Product %d", i)})
	}

	remoteRes = vuetifyx.RegisterRemoteResource(
		&vuetifyx.RemoteResource{
			Model:        &Product{},
			DB:           db,
			SearchFields: []string{"Name"},
			RemoteURL:    "/samples/vuetify-auto-complete",
		},
	)
}

func VuetifyAutocomplete(ctx *web.EventContext) (pr web.PageResponse, err error) {

	result := h.Ul()
	for _, v := range globalState.Values1 {
		result.AppendChildren(h.Li().Text(v))
	}
	pr.Body = VContainer(
		h.H1("An auto complete that you can select multiple with static data"),
		VAutocomplete().
			Items(options1).
			FieldName("Values1").
			ItemText("Name").
			ItemValue("Login").
			Label("Static Options").
			Value(globalState.Values1),
		result,

		h.H1("An auto complete that you can select multiple with remote resource"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load Options from Remote").
			SetRemoteResource(remoteRes),

		h.H1("VSelect"),
		VSelect().
			Items(options1).    // Items is the data source
			ItemText("Name").   // ItemText is the value that would be displayed to user. the argument is the corresponding field name in the Items. here is user.Name
			ItemValue("Login"). // ItemValue is the value that will be passed with the form. same with ItemText, here is user.Login
			FieldName("Value3").
			Solo(true).
			Value(globalState.Value3),
		h.Pre(globalState.Value3),
		VBtn("Update").
			Color("success").
			OnClick("update"),
	)
	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	globalState = &myFormValue{}
	ctx.MustUnmarshalForm(globalState)

	selectedItems1 = []*User{}
	for _, login := range globalState.Values1 {
		for _, u := range options1 {
			if u.Login == login {
				selectedItems1 = append(selectedItems1, u)
			}
		}
	}

	selectedItems2 = []*User{}
	for _, login := range globalState.Values2 {
		for _, u := range options2 {
			if u.Login == login {
				selectedItems2 = append(selectedItems2, u)
			}
		}
	}

	selectedItems3 = []*User{}
	for _, u := range options1 {
		if u.Login == globalState.Value3 {
			selectedItems3 = append(selectedItems3, u)
		}
	}
	r.Reload = true

	return
}

var VuetifyAutocompletePB = web.Page(VuetifyAutocomplete).
	EventFunc("update", update).
	EventFunc(vuetifyx.AutocompleteRemoteResEvent, FetchItemsFromRemoteResource)
	// EventFunc(vuetifyx.AutocompleteRemoteResEvent, vuetifyx.FetchItemsFromRemoteResource)

const VuetifyAutoCompletePath = "/samples/vuetify-auto-complete"

// @snippet_end

// rewrite this event handle to fix "sql: Scan called without calling Next" on sqlite3 by using the model Product directly on db query.
func FetchItemsFromRemoteResource(ctx *web.EventContext) (r web.EventResponse, err error) {
	var res = remoteRes
	var db = res.DB
	if res.NewDB != nil {
		db = res.NewDB(ctx)
	}

	db = db.Session(&gorm.Session{})
	if searchKey := ctx.R.FormValue("keyword"); searchKey != "" {
		var sqlKeys []string
		var values []interface{}
		for _, key := range res.SearchFields {
			key = strings.ToLower(key)
			if key == "id" {
				if v, err := strconv.Atoi(searchKey); err == nil {
					sqlKeys = append(sqlKeys, fmt.Sprintf("%s = ?", key))
					values = append(values, v)
				}
				continue
			}

			sqlKeys = append(sqlKeys, fmt.Sprintf("%s ILIKE ?", key))
			values = append(values, fmt.Sprintf("%%%s%%", searchKey))

		}
		db = db.Where(strings.Join(sqlKeys, " or "), values...)
	}

	var (
		count   int64
		current int
		total   int
		offset  int
	)

	db.Model(res.Model).Count(&count)
	if v, err := strconv.Atoi(ctx.R.FormValue("current")); err == nil {
		current = v
	}

	if res.IsPaging {
		offset = (current - 1) * res.SizePerPage
	} else {
		offset = current
	}

	var datas = []Product{}
	// var datas = reflect.MakeSlice(reflect.SliceOf(reflect.Indirect(reflect.ValueOf(res.Model)).Type()), 0, 0).Interface()
	db.Model(res.Model).Offset(offset).Limit(res.SizePerPage).Find(&datas)

	reflectValue := reflect.Indirect(reflect.ValueOf(datas))
	var items []vuetifyx.OptionItem
	for i := 0; i < reflectValue.Len(); i++ {
		var value = fmt.Sprintf("%v", reflectValue.Index(i).FieldByName(res.OptionValue).Interface())

		var text string
		if res.OptionText != nil {
			text = res.OptionText(reflectValue.Index(i).Interface())
		} else {
			text = fmt.Sprintf("%v", reflectValue.Index(i).FieldByName(res.SearchFields[0]).Interface())
		}

		var icon string
		if res.OptionIcon != nil {
			icon = res.OptionIcon(reflectValue.Index(i).Interface())
		}

		items = append(items, vuetifyx.OptionItem{
			Text:  text,
			Value: value,
			Icon:  icon,
		})
	}

	if res.IsPaging {
		total = int(count) / res.SizePerPage
		if int(count)%res.SizePerPage > 0 {
			total++
		}
	} else {
		total = int(count)
		current += len(items)
	}

	r.Data = vuetifyx.RemoteResourceResult{
		Items:   items,
		Total:   total,
		Current: current,
	}
	return
}
