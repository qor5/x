package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VChipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VChip(children ...h.HTMLComponent) (r *VChipBuilder) {
	r = &VChipBuilder{
		tag: h.Tag("v-chip").Children(children...),
	}
	return
}

func (b *VChipBuilder) Label(v bool) (r *VChipBuilder) {
	b.tag.Attr(":label", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Filter(v bool) (r *VChipBuilder) {
	b.tag.Attr(":filter", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) ActiveClass(v string) (r *VChipBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VChipBuilder) AppendAvatar(v string) (r *VChipBuilder) {
	b.tag.Attr("append-avatar", v)
	return b
}

func (b *VChipBuilder) AppendIcon(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Closable(v bool) (r *VChipBuilder) {
	b.tag.Attr(":closable", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) CloseIcon(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":close-icon", h.JSONString(v))
	return b
}

func (b *VChipBuilder) CloseLabel(v string) (r *VChipBuilder) {
	b.tag.Attr("close-label", v)
	return b
}

func (b *VChipBuilder) Draggable(v bool) (r *VChipBuilder) {
	b.tag.Attr(":draggable", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) FilterIcon(v string) (r *VChipBuilder) {
	b.tag.Attr("filter-icon", v)
	return b
}

func (b *VChipBuilder) Link(v bool) (r *VChipBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Pill(v bool) (r *VChipBuilder) {
	b.tag.Attr(":pill", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) PrependAvatar(v string) (r *VChipBuilder) {
	b.tag.Attr("prepend-avatar", v)
	return b
}

func (b *VChipBuilder) PrependIcon(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Ripple(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Value(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Text(v string) (r *VChipBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VChipBuilder) ModelValue(v bool) (r *VChipBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Border(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Density(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Elevation(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Disabled(v bool) (r *VChipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) SelectedClass(v string) (r *VChipBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VChipBuilder) Rounded(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Tile(v bool) (r *VChipBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Href(v string) (r *VChipBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VChipBuilder) Replace(v bool) (r *VChipBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Exact(v bool) (r *VChipBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) To(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Size(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Tag(v string) (r *VChipBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VChipBuilder) Theme(v string) (r *VChipBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VChipBuilder) Color(v string) (r *VChipBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VChipBuilder) Variant(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VChipBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VChipBuilder) Attr(vs ...interface{}) (r *VChipBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VChipBuilder) Children(children ...h.HTMLComponent) (r *VChipBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VChipBuilder) AppendChildren(children ...h.HTMLComponent) (r *VChipBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VChipBuilder) PrependChildren(children ...h.HTMLComponent) (r *VChipBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VChipBuilder) Class(names ...string) (r *VChipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VChipBuilder) ClassIf(name string, add bool) (r *VChipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VChipBuilder) On(name string, value string) (r *VChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VChipBuilder) Bind(name string, value string) (r *VChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VChipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
