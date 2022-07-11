package vuetifyx

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

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

func (b *VXAutocompleteBuilder) SetRemoteResource(res *RemoteResource) (r *VXAutocompleteBuilder) {
	b.tag.Attr("remote-res", res.Name)
	b.tag.Attr("remote-url", res.RemoteURL)
	b.tag.Attr("is-paging", res.IsPaging)
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
	Model interface{}                          `json:"-"`
	DB    *gorm.DB                             `json:"-"`
	NewDB func(ctx *web.EventContext) *gorm.DB `json:"-"`

	Name         string                   `json:"name"`
	RemoteURL    string                   `json:"remote_url"`
	IsPaging     bool                     `json:"is_paging"`
	SizePerPage  int                      `json:"-"`
	SearchFields []string                 `json:"-"`
	OptionIcon   func(interface{}) string `json:"-"`
	OptionText   func(interface{}) string `json:"-"`
	OptionValue  string                   `json:"-"`
}

func RegisterRemoteResource(res *RemoteResource) *RemoteResource {
	if res.Model == nil {
		panic("model is required")
	}

	if reflect.TypeOf(res.Model).Kind() != reflect.Ptr {
		panic("model must be a pointer")
	}

	if res.DB == nil && res.NewDB == nil {
		panic("db is required")
	}

	if len(res.SearchFields) == 0 {
		res.SearchFields = []string{"ID"}
	}

	if res.OptionValue == "" {
		res.OptionValue = "ID"
	}

	if res.SizePerPage == 0 {
		res.SizePerPage = 20
	}

	if res.Name == "" {
		res.Name = reflect.ValueOf(res.Model).Elem().Type().Name()
	}

	if res.RemoteURL == "" {
		res.RemoteURL = "/admin"
	}

	if remoteResources[res.Name] != nil {
		panic(fmt.Sprintf("duplicate remote resource name: %s", res.Name))
	}

	remoteResources[res.Name] = res
	return res
}

type RemoteResourceResult struct {
	Items   []OptionItem `json:"items"`
	Total   int          `json:"total"`
	Current int          `json:"current"`
}

type OptionItem struct {
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
	Icon  string `json:"icon,omitempty"`
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

	var datas = reflect.MakeSlice(reflect.SliceOf(reflect.Indirect(reflect.ValueOf(res.Model)).Type()), 0, 0).Interface()
	db.Model(res.Model).Offset(offset).Limit(res.SizePerPage).Find(&datas)

	reflectValue := reflect.Indirect(reflect.ValueOf(datas))
	var items []OptionItem
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

		items = append(items, OptionItem{
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

	r.Data = RemoteResourceResult{
		Items:   items,
		Total:   total,
		Current: current,
	}
	return
}
