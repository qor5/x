package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCard(children ...h.HTMLComponent) (r *VCardBuilder) {
	r = &VCardBuilder{
		tag: h.Tag("v-card").Children(children...),
	}
	return
}

func (b *VCardBuilder) Title(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":title", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Subtitle(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":subtitle", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Text(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":text", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Image(v string) (r *VCardBuilder) {
	b.tag.Attr("image", v)
	return b
}

func (b *VCardBuilder) Flat(v bool) (r *VCardBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) AppendAvatar(v string) (r *VCardBuilder) {
	b.tag.Attr("append-avatar", v)
	return b
}

func (b *VCardBuilder) AppendIcon(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Disabled(v bool) (r *VCardBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Hover(v bool) (r *VCardBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Link(v bool) (r *VCardBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) PrependAvatar(v string) (r *VCardBuilder) {
	b.tag.Attr("prepend-avatar", v)
	return b
}

func (b *VCardBuilder) PrependIcon(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Ripple(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Border(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Density(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Height(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VCardBuilder) MaxHeight(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VCardBuilder) MaxWidth(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VCardBuilder) MinHeight(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VCardBuilder) MinWidth(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Width(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Elevation(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Loading(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Location(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Position(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Rounded(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Tile(v bool) (r *VCardBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Href(v string) (r *VCardBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VCardBuilder) Replace(v bool) (r *VCardBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Exact(v bool) (r *VCardBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) To(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VCardBuilder) Tag(v string) (r *VCardBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCardBuilder) Theme(v string) (r *VCardBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VCardBuilder) Color(v string) (r *VCardBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCardBuilder) Variant(v interface{}) (r *VCardBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VCardBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCardBuilder) Attr(vs ...interface{}) (r *VCardBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCardBuilder) Children(children ...h.HTMLComponent) (r *VCardBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCardBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCardBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCardBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCardBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCardBuilder) Class(names ...string) (r *VCardBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardBuilder) ClassIf(name string, add bool) (r *VCardBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardBuilder) On(name string, value string) (r *VCardBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardBuilder) Bind(name string, value string) (r *VCardBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
