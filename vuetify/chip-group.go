package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VChipGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VChipGroup(children ...h.HTMLComponent) (r *VChipGroupBuilder) {
	r = &VChipGroupBuilder{
		tag: h.Tag("v-chip-group").Children(children...),
	}
	return
}

func (b *VChipGroupBuilder) ActiveClass(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VChipGroupBuilder) CenterActive(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":center-active", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Color(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VChipGroupBuilder) Column(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":column", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Dark(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Light(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Mandatory(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Max(v int) (r *VChipGroupBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) MobileBreakpoint(v int) (r *VChipGroupBuilder) {
	b.tag.Attr(":mobile-breakpoint", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Multiple(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) NextIcon(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VChipGroupBuilder) PrevIcon(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VChipGroupBuilder) ShowArrows(v bool) (r *VChipGroupBuilder) {
	b.tag.Attr(":show-arrows", fmt.Sprint(v))
	return b
}

func (b *VChipGroupBuilder) Tag(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VChipGroupBuilder) Value(v interface{}) (r *VChipGroupBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VChipGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VChipGroupBuilder) Attr(vs ...interface{}) (r *VChipGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VChipGroupBuilder) Children(children ...h.HTMLComponent) (r *VChipGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VChipGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VChipGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VChipGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VChipGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VChipGroupBuilder) Class(names ...string) (r *VChipGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VChipGroupBuilder) ClassIf(name string, add bool) (r *VChipGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VChipGroupBuilder) On(name string, value string) (r *VChipGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VChipGroupBuilder) Bind(name string, value string) (r *VChipGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VChipGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
