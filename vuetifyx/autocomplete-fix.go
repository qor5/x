package vuetifyx

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/goplaid/web"
	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type VXAutocompleteBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VXAutocomplete(children ...h.HTMLComponent) (r *VXAutocompleteBuilder) {
	r = &VXAutocompleteBuilder{
		tag: h.Tag("vx-autocomplete").Children(children...),
	}
	r.Multiple(true)

	return
}

func (b *VXAutocompleteBuilder) ErrorMessages(v ...string) (r *VXAutocompleteBuilder) {
	vuetify.SetErrorMessages(b.tag, v)
	return b
}

func (b *VXAutocompleteBuilder) Items(v interface{}) (r *VXAutocompleteBuilder) {
	b.items = v
	return b
}

func (b *VXAutocompleteBuilder) FieldName(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VXAutocompleteBuilder) SelectedItems(v interface{}) (r *VXAutocompleteBuilder) {
	b.selectedItems = v
	return b
}

func (b *VXAutocompleteBuilder) ItemsEventFunc(eventFuncId string) (r *VXAutocompleteBuilder) {

	b.tag.Attr(":items-event-func-id", &web.EventFuncID{
		ID: eventFuncId,
	})
	return b
}

func (b *VXAutocompleteBuilder) SetRemoteResource(res *RemoteResource) (r *VXAutocompleteBuilder) {
	if res.Model == nil {
		panic("model is required")
	}

	if res.DB == nil && res.NewDB == nil {
		panic("db is required")
	}

	if len(res.SearchKeys) == 0 {
		panic("search keys are required")
	}

	if res.SetOptionValue == nil {
		panic("set option value is required")
	}

	if res.SizePerPage == 0 {
		res.SizePerPage = 20
	}

	var resName = res.Name
	if resName == "" {
		resName = reflect.ValueOf(res.Model).Elem().Type().Name()
	}

	if remoteResources[resName] != nil {
		panic(fmt.Sprintf("duplicate remote resource name: %s", resName))
	}

	remoteResources[resName] = res
	b.tag.Attr(":remote-res", resName)
	return b
}

func (b *VXAutocompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)
	return b.tag.MarshalHTML(ctx)
}

var remoteResources = map[string]*RemoteResource{}

const AutocompleteRemoteResEvent = "autocomplete-remote-res-event"

type RemoteResource struct {
	Name           string
	Model          interface{}
	DB             *gorm.DB
	NewDB          func(ctx *web.EventContext) *gorm.DB
	SizePerPage    int
	SearchKeys     []string
	SetOptionText  func(interface{}) string
	SetOptionValue func(interface{}) string
}

func FetchItemsFromRemoteResource(ctx *web.EventContext) (r web.EventResponse, err error) {
	var name = ctx.R.FormValue("name")

	if remoteResources[name] == nil {
		return r, fmt.Errorf("remote resource not found: %s", name)
	}

	var res = remoteResources[name]
	var db = res.DB
	if res.NewDB != nil {
		db = res.NewDB(ctx)
	}

	if searchKey := ctx.R.FormValue("keyword"); searchKey != "" {
		for _, key := range res.SearchKeys {
			db = db.Or(key+" ILIKE ?", fmt.Sprintf("%%%s%%", searchKey))
		}
	}

	var datas = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(res.Model)), 0, 0)

	var offset int
	offset, _ = strconv.Atoi(ctx.R.FormValue("offset"))
	db.Offset(offset).Limit(res.SizePerPage).Find(&datas)

	reflectValue := reflect.Indirect(reflect.ValueOf(datas))
	var items []SelectItem
	for i := 0; i < reflectValue.Len(); i++ {
		var value = res.SetOptionValue(reflectValue.Index(i).Interface())
		var text string

		if res.SetOptionText != nil {
			text = res.SetOptionText(reflectValue.Index(i).Interface())
		} else {
			text = value
		}

		items = append(items, SelectItem{
			Text:  text,
			Value: value,
		})
	}

	r.Data = items
	return
}
