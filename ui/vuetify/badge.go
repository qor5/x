package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBadgeBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBadge(children ...h.HTMLComponent) (r *VBadgeBuilder) {
	r = &VBadgeBuilder{
		tag: h.Tag("v-badge").Children(children...),
	}
	return
}

func (b *VBadgeBuilder) Bordered(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":bordered", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Color(v string) (r *VBadgeBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBadgeBuilder) Content(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":content", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) Dot(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":dot", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Floating(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Icon(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) Inline(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Label(v string) (r *VBadgeBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VBadgeBuilder) Max(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) ModelValue(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) OffsetX(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":offset-x", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) OffsetY(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":offset-y", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) TextColor(v string) (r *VBadgeBuilder) {
	b.tag.Attr("text-color", v)
	return b
}

func (b *VBadgeBuilder) Location(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) Rounded(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) Tile(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Tag(v string) (r *VBadgeBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBadgeBuilder) Theme(v string) (r *VBadgeBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VBadgeBuilder) Transition(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VBadgeBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBadgeBuilder) Attr(vs ...interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBadgeBuilder) Children(children ...h.HTMLComponent) (r *VBadgeBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBadgeBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBadgeBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBadgeBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBadgeBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBadgeBuilder) Class(names ...string) (r *VBadgeBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBadgeBuilder) ClassIf(name string, add bool) (r *VBadgeBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBadgeBuilder) On(name string, value string) (r *VBadgeBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBadgeBuilder) Bind(name string, value string) (r *VBadgeBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBadgeBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
