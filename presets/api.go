package presets

import (
	"mime/multipart"
	"net/url"

	"github.com/sunfmin/bran/ui"
	v "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

// UI Layer

type ComponentFunc func(ctx *ui.EventContext) h.HTMLComponent

type BulkComponentFunc func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent

type FieldComponentFunc func(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent

type FilterDataFunc func(ctx *ui.EventContext) v.FilterData

type FilterTab struct {
	Label string
	Query url.Values
}

type FilterTabsFunc func(ctx *ui.EventContext) []*FilterTab

type BulkActionUpdateFunc func(selectedIds []string, form *multipart.Form, ctx *ui.EventContext) (err error)

type UpdateFunc func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) (err error)

type SetterFunc func(obj interface{}, form *multipart.Form, ctx *ui.EventContext)

type MessagesFunc func(ctx *ui.EventContext) *Messages

// Data Layer
type DataOperator interface {
	Search(obj interface{}, params *SearchParams) (r interface{}, totalCount int, err error)
	Fetch(obj interface{}, id string) (r interface{}, err error)
	UpdateField(obj interface{}, id string, fieldName string, value interface{}) (err error)
	Save(obj interface{}, id string) (err error)
	Delete(obj interface{}, id string) (err error)
}

type SearchOpFunc func(model interface{}, params *SearchParams) (r interface{}, totalCount int, err error)
type FetchOpFunc func(obj interface{}, id string) (r interface{}, err error)
type UpdateFieldOpFunc func(obj interface{}, id string, fieldName string, value interface{}) (err error)
type SaveOpFunc func(obj interface{}, id string) (err error)
type DeleteOpFunc func(obj interface{}, id string) (err error)

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
