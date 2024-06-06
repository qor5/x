package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTableBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VTableBuilder) FixedHeader(v bool) (r *VTableBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VTableBuilder) FixedFooter(v bool) (r *VTableBuilder) {
	b.tag.Attr(":fixed-footer", fmt.Sprint(v))
	return b
}

func (b *VTableBuilder) Height(v interface{}) (r *VTableBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTableBuilder) Hover(v bool) (r *VTableBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VTableBuilder) Density(v interface{}) (r *VTableBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTableBuilder) Tag(v string) (r *VTableBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTableBuilder) Theme(v string) (r *VTableBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTableBuilder) Attr(vs ...interface{}) (r *VTableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTableBuilder) Children(children ...h.HTMLComponent) (r *VTableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTableBuilder) Class(names ...string) (r *VTableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTableBuilder) ClassIf(name string, add bool) (r *VTableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTableBuilder) On(name string, value string) (r *VTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTableBuilder) Bind(name string, value string) (r *VTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
