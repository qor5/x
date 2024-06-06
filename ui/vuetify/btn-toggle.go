package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBtnToggleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBtnToggle(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	r = &VBtnToggleBuilder{
		tag: h.Tag("v-btn-toggle").Children(children...),
	}
	return
}

func (b *VBtnToggleBuilder) BaseColor(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VBtnToggleBuilder) Divided(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":divided", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Border(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) Density(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) Elevation(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) Rounded(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) Tile(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Tag(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBtnToggleBuilder) Theme(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VBtnToggleBuilder) Color(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnToggleBuilder) Variant(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) ModelValue(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) Multiple(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Max(v int) (r *VBtnToggleBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) SelectedClass(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VBtnToggleBuilder) Disabled(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Mandatory(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VBtnToggleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBtnToggleBuilder) Attr(vs ...interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBtnToggleBuilder) Children(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBtnToggleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBtnToggleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBtnToggleBuilder) Class(names ...string) (r *VBtnToggleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBtnToggleBuilder) ClassIf(name string, add bool) (r *VBtnToggleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBtnToggleBuilder) On(name string, value string) (r *VBtnToggleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBtnToggleBuilder) Bind(name string, value string) (r *VBtnToggleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBtnToggleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
