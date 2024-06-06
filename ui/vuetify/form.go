package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFormBuilder struct {
	tag *h.HTMLTagBuilder
}

func VForm(children ...h.HTMLComponent) (r *VFormBuilder) {
	r = &VFormBuilder{
		tag: h.Tag("v-form").Children(children...),
	}
	return
}

func (b *VFormBuilder) ModelValue(v bool) (r *VFormBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VFormBuilder) Disabled(v bool) (r *VFormBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VFormBuilder) FastFail(v bool) (r *VFormBuilder) {
	b.tag.Attr(":fast-fail", fmt.Sprint(v))
	return b
}

func (b *VFormBuilder) Readonly(v bool) (r *VFormBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VFormBuilder) ValidateOn(v interface{}) (r *VFormBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VFormBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFormBuilder) Attr(vs ...interface{}) (r *VFormBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFormBuilder) Children(children ...h.HTMLComponent) (r *VFormBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFormBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFormBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFormBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFormBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFormBuilder) Class(names ...string) (r *VFormBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFormBuilder) ClassIf(name string, add bool) (r *VFormBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFormBuilder) On(name string, value string) (r *VFormBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFormBuilder) Bind(name string, value string) (r *VFormBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFormBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
