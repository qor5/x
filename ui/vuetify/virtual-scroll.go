package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VVirtualScrollBuilder struct {
	tag *h.HTMLTagBuilder
}

func VVirtualScroll(children ...h.HTMLComponent) (r *VVirtualScrollBuilder) {
	r = &VVirtualScrollBuilder{
		tag: h.Tag("v-virtual-scroll").Children(children...),
	}
	return
}

func (b *VVirtualScrollBuilder) Items(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) Renderless(v bool) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":renderless", fmt.Sprint(v))
	return b
}

func (b *VVirtualScrollBuilder) ItemHeight(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":item-height", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) Height(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) MaxHeight(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) MaxWidth(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) MinHeight(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) MinWidth(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) Width(v interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VVirtualScrollBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VVirtualScrollBuilder) Attr(vs ...interface{}) (r *VVirtualScrollBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VVirtualScrollBuilder) Children(children ...h.HTMLComponent) (r *VVirtualScrollBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VVirtualScrollBuilder) AppendChildren(children ...h.HTMLComponent) (r *VVirtualScrollBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VVirtualScrollBuilder) PrependChildren(children ...h.HTMLComponent) (r *VVirtualScrollBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VVirtualScrollBuilder) Class(names ...string) (r *VVirtualScrollBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VVirtualScrollBuilder) ClassIf(name string, add bool) (r *VVirtualScrollBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VVirtualScrollBuilder) On(name string, value string) (r *VVirtualScrollBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VVirtualScrollBuilder) Bind(name string, value string) (r *VVirtualScrollBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VVirtualScrollBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
