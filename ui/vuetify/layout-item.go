package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLayoutItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLayoutItem(children ...h.HTMLComponent) (r *VLayoutItemBuilder) {
	r = &VLayoutItemBuilder{
		tag: h.Tag("v-layout-item").Children(children...),
	}
	return
}

func (b *VLayoutItemBuilder) Position(v interface{}) (r *VLayoutItemBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VLayoutItemBuilder) Size(v interface{}) (r *VLayoutItemBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VLayoutItemBuilder) ModelValue(v bool) (r *VLayoutItemBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VLayoutItemBuilder) Name(v string) (r *VLayoutItemBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VLayoutItemBuilder) Order(v interface{}) (r *VLayoutItemBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VLayoutItemBuilder) Absolute(v bool) (r *VLayoutItemBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VLayoutItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VLayoutItemBuilder) Attr(vs ...interface{}) (r *VLayoutItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VLayoutItemBuilder) Children(children ...h.HTMLComponent) (r *VLayoutItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VLayoutItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VLayoutItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VLayoutItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VLayoutItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VLayoutItemBuilder) Class(names ...string) (r *VLayoutItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLayoutItemBuilder) ClassIf(name string, add bool) (r *VLayoutItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLayoutItemBuilder) On(name string, value string) (r *VLayoutItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLayoutItemBuilder) Bind(name string, value string) (r *VLayoutItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLayoutItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
