package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardMediaBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardMedia() (r *VCardMediaBuilder) {
	r = &VCardMediaBuilder{
		tag: h.Tag("v-card-media"),
	}
	return
}

func (b *VCardMediaBuilder) Alt(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("alt", v)
	return b
}

func (b *VCardMediaBuilder) AspectRatio(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("aspect-ratio", v)
	return b
}

func (b *VCardMediaBuilder) Contain(v bool) (r *VCardMediaBuilder) {
	b.tag.Attr(":contain", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) Gradient(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("gradient", v)
	return b
}

func (b *VCardMediaBuilder) Height(v int) (r *VCardMediaBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) LazySrc(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("lazy-src", v)
	return b
}

func (b *VCardMediaBuilder) MaxHeight(v int) (r *VCardMediaBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) MaxWidth(v int) (r *VCardMediaBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) MinHeight(v int) (r *VCardMediaBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) MinWidth(v int) (r *VCardMediaBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) Position(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("position", v)
	return b
}

func (b *VCardMediaBuilder) Sizes(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("sizes", v)
	return b
}

func (b *VCardMediaBuilder) Src(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("src", v)
	return b
}

func (b *VCardMediaBuilder) Srcset(v string) (r *VCardMediaBuilder) {
	b.tag.Attr("srcset", v)
	return b
}

func (b *VCardMediaBuilder) Transition(v bool) (r *VCardMediaBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) Width(v int) (r *VCardMediaBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VCardMediaBuilder) Class(names ...string) (r *VCardMediaBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardMediaBuilder) ClassIf(name string, add bool) (r *VCardMediaBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardMediaBuilder) On(name string, value string) (r *VCardMediaBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardMediaBuilder) Bind(name string, value string) (r *VCardMediaBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardMediaBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
