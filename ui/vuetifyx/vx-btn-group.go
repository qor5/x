package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXBtnGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXBtnGroup(children ...h.HTMLComponent) (r *VXBtnGroupBuilder) {
	r = &VXBtnGroupBuilder{
		tag: h.Tag("vx-btn-group").Children(children...),
	}
	return
}

func (b *VXBtnGroupBuilder) BaseColor(v string) (r *VXBtnGroupBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VXBtnGroupBuilder) DividerColor(v string) (r *VXBtnGroupBuilder) {
	b.tag.Attr("divider-color", v)
	return b
}

func (b *VXBtnGroupBuilder) DividerWidth(v interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":divider-color", h.JSONString(v))
	return b
}

func (b *VXBtnGroupBuilder) Divided(v bool) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":divided", fmt.Sprint(v))
	return b
}

func (b *VXBtnGroupBuilder) Border(v interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VXBtnGroupBuilder) Density(v interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXBtnGroupBuilder) Elevation(v interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VXBtnGroupBuilder) Rounded(v interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXBtnGroupBuilder) Tile(v bool) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXBtnGroupBuilder) Tag(v string) (r *VXBtnGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXBtnGroupBuilder) Theme(v string) (r *VXBtnGroupBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXBtnGroupBuilder) Color(v string) (r *VXBtnGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXBtnGroupBuilder) Variant(v interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXBtnGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXBtnGroupBuilder) Attr(vs ...interface{}) (r *VXBtnGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXBtnGroupBuilder) Children(children ...h.HTMLComponent) (r *VXBtnGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXBtnGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXBtnGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXBtnGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXBtnGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXBtnGroupBuilder) Class(names ...string) (r *VXBtnGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXBtnGroupBuilder) ClassIf(name string, add bool) (r *VXBtnGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXBtnGroupBuilder) On(name string, value string) (r *VXBtnGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXBtnGroupBuilder) Bind(name string, value string) (r *VXBtnGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXBtnGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
