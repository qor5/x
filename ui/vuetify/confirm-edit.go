package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VConfirmEditBuilder struct {
	tag *h.HTMLTagBuilder
}

func VConfirmEdit(children ...h.HTMLComponent) (r *VConfirmEditBuilder) {
	r = &VConfirmEditBuilder{
		tag: h.Tag("v-confirm-edit").Children(children...),
	}
	return
}

func (b *VConfirmEditBuilder) ModelValue(v interface{}) (r *VConfirmEditBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VConfirmEditBuilder) Color(v string) (r *VConfirmEditBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VConfirmEditBuilder) CancelText(v string) (r *VConfirmEditBuilder) {
	b.tag.Attr("cancel-text", v)
	return b
}

func (b *VConfirmEditBuilder) OkText(v string) (r *VConfirmEditBuilder) {
	b.tag.Attr("ok-text", v)
	return b
}

func (b *VConfirmEditBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VConfirmEditBuilder) Attr(vs ...interface{}) (r *VConfirmEditBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VConfirmEditBuilder) Children(children ...h.HTMLComponent) (r *VConfirmEditBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VConfirmEditBuilder) AppendChildren(children ...h.HTMLComponent) (r *VConfirmEditBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VConfirmEditBuilder) PrependChildren(children ...h.HTMLComponent) (r *VConfirmEditBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VConfirmEditBuilder) Class(names ...string) (r *VConfirmEditBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VConfirmEditBuilder) ClassIf(name string, add bool) (r *VConfirmEditBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VConfirmEditBuilder) On(name string, value string) (r *VConfirmEditBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VConfirmEditBuilder) Bind(name string, value string) (r *VConfirmEditBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VConfirmEditBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
