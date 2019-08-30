package presets

import (
	"net/http"
	"net/url"

	"github.com/sunfmin/bran/ui"
	v "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type ComponentFunc func(ctx *ui.EventContext) h.HTMLComponent
type FieldComponentFunc func(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent

type ActionComponentFunc func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent
type ActionUpdateFunc func(selectedIds []string, ctx *ui.EventContext) (err error)

type MessagesFunc func(r *http.Request) *Messages

// Data Layer
type DataOperator interface {
	Search(obj interface{}, params *SearchParams, ctx *ui.EventContext) (r interface{}, totalCount int, err error)
	Fetch(obj interface{}, id string, ctx *ui.EventContext) (r interface{}, err error)
	Save(obj interface{}, id string, ctx *ui.EventContext) (err error)
	Delete(obj interface{}, id string, ctx *ui.EventContext) (err error)
}

type SetterFunc func(obj interface{}, ctx *ui.EventContext)
type ValidateFunc func(obj interface{}, ctx *ui.EventContext) (err ValidationErrors)

type SearchFunc func(model interface{}, params *SearchParams, ctx *ui.EventContext) (r interface{}, totalCount int, err error)
type FetchFunc func(obj interface{}, id string, ctx *ui.EventContext) (r interface{}, err error)
type SaveFunc func(obj interface{}, id string, ctx *ui.EventContext) (err error)
type DeleteFunc func(obj interface{}, id string, ctx *ui.EventContext) (err error)

type FilterDataFunc func(ctx *ui.EventContext) v.FilterData

type FilterTab struct {
	Label string
	Query url.Values
}

type FilterTabsFunc func(ctx *ui.EventContext) []*FilterTab

type SQLCondition struct {
	Query string
	Args  []interface{}
}

type SearchParams struct {
	KeywordColumns []string
	Keyword        string
	SQLConditions  []*SQLCondition
	PerPage        int64
	Page           int64
	OrderBy        string
}
