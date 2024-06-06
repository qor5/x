package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDividerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDivider(children ...h.HTMLComponent) (r *VDividerBuilder) {
	r = &VDividerBuilder{
		tag: h.Tag("v-divider").Children(children...),
	}
	return
}

func (b *VDividerBuilder) Length(v interface{}) (r *VDividerBuilder) {
	b.tag.Attr(":length", h.JSONString(v))
	return b
}

func (b *VDividerBuilder) Color(v string) (r *VDividerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDividerBuilder) Inset(v bool) (r *VDividerBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VDividerBuilder) Opacity(v interface{}) (r *VDividerBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VDividerBuilder) Thickness(v interface{}) (r *VDividerBuilder) {
	b.tag.Attr(":thickness", h.JSONString(v))
	return b
}

func (b *VDividerBuilder) Vertical(v bool) (r *VDividerBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VDividerBuilder) Theme(v string) (r *VDividerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDividerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDividerBuilder) Attr(vs ...interface{}) (r *VDividerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDividerBuilder) Children(children ...h.HTMLComponent) (r *VDividerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDividerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDividerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDividerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDividerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDividerBuilder) Class(names ...string) (r *VDividerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDividerBuilder) ClassIf(name string, add bool) (r *VDividerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDividerBuilder) On(name string, value string) (r *VDividerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDividerBuilder) Bind(name string, value string) (r *VDividerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDividerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
