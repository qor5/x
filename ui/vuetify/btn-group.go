package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBtnGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBtnGroup(children ...h.HTMLComponent) (r *VBtnGroupBuilder) {
	r = &VBtnGroupBuilder{
		tag: h.Tag("v-btn-group").Children(children...),
	}
	return
}

func (b *VBtnGroupBuilder) BaseColor(v string) (r *VBtnGroupBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VBtnGroupBuilder) Divided(v bool) (r *VBtnGroupBuilder) {
	b.tag.Attr(":divided", fmt.Sprint(v))
	return b
}

func (b *VBtnGroupBuilder) Border(v interface{}) (r *VBtnGroupBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VBtnGroupBuilder) Density(v interface{}) (r *VBtnGroupBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VBtnGroupBuilder) Elevation(v interface{}) (r *VBtnGroupBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VBtnGroupBuilder) Rounded(v interface{}) (r *VBtnGroupBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VBtnGroupBuilder) Tile(v bool) (r *VBtnGroupBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBtnGroupBuilder) Tag(v string) (r *VBtnGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBtnGroupBuilder) Theme(v string) (r *VBtnGroupBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VBtnGroupBuilder) Color(v string) (r *VBtnGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnGroupBuilder) Variant(v interface{}) (r *VBtnGroupBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VBtnGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBtnGroupBuilder) Attr(vs ...interface{}) (r *VBtnGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBtnGroupBuilder) Children(children ...h.HTMLComponent) (r *VBtnGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBtnGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBtnGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBtnGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBtnGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBtnGroupBuilder) Class(names ...string) (r *VBtnGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBtnGroupBuilder) ClassIf(name string, add bool) (r *VBtnGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBtnGroupBuilder) On(name string, value string) (r *VBtnGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBtnGroupBuilder) Bind(name string, value string) (r *VBtnGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBtnGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
