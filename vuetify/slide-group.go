package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideGroup(children ...h.HTMLComponent) (r *VSlideGroupBuilder) {
	r = &VSlideGroupBuilder{
		tag: h.Tag("v-slide-group").Children(children...),
	}
	return
}

func (b *VSlideGroupBuilder) ActiveClass(v string) (r *VSlideGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VSlideGroupBuilder) CenterActive(v bool) (r *VSlideGroupBuilder) {
	b.tag.Attr(":center-active", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) Dark(v bool) (r *VSlideGroupBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) Light(v bool) (r *VSlideGroupBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) Mandatory(v bool) (r *VSlideGroupBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) Max(v int) (r *VSlideGroupBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) MobileBreakpoint(v int) (r *VSlideGroupBuilder) {
	b.tag.Attr(":mobile-breakpoint", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) Multiple(v bool) (r *VSlideGroupBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) NextIcon(v string) (r *VSlideGroupBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VSlideGroupBuilder) PrevIcon(v string) (r *VSlideGroupBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VSlideGroupBuilder) ShowArrows(v bool) (r *VSlideGroupBuilder) {
	b.tag.Attr(":show-arrows", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupBuilder) Tag(v string) (r *VSlideGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VSlideGroupBuilder) Value(v interface{}) (r *VSlideGroupBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSlideGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideGroupBuilder) Attr(vs ...interface{}) (r *VSlideGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideGroupBuilder) Children(children ...h.HTMLComponent) (r *VSlideGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideGroupBuilder) Class(names ...string) (r *VSlideGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideGroupBuilder) ClassIf(name string, add bool) (r *VSlideGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideGroupBuilder) On(name string, value string) (r *VSlideGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideGroupBuilder) Bind(name string, value string) (r *VSlideGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
