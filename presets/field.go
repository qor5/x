package presets

import (
	"log"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type FieldBuilder struct {
	NameLabel
	compFunc   FieldComponentFunc
	setterFunc SetterFunc
}

func NewField(name string) (r *FieldBuilder) {
	r = &FieldBuilder{}
	r.name = name
	r.compFunc = emptyComponentFunc
	return
}

func emptyComponentFunc(obj interface{}, field *FieldContext, ctx *ui.EventContext) (r h.HTMLComponent) {
	log.Printf("No ComponentFunc for field %v\n", field.Name)
	return
}

func (b *FieldBuilder) Label(v string) (r *FieldBuilder) {
	b.label = v
	return b
}

func (b *FieldBuilder) Clone() (r *FieldBuilder) {
	r = &FieldBuilder{}
	r.name = b.name
	r.label = b.label
	r.compFunc = b.compFunc
	r.setterFunc = b.setterFunc
	return r
}

func (b *FieldBuilder) ComponentFunc(v FieldComponentFunc) (r *FieldBuilder) {
	if v == nil {
		panic("value required")
	}
	b.compFunc = v
	return b
}

func (b *FieldBuilder) SetterFunc(v SetterFunc) (r *FieldBuilder) {
	b.setterFunc = v
	return b
}

type NameLabel struct {
	name  string
	label string
}

type FieldBuilders struct {
	fields []*FieldBuilder
}

func (b *FieldBuilders) Field(name string) (r *FieldBuilder) {
	r = b.GetField(name)
	if r != nil {
		return
	}

	r = NewField(name)
	b.fields = append(b.fields, r)
	return
}

func (b *FieldBuilders) GetField(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	return
}
