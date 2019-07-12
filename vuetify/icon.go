package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VIcon(name string) (r *VIconBuilder) {
	r = &VIconBuilder{
		tag: h.Tag("v-icon").Name(name),
	}
	return
}
func (b *VIconBuilder) Color(v string) (r *VIconBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VIconBuilder) Dark(v bool) (r *VIconBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Disabled(v bool) (r *VIconBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Large(v bool) (r *VIconBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Left(v bool) (r *VIconBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Light(v bool) (r *VIconBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Medium(v bool) (r *VIconBuilder) {
	b.tag.Attr(":medium", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Right(v bool) (r *VIconBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Size(v int) (r *VIconBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Small(v bool) (r *VIconBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) XLarge(v bool) (r *VIconBuilder) {
	b.tag.Attr(":x-large", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
