package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideYReverseTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideYReverseTransition(children ...h.HTMLComponent) (r *VSlideYReverseTransitionBuilder) {
	r = &VSlideYReverseTransitionBuilder{
		tag: h.Tag("v-slide-y-reverse-transition").Children(children...),
	}
	return
}

func (b *VSlideYReverseTransitionBuilder) Disabled(v bool) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSlideYReverseTransitionBuilder) Group(v bool) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VSlideYReverseTransitionBuilder) HideOnLeave(v bool) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VSlideYReverseTransitionBuilder) LeaveAbsolute(v bool) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VSlideYReverseTransitionBuilder) Mode(v string) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VSlideYReverseTransitionBuilder) Origin(v string) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VSlideYReverseTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideYReverseTransitionBuilder) Attr(vs ...interface{}) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideYReverseTransitionBuilder) Children(children ...h.HTMLComponent) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideYReverseTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideYReverseTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideYReverseTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideYReverseTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideYReverseTransitionBuilder) Class(names ...string) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideYReverseTransitionBuilder) ClassIf(name string, add bool) (r *VSlideYReverseTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideYReverseTransitionBuilder) On(name string, value string) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideYReverseTransitionBuilder) Bind(name string, value string) (r *VSlideYReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideYReverseTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
