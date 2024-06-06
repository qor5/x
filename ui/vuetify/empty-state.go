package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VEmptyStateBuilder struct {
	tag *h.HTMLTagBuilder
}

func VEmptyState(children ...h.HTMLComponent) (r *VEmptyStateBuilder) {
	r = &VEmptyStateBuilder{
		tag: h.Tag("v-empty-state").Children(children...),
	}
	return
}

func (b *VEmptyStateBuilder) Headline(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("headline", v)
	return b
}

func (b *VEmptyStateBuilder) Title(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VEmptyStateBuilder) Text(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VEmptyStateBuilder) ActionText(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("action-text", v)
	return b
}

func (b *VEmptyStateBuilder) BgColor(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VEmptyStateBuilder) Color(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VEmptyStateBuilder) Icon(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) Image(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("image", v)
	return b
}

func (b *VEmptyStateBuilder) Justify(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":justify", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) TextWidth(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":text-width", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) Href(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VEmptyStateBuilder) To(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("to", v)
	return b
}

func (b *VEmptyStateBuilder) Height(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) MaxHeight(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) MaxWidth(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) MinHeight(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) MinWidth(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) Width(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) Size(v interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VEmptyStateBuilder) Theme(v string) (r *VEmptyStateBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VEmptyStateBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VEmptyStateBuilder) Attr(vs ...interface{}) (r *VEmptyStateBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VEmptyStateBuilder) Children(children ...h.HTMLComponent) (r *VEmptyStateBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VEmptyStateBuilder) AppendChildren(children ...h.HTMLComponent) (r *VEmptyStateBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VEmptyStateBuilder) PrependChildren(children ...h.HTMLComponent) (r *VEmptyStateBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VEmptyStateBuilder) Class(names ...string) (r *VEmptyStateBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VEmptyStateBuilder) ClassIf(name string, add bool) (r *VEmptyStateBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VEmptyStateBuilder) On(name string, value string) (r *VEmptyStateBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VEmptyStateBuilder) Bind(name string, value string) (r *VEmptyStateBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VEmptyStateBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
