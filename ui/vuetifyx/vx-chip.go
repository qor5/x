package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXChipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXChip(text string) (r *VXChipBuilder) {
	r = &VXChipBuilder{
		tag: h.Tag("vx-chip").Children(h.Text(text)),
	}
	return
}

func (b *VXChipBuilder) Label(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":label", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Filter(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":filter", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Round(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":round", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Presets(v string) (r *VXChipBuilder) {
	b.tag.Attr("presets", v)
	return b
}

func (b *VXChipBuilder) ActiveClass(v string) (r *VXChipBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VXChipBuilder) AppendAvatar(v string) (r *VXChipBuilder) {
	b.tag.Attr("append-avatar", v)
	return b
}

func (b *VXChipBuilder) AppendIcon(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Closable(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":closable", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) CloseIcon(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":close-icon", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) CloseLabel(v string) (r *VXChipBuilder) {
	b.tag.Attr("close-label", v)
	return b
}

func (b *VXChipBuilder) Draggable(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":draggable", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) FilterIcon(v string) (r *VXChipBuilder) {
	b.tag.Attr("filter-icon", v)
	return b
}

func (b *VXChipBuilder) Link(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Pill(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":pill", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) PrependAvatar(v string) (r *VXChipBuilder) {
	b.tag.Attr("prepend-avatar", v)
	return b
}

func (b *VXChipBuilder) PrependIcon(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Ripple(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Value(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Text(v string) (r *VXChipBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VXChipBuilder) ModelValue(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Border(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Density(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Elevation(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Disabled(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) SelectedClass(v string) (r *VXChipBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VXChipBuilder) Rounded(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Tile(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Href(v string) (r *VXChipBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VXChipBuilder) Replace(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Exact(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) To(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Size(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Tag(v string) (r *VXChipBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXChipBuilder) Theme(v string) (r *VXChipBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXChipBuilder) Color(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":color", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Variant(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXChipBuilder) Attr(vs ...interface{}) (r *VXChipBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXChipBuilder) Children(children ...h.HTMLComponent) (r *VXChipBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXChipBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXChipBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXChipBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXChipBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXChipBuilder) Class(names ...string) (r *VXChipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXChipBuilder) ClassIf(name string, add bool) (r *VXChipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXChipBuilder) On(name string, value string) (r *VXChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXChipBuilder) Bind(name string, value string) (r *VXChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXChipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
