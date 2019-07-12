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

func (b *VFlexBuilder) Col(size SizeType, columns int) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf(":%s%d", size, columns), fmt.Sprint(true))
	return b
}

func (b *VFlexBuilder) Offset(size SizeType, columns int) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf(":offset-%s%d", size, columns), fmt.Sprint(true))
	return b
}

func (b *VFlexBuilder) Order(size SizeType, columns int) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf(":order-%s%d", size, columns), fmt.Sprint(true))

	return b
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

func (b *VFlexBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
