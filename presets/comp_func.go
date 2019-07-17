package presets

import (
	"fmt"
	"reflect"

	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
)

type Field struct {
	Name  string
	Label string
}

type CompFunc func(obj interface{}, field *Field, ctx *ui.EventContext) h.HTMLComponent

type FieldTypeBuilder struct {
	valType           reflect.Type
	listingCompFunc   CompFunc
	detailingCompFunc CompFunc
	editingCompFunc   CompFunc
}

func NewFieldType(t reflect.Type) (r *FieldTypeBuilder) {
	r = &FieldTypeBuilder{valType: t}
	return
}

func (b *FieldTypeBuilder) ListingComponentFunc(v CompFunc) (r *FieldTypeBuilder) {
	b.listingCompFunc = v
	return b
}

func (b *FieldTypeBuilder) EditingComponentFunc(v CompFunc) (r *FieldTypeBuilder) {
	b.editingCompFunc = v
	return b
}

func (b *FieldTypeBuilder) DetailingComponentFunc(v CompFunc) (r *FieldTypeBuilder) {
	b.detailingCompFunc = v
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
	fieldTypes []*FieldTypeBuilder
}

func (b *FieldTypes) FieldType(v interface{}) (r *FieldTypeBuilder) {
	return b.fieldTypeByType(reflect.TypeOf(v))
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

func stringVal(obj interface{}, fieldName string) (r string) {
	val := reflectutils.MustGet(obj, fieldName)
	switch vt := val.(type) {
	case []rune:
		return string(vt)
	case []byte:
		return string(vt)
	}
	return fmt.Sprint(val)
}

func cfText(obj interface{}, field *Field, ctx *ui.EventContext) h.HTMLComponent {
	return h.Text(stringVal(obj, field.Name))
}

func cfCheckbox(obj interface{}, field *Field, ctx *ui.EventContext) h.HTMLComponent {
	return VCheckbox().FieldName(field.Name).Label(field.Label).Value(stringVal(obj, field.Name))
}

func cfNumber(obj interface{}, field *Field, ctx *ui.EventContext) h.HTMLComponent {
	return VTextField().Type("number").FieldName(field.Name).Label(field.Label).Value(stringVal(obj, field.Name))
}

func cfTextField(obj interface{}, field *Field, ctx *ui.EventContext) h.HTMLComponent {
	return VTextField().Type("text").FieldName(field.Name).Label(field.Label).Value(stringVal(obj, field.Name))
}

func builtInFieldTypes() (r FieldTypes) {
	r.FieldType(true).ListingComponentFunc(cfText).DetailingComponentFunc(cfText).EditingComponentFunc(cfCheckbox)

	for _, v := range numberVals {
		r.FieldType(v).ListingComponentFunc(cfText).DetailingComponentFunc(cfText).EditingComponentFunc(cfNumber)
	}

	for _, v := range stringVals {
		r.FieldType(v).ListingComponentFunc(cfText).DetailingComponentFunc(cfText).EditingComponentFunc(cfTextField)
	}
	return
}
