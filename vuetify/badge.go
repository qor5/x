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

func (b *VBadgeBuilder) Bottom(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Color(v string) (r *VBadgeBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBadgeBuilder) Left(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Mode(v string) (r *VBadgeBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VBadgeBuilder) Origin(v string) (r *VBadgeBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VBadgeBuilder) Overlap(v bool) (r *VBadgeBuilder) {
	b.tag.Attr(":overlap", fmt.Sprint(v))
	return b
}

func (b *VBadgeBuilder) Transition(v string) (r *VBadgeBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VBadgeBuilder) Value(v interface{}) (r *VBadgeBuilder) {
	b.tag.Attr(":value", v)
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
