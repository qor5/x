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

func (b *VBannerBuilder) App(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Color(v string) (r *VBannerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBannerBuilder) Dark(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Elevation(v int) (r *VBannerBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Height(v int) (r *VBannerBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Icon(v string) (r *VBannerBuilder) {
	b.tag.Attr("icon", v)
	return b
}

func (b *VBannerBuilder) IconColor(v string) (r *VBannerBuilder) {
	b.tag.Attr("icon-color", v)
	return b
}

func (b *VBannerBuilder) Light(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) MaxHeight(v int) (r *VBannerBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) MaxWidth(v int) (r *VBannerBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) MinHeight(v int) (r *VBannerBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) MinWidth(v int) (r *VBannerBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) MobileBreakpoint(v int) (r *VBannerBuilder) {
	b.tag.Attr(":mobile-breakpoint", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Outlined(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Rounded(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Shaped(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) SingleLine(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Sticky(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Tag(v string) (r *VBannerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBannerBuilder) Tile(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Value(v bool) (r *VBannerBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VBannerBuilder) Width(v int) (r *VBannerBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
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
