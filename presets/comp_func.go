package presets

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/iancoleman/strcase"

	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type FieldContext struct {
	Name   string
	Label  string
	Errors []string
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
	valType  reflect.Type
	mode     FieldMode
	compFunc FieldComponentFunc
}

type FieldMode int

const (
	WRITE FieldMode = iota
	LIST
)

func NewFieldType(t reflect.Type) (r *FieldTypeBuilder) {
	r = &FieldTypeBuilder{valType: t}
	return
}

func (b *FieldTypeBuilder) ComponentFunc(v FieldComponentFunc) (r *FieldTypeBuilder) {
	b.compFunc = v
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
	mode             FieldMode
	fieldTypes       []*FieldTypeBuilder
	excludesPatterns []string
}

func NewFieldTypes(t FieldMode) (r *FieldTypes) {
	r = &FieldTypes{
		mode: t,
	}
	r.builtInFieldTypes()
	return
}

func (b *FieldTypes) FieldType(v interface{}) (r *FieldTypeBuilder) {
	return b.fieldTypeByType(reflect.TypeOf(v))
}

func (b *FieldTypes) Exclude(patterns ...string) (r *FieldTypes) {
	b.excludesPatterns = patterns
	return b
}

func (b *FieldTypes) InspectFields(val interface{}) (r *FieldBuilders) {
	r, _ = b.inspectFieldsAndCollectName(val, nil)
	return
}

func (b *FieldTypes) inspectFieldsAndCollectName(val interface{}, collectType reflect.Type) (r *FieldBuilders, names []string) {
	v := reflect.ValueOf(val)

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v = v.Elem()

	t := v.Type()

	r = &FieldBuilders{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		ft := b.fieldTypeByType(f.Type)

		if !hasMatched(b.excludesPatterns, f.Name) && ft.compFunc != nil {
			r.Field(f.Name).ComponentFunc(ft.compFunc)
		}

		if collectType != nil && f.Type == collectType {
			names = append(names, strcase.ToSnake(f.Name))
		}
	}

	return
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

func cfTextTd(obj interface{}, field *FieldContext, ctx *ui.EventContext) h.HTMLComponent {
	if field.Name == "ID" {
		id := field.StringValue(obj)
		if len(id) > 0 {
			mi := GetModelInfo(ctx.R)
			if mi == nil {
				return h.Td().Text(id)
			}

			a := ui.Bind(h.A().Text(id))
			if mi.HasDetailing() {
				a.PushStateURL(
					mi.DetailingHref(id),
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

func (b *FieldTypes) builtInFieldTypes() {

	if b.mode == LIST {
		b.FieldType(true).
			ComponentFunc(cfTextTd)

		for _, v := range numberVals {
			b.FieldType(v).
				ComponentFunc(cfTextTd)
		}

		for _, v := range stringVals {
			b.FieldType(v).
				ComponentFunc(cfTextTd)
		}
		return
	}

	b.FieldType(true).
		ComponentFunc(cfCheckbox)

	for _, v := range numberVals {
		b.FieldType(v).
			ComponentFunc(cfNumber)
	}

	for _, v := range stringVals {
		b.FieldType(v).
			ComponentFunc(cfTextField)
	}

	b.Exclude("ID")
	return
}
