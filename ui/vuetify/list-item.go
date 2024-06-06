package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItem(children ...h.HTMLComponent) (r *VListItemBuilder) {
	r = &VListItemBuilder{
		tag: h.Tag("v-list-item").Children(children...),
	}
	return
}

func (b *VListItemBuilder) Title(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":title", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Subtitle(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":subtitle", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Active(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) ActiveClass(v string) (r *VListItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListItemBuilder) ActiveColor(v string) (r *VListItemBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VListItemBuilder) AppendAvatar(v string) (r *VListItemBuilder) {
	b.tag.Attr("append-avatar", v)
	return b
}

func (b *VListItemBuilder) AppendIcon(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) BaseColor(v string) (r *VListItemBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VListItemBuilder) Disabled(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Link(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Nav(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":nav", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) PrependAvatar(v string) (r *VListItemBuilder) {
	b.tag.Attr("prepend-avatar", v)
	return b
}

func (b *VListItemBuilder) PrependIcon(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Ripple(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Value(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Slim(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Border(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Density(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Height(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) MaxHeight(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) MaxWidth(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) MinHeight(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) MinWidth(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Width(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Elevation(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Rounded(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Tile(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Href(v string) (r *VListItemBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VListItemBuilder) Replace(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Exact(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) To(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Tag(v string) (r *VListItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListItemBuilder) Theme(v string) (r *VListItemBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VListItemBuilder) Color(v string) (r *VListItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListItemBuilder) Variant(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Lines(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":lines", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemBuilder) Attr(vs ...interface{}) (r *VListItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemBuilder) Children(children ...h.HTMLComponent) (r *VListItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemBuilder) Class(names ...string) (r *VListItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemBuilder) ClassIf(name string, add bool) (r *VListItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemBuilder) On(name string, value string) (r *VListItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemBuilder) Bind(name string, value string) (r *VListItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
