package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VItem(children ...h.HTMLComponent) (r *VItemBuilder) {
	r = &VItemBuilder{
		tag: h.Tag("v-item").Children(children...),
	}
	return
}

func (b *VItemBuilder) Value(v interface{}) (r *VItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VItemBuilder) Disabled(v bool) (r *VItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VItemBuilder) SelectedClass(v string) (r *VItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VItemBuilder) Attr(vs ...interface{}) (r *VItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VItemBuilder) Children(children ...h.HTMLComponent) (r *VItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VItemBuilder) Class(names ...string) (r *VItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VItemBuilder) ClassIf(name string, add bool) (r *VItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VItemBuilder) On(name string, value string) (r *VItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VItemBuilder) Bind(name string, value string) (r *VItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
