package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBannerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBanner(children ...h.HTMLComponent) (r *VBannerBuilder) {
	r = &VBannerBuilder{
		tag: h.Tag("v-banner").Children(children...),
	}
	return
}

func (b *VBannerBuilder) Text(v string) (r *VBannerBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VBannerBuilder) Avatar(v string) (r *VBannerBuilder) {
	b.tag.Attr("avatar", v)
	return b
}

func (b *VBannerBuilder) BgColor(v string) (r *VBannerBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VBannerBuilder) Color(v string) (r *VBannerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBannerBuilder) Icon(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Stacked(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Sticky(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Border(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Density(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Height(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) MaxHeight(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) MaxWidth(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) MinHeight(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) MinWidth(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Width(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Mobile(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) MobileBreakpoint(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Elevation(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Location(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Position(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Rounded(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) Tile(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Tag(v string) (r *VBannerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBannerBuilder) Theme(v string) (r *VBannerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VBannerBuilder) Lines(v interface{}) (r *VBannerBuilder) {
	b.tag.Attr(":lines", h.JSONString(v))
	return b
}

func (b *VBannerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBannerBuilder) Attr(vs ...interface{}) (r *VBannerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBannerBuilder) Children(children ...h.HTMLComponent) (r *VBannerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBannerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBannerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBannerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBannerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBannerBuilder) Class(names ...string) (r *VBannerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBannerBuilder) ClassIf(name string, add bool) (r *VBannerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBannerBuilder) On(name string, value string) (r *VBannerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBannerBuilder) Bind(name string, value string) (r *VBannerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBannerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
