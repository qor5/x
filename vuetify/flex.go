package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFlexBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFlex(children ...h.HTMLComponent) (r *VFlexBuilder) {
	r = &VFlexBuilder{
		tag: h.Tag("v-flex").Children(children...),
	}
	return
}

func (b *VFlexBuilder) AlignSelfBaseline(v bool) (r *VFlexBuilder) {
	b.tag.Attr(":align-self-baseline", fmt.Sprint(v))
	return b
}

func (b *VFlexBuilder) AlignSelfCenter(v bool) (r *VFlexBuilder) {
	b.tag.Attr(":align-self-center", fmt.Sprint(v))
	return b
}

func (b *VFlexBuilder) AlignSelfEnd(v bool) (r *VFlexBuilder) {
	b.tag.Attr(":align-self-end", fmt.Sprint(v))
	return b
}

func (b *VFlexBuilder) AlignSelfStart(v bool) (r *VFlexBuilder) {
	b.tag.Attr(":align-self-start", fmt.Sprint(v))
	return b
}

func (b *VFlexBuilder) Grow(v bool) (r *VFlexBuilder) {
	b.tag.Attr(":grow", fmt.Sprint(v))
	return b
}

func (b *VFlexBuilder) Id(v string) (r *VFlexBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VFlexBuilder) Shrink(v bool) (r *VFlexBuilder) {
	b.tag.Attr(":shrink", fmt.Sprint(v))
	return b
}

func (b *VFlexBuilder) Tag(v string) (r *VFlexBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VFlexBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFlexBuilder) Attr(vs ...interface{}) (r *VFlexBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFlexBuilder) Children(children ...h.HTMLComponent) (r *VFlexBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFlexBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFlexBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFlexBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFlexBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFlexBuilder) Class(names ...string) (r *VFlexBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFlexBuilder) ClassIf(name string, add bool) (r *VFlexBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFlexBuilder) On(name string, value string) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFlexBuilder) Bind(name string, value string) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFlexBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
