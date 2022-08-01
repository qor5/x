package e10_vuetify_autocomplete

// @snippet_begin(VuetifyAutoCompleteSample)

import (
	"fmt"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gorm2op"
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

var loadMoreRes *vuetifyx.AutocompleteDataSource
var pagingRes *vuetifyx.AutocompleteDataSource
var ExamplePreset *presets.Builder

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

	ExamplePreset = presets.New()
	ExamplePreset.URIPrefix(VuetifyAutoCompletePresetPath).DataOperator(gorm2op.DataOperator(db))
	listing := ExamplePreset.Model(&Product{}).Listing()
	loadMoreRes = listing.ConfigureAutocompleteDataSource(
		&presets.AutocompleteDataSourceConfig{
			OptionValue: "ID",
			OptionText:  "Name",
			KeywordColumns: []string{
				"Name",
			},
			PerPage: 50,
		},
		"loadMore",
	)

	pagingRes = listing.ConfigureAutocompleteDataSource(
		&presets.AutocompleteDataSourceConfig{
			OptionValue: "ID",
			OptionText:  "Name",
			KeywordColumns: []string{
				"Name",
			},
			PerPage:  20,
			IsPaging: true,
			OrderBy:  "Name",
		},
		"paging",
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

		h.H1("An auto complete that you can select multiple with remote resource by load more"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load Options from Remote").
			SetRemoteResource(loadMoreRes),

		h.H1("An auto complete that you can select multiple with remote resource by paging"),
		vuetifyx.VXAutocomplete().
			FieldName("Values2").
			Label("Load Options from Remote").
			SetRemoteResource(pagingRes),

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
	EventFunc("update", update)

const VuetifyAutoCompletePath = "/samples/vuetify-auto-complete"
const VuetifyAutoCompletePresetPath = "/samples/vuetify-auto-complete-preset"

// @snippet_end
