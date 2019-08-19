package presets

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type FieldContext struct {
	Name      string
	Label     string
	ModelInfo *ModelInfo
}

func (fc *FieldContext) StringValue(obj interface{}) (r string) {
	fieldName := fc.Name
	val := reflectutils.MustGet(obj, fieldName)
	switch vt := val.(type) {
	case []rune:
		return string(vt)
	case []byte:
		return string(vt)
	}
	return fmt.Sprint(val)
}

type FieldTypeBuilder struct {
	valType           reflect.Type
	listingCompFunc   FieldComponentFunc
	detailingCompFunc FieldComponentFunc
	editingCompFunc   FieldComponentFunc
}

type PageType string

const (
	LISTING   PageType = "listing"
	EDITING   PageType = "editing"
	DETAILING PageType = "detailing"
)

func NewFieldType(t reflect.Type) (r *FieldTypeBuilder) {
	r = &FieldTypeBuilder{valType: t}
	return
}

func (b *FieldTypeBuilder) ComponentFunc(t PageType, v FieldComponentFunc) (r *FieldTypeBuilder) {
	switch t {
	case LISTING:
		b.listingCompFunc = v
	case EDITING:
		b.editingCompFunc = v
	case DETAILING:
		b.detailingCompFunc = v
	}
	return b
}

var numberVals = []interface{}{
	int(0), int8(0), int16(0), int32(0), int64(0),
	uint(0), uint(8), uint16(0), uint32(0), uint64(0),
	float32(0.0), float64(0.0),
}

var stringVals = []interface{}{
	string(""),
	[]rune(""),
	[]byte(""),
}

type FieldTypes struct {
	fieldTypes                []*FieldTypeBuilder
	excludesListingPatterns   []string
	excludesEditingPatterns   []string
	excludesDetailingPatterns []string
}

func (b *FieldTypes) FieldType(v interface{}) (r *FieldTypeBuilder) {
	return b.fieldTypeByType(reflect.TypeOf(v))
}

func (b *FieldTypes) ExcludeFields(t PageType, patterns ...string) {
	switch t {
	case LISTING:
		b.excludesListingPatterns = patterns
	case EDITING:
		b.excludesEditingPatterns = patterns
	case DETAILING:
		b.excludesDetailingPatterns = patterns
	}
}

func hasMatched(patterns []string, name string) bool {
	for _, p := range patterns {
		ok, err := filepath.Match(p, name)
		if err != nil {
			panic(err)
		}
		if ok {
			return true
		}
	}
	return false
}

func (b *FieldTypes) fieldNameExcluded(t PageType, name string) bool {
	switch t {
	case LISTING:
		return hasMatched(b.excludesListingPatterns, name)
	case EDITING:
		return hasMatched(b.excludesEditingPatterns, name)
	case DETAILING:
		return hasMatched(b.excludesDetailingPatterns, name)
	}
	return false
}

func (b *FieldTypes) fieldTypeByType(tv reflect.Type) (r *FieldTypeBuilder) {
	for _, ft := range b.fieldTypes {
		if ft.valType == tv {
			return ft
		}
	}
	r = NewFieldType(tv)
	b.fieldTypes = append(b.fieldTypes, r)
	return
}

func cfText(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent {
	return h.Text(field.StringValue(obj))
}

func cfTextTd(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent {
	if field.Name == "ID" {
		id := field.StringValue(obj)
		if len(id) > 0 {
			a := ui.Bind(h.A().Text(id))
			if field.ModelInfo.HasDetailing() {
				a.PushStateURL(
					field.ModelInfo.DetailingHref(id),
				)
			} else {
				a.OnClick("formDrawerEdit", id)
			}
			return h.Td(a)
		}
	}
	return h.Td(h.Text(field.StringValue(obj)))
}

func cfCheckbox(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent {
	return VCheckbox().
		FieldName(field.Name).
		Label(field.Label).
		InputValue(reflectutils.MustGet(obj, field.Name).(bool))
}

func cfNumber(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent {
	return VTextField().
		Type("number").
		FieldName(field.Name).
		Label(field.Label).
		Value(fmt.Sprint(reflectutils.MustGet(obj, field.Name)))
}

func cfTextField(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent {
	return VTextField().
		Type("text").
		FieldName(field.Name).
		Label(field.Label).
		Value(reflectutils.MustGet(obj, field.Name).(string))
}

func builtInFieldTypes() (r FieldTypes) {
	r.FieldType(true).
		ComponentFunc(LISTING, cfTextTd).
		ComponentFunc(DETAILING, cfText).
		ComponentFunc(EDITING, cfCheckbox)

	for _, v := range numberVals {
		r.FieldType(v).
			ComponentFunc(LISTING, cfTextTd).
			ComponentFunc(DETAILING, cfText).
			ComponentFunc(EDITING, cfNumber)
	}

	for _, v := range stringVals {
		r.FieldType(v).
			ComponentFunc(LISTING, cfTextTd).
			ComponentFunc(DETAILING, cfText).
			ComponentFunc(EDITING, cfTextField)
	}

	r.ExcludeFields(EDITING, "ID")
	return
}
