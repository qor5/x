package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCarouselTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCarouselTransition(children ...h.HTMLComponent) (r *VCarouselTransitionBuilder) {
	r = &VCarouselTransitionBuilder{
		tag: h.Tag("v-carousel-transition").Children(children...),
	}
	return
}

func (b *VCarouselTransitionBuilder) Group(v bool) (r *VCarouselTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VCarouselTransitionBuilder) HideOnLeave(v bool) (r *VCarouselTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VCarouselTransitionBuilder) LeaveAbsolute(v bool) (r *VCarouselTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VCarouselTransitionBuilder) Mode(v string) (r *VCarouselTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VCarouselTransitionBuilder) Origin(v string) (r *VCarouselTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VCarouselTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCarouselTransitionBuilder) Attr(vs ...interface{}) (r *VCarouselTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCarouselTransitionBuilder) Children(children ...h.HTMLComponent) (r *VCarouselTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCarouselTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCarouselTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCarouselTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCarouselTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCarouselTransitionBuilder) Class(names ...string) (r *VCarouselTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCarouselTransitionBuilder) ClassIf(name string, add bool) (r *VCarouselTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCarouselTransitionBuilder) On(name string, value string) (r *VCarouselTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCarouselTransitionBuilder) Bind(name string, value string) (r *VCarouselTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCarouselTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
