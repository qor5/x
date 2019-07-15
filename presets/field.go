package presets

import (
	"mime/multipart"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type FieldBuilder struct {
	name        string
	label       string
	compFunc    CompFunc
	setterFunc  SetterFunc
	inplaceEdit *InplaceEditBuilder
}

type CompFunc func(obj interface{}, ctx *ui.EventContext) h.HTMLComponent
type UpdateFunc func(obj interface{}, form *multipart.Form, ctx *ui.EventContext) (err error)

func (b *FieldBuilder) Label(v string) (r *FieldBuilder) {
	b.label = v
	return b
}

func (b *FieldBuilder) ComponentFunc(v CompFunc) (r *FieldBuilder) {
	b.compFunc = v
	return b
}

type SetterFunc func(obj interface{}, form *multipart.Form, ctx *ui.EventContext)

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
	compFunc   CompFunc
	updateFunc UpdateFunc
}

func (b *InplaceEditBuilder) ComponentFunc(v CompFunc) (r *InplaceEditBuilder) {
	b.compFunc = v
	return b
}

func (b *InplaceEditBuilder) UpdateFunc(v UpdateFunc) (r *InplaceEditBuilder) {
	b.updateFunc = v
	return b
}
