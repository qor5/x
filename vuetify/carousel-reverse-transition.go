package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCarouselReverseTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCarouselReverseTransition(children ...h.HTMLComponent) (r *VCarouselReverseTransitionBuilder) {
	r = &VCarouselReverseTransitionBuilder{
		tag: h.Tag("v-carousel-reverse-transition").Children(children...),
	}
	return
}

func (b *VCarouselReverseTransitionBuilder) Group(v bool) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VCarouselReverseTransitionBuilder) HideOnLeave(v bool) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VCarouselReverseTransitionBuilder) LeaveAbsolute(v bool) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VCarouselReverseTransitionBuilder) Mode(v string) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VCarouselReverseTransitionBuilder) Origin(v string) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VCarouselReverseTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCarouselReverseTransitionBuilder) Attr(vs ...interface{}) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCarouselReverseTransitionBuilder) Children(children ...h.HTMLComponent) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCarouselReverseTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCarouselReverseTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCarouselReverseTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCarouselReverseTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCarouselReverseTransitionBuilder) Class(names ...string) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCarouselReverseTransitionBuilder) ClassIf(name string, add bool) (r *VCarouselReverseTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCarouselReverseTransitionBuilder) On(name string, value string) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCarouselReverseTransitionBuilder) Bind(name string, value string) (r *VCarouselReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCarouselReverseTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
