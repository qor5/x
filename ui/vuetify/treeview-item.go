package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTreeviewItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTreeviewItem(children ...h.HTMLComponent) (r *VTreeviewItemBuilder) {
	r = &VTreeviewItemBuilder{
		tag: h.Tag("v-treeview-item").Children(children...),
	}
	return
}

func (b *VTreeviewItemBuilder) Title(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":title", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Subtitle(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":subtitle", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Loading(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) ToggleIcon(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":toggle-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Active(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) ActiveClass(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTreeviewItemBuilder) ActiveColor(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VTreeviewItemBuilder) AppendAvatar(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("append-avatar", v)
	return b
}

func (b *VTreeviewItemBuilder) AppendIcon(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) BaseColor(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VTreeviewItemBuilder) Disabled(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) Link(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) Nav(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":nav", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) PrependAvatar(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("prepend-avatar", v)
	return b
}

func (b *VTreeviewItemBuilder) PrependIcon(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Ripple(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Value(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Slim(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) Border(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Density(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Height(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) MaxHeight(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) MaxWidth(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) MinHeight(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) MinWidth(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Width(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Elevation(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Rounded(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Tile(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) Href(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VTreeviewItemBuilder) Replace(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) Exact(v bool) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VTreeviewItemBuilder) To(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Tag(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTreeviewItemBuilder) Theme(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTreeviewItemBuilder) Color(v string) (r *VTreeviewItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTreeviewItemBuilder) Variant(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) Lines(v interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(":lines", h.JSONString(v))
	return b
}

func (b *VTreeviewItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTreeviewItemBuilder) Attr(vs ...interface{}) (r *VTreeviewItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTreeviewItemBuilder) Children(children ...h.HTMLComponent) (r *VTreeviewItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTreeviewItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTreeviewItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTreeviewItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTreeviewItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTreeviewItemBuilder) Class(names ...string) (r *VTreeviewItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTreeviewItemBuilder) ClassIf(name string, add bool) (r *VTreeviewItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTreeviewItemBuilder) On(name string, value string) (r *VTreeviewItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTreeviewItemBuilder) Bind(name string, value string) (r *VTreeviewItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTreeviewItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
