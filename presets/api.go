package presets

import (
	"mime/multipart"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

// UI Layer

type ComponentFunc func(ctx *ui.EventContext) h.HTMLComponent

type BulkComponentFunc func(selectedIds []string, ctx *ui.EventContext) h.HTMLComponent

type FieldComponentFunc func(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent

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
}

type SearchOpFunc func(model interface{}, params *SearchParams) (r interface{}, totalCount int, err error)
type FetchOpFunc func(obj interface{}, id string) (r interface{}, err error)
type UpdateFieldOpFunc func(obj interface{}, id string, fieldName string, value interface{}) (err error)
type SaveOpFunc func(obj interface{}, id string) (err error)

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
}
