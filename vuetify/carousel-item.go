package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCarouselItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCarouselItem(children ...h.HTMLComponent) (r *VCarouselItemBuilder) {
	r = &VCarouselItemBuilder{
		tag: h.Tag("v-carousel-item").Children(children...),
	}
	return
}

func (b *VCarouselItemBuilder) ActiveClass(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VCarouselItemBuilder) Disabled(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Lazy(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) ReverseTransition(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":reverse-transition", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Transition(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Value(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VCarouselItemBuilder) Class(names ...string) (r *VCarouselItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCarouselItemBuilder) ClassIf(name string, add bool) (r *VCarouselItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCarouselItemBuilder) On(name string, value string) (r *VCarouselItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCarouselItemBuilder) Bind(name string, value string) (r *VCarouselItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCarouselItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
