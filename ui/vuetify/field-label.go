package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFieldLabelBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFieldLabel(children ...h.HTMLComponent) (r *VFieldLabelBuilder) {
	r = &VFieldLabelBuilder{
		tag: h.Tag("v-field-label").Children(children...),
	}
	return
}

func (b *VFieldLabelBuilder) Floating(v bool) (r *VFieldLabelBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VFieldLabelBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFieldLabelBuilder) Attr(vs ...interface{}) (r *VFieldLabelBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFieldLabelBuilder) Children(children ...h.HTMLComponent) (r *VFieldLabelBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFieldLabelBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFieldLabelBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFieldLabelBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFieldLabelBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFieldLabelBuilder) Class(names ...string) (r *VFieldLabelBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFieldLabelBuilder) ClassIf(name string, add bool) (r *VFieldLabelBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFieldLabelBuilder) On(name string, value string) (r *VFieldLabelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFieldLabelBuilder) Bind(name string, value string) (r *VFieldLabelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFieldLabelBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
