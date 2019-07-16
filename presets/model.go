package presets

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/sunfmin/bran/vuetify"

	"github.com/sunfmin/reflectutils"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type ModelBuilder struct {
	p             *Builder
	model         interface{}
	uriName       string
	labels        []string
	placeholders  []string
	searchColumns []string
	listing       *ListingBuilder
	editing       *EditingBuilder
	detailing     *DetailingBuilder
}

func (b *ModelBuilder) inspectModel() {
	v := reflect.ValueOf(b.model)

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v = v.Elem()

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type.String())
	}
}

var basicTypes = strings.Fields(`
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte
rune
float32 float64
complex64 complex128
`)

var numberTypes = strings.Fields(`
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64
float32 float64
`)

var stringTypes = strings.Fields(`
string
[]rune
[]byte
`)

func (b *ModelBuilder) newListing() (r *ListingBuilder) {
	r = &ListingBuilder{filtering: &FilteringBuilder{}}
	r.RegisterComponentFunc(basicTypes, func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
		val, _ := reflectutils.Get(obj, fieldName)
		return h.Text(fmt.Sprint(val))
	}, true)
	return
}

func (b *ModelBuilder) newEditing() (r *EditingBuilder) {
	r = &EditingBuilder{}

	r.RegisterComponentFunc(numberTypes, func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
		//val, _ := reflectutils.Get(obj, fieldName)
		return VTextField().Type("number").FieldName(fieldName)
	}, true)

	r.RegisterComponentFunc(stringTypes, func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent {
		//val, _ := reflectutils.Get(obj, fieldName)
		return VTextField().Type("text").FieldName(fieldName)
	}, true)

	return
}

func (b *ModelBuilder) newDetailing() (r *DetailingBuilder) {
	r = &DetailingBuilder{}
	return
}

func (b *ModelBuilder) URIName(v string) (r *ModelBuilder) {
	b.uriName = v
	return b
}

func (b *ModelBuilder) Labels(vs ...string) (r *ModelBuilder) {
	b.labels = append(b.labels, vs...)
	return b
}

func (b *ModelBuilder) Placeholders(vs ...string) (r *ModelBuilder) {
	b.placeholders = append(b.placeholders, vs...)
	return b
}

func (b *ModelBuilder) SearchColumns(vs ...string) (r *ModelBuilder) {
	b.searchColumns = vs
	return b
}
