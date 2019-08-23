package presets

import (
	"log"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type FieldBuilder struct {
	NameLabel
	compFunc    FieldComponentFunc
	setterFunc  SetterFunc
	inplaceEdit *InplaceEditBuilder
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

func (b *FieldBuilder) InplaceEdit() (r *InplaceEditBuilder) {
	r = &InplaceEditBuilder{}
	b.inplaceEdit = r
	return
}

type InplaceEditBuilder struct {
	compFunc   FieldComponentFunc
	updateFunc UpdateFunc
}

func (b *InplaceEditBuilder) ComponentFunc(v FieldComponentFunc) (r *InplaceEditBuilder) {
	b.compFunc = v
	return b
}

func (b *InplaceEditBuilder) UpdateFunc(v UpdateFunc) (r *InplaceEditBuilder) {
	b.updateFunc = v
	return b
}

type NameLabel struct {
	name  string
	label string
}
