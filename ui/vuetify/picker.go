package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VPickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VPicker(children ...h.HTMLComponent) (r *VPickerBuilder) {
	r = &VPickerBuilder{
		tag: h.Tag("v-picker").Children(children...),
	}
	return
}

func (b *VPickerBuilder) Title(v string) (r *VPickerBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VPickerBuilder) BgColor(v string) (r *VPickerBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VPickerBuilder) Landscape(v bool) (r *VPickerBuilder) {
	b.tag.Attr(":landscape", fmt.Sprint(v))
	return b
}

func (b *VPickerBuilder) HideHeader(v bool) (r *VPickerBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VPickerBuilder) Color(v string) (r *VPickerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VPickerBuilder) Border(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Height(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) MaxHeight(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) MaxWidth(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) MinHeight(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) MinWidth(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Width(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Elevation(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Location(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Position(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Rounded(v interface{}) (r *VPickerBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VPickerBuilder) Tile(v bool) (r *VPickerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VPickerBuilder) Tag(v string) (r *VPickerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VPickerBuilder) Theme(v string) (r *VPickerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VPickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VPickerBuilder) Attr(vs ...interface{}) (r *VPickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VPickerBuilder) Children(children ...h.HTMLComponent) (r *VPickerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VPickerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VPickerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VPickerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VPickerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VPickerBuilder) Class(names ...string) (r *VPickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VPickerBuilder) ClassIf(name string, add bool) (r *VPickerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VPickerBuilder) On(name string, value string) (r *VPickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VPickerBuilder) Bind(name string, value string) (r *VPickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VPickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
