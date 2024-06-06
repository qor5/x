package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideXReverseTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideXReverseTransition(children ...h.HTMLComponent) (r *VSlideXReverseTransitionBuilder) {
	r = &VSlideXReverseTransitionBuilder{
		tag: h.Tag("v-slide-x-reverse-transition").Children(children...),
	}
	return
}

func (b *VSlideXReverseTransitionBuilder) Disabled(v bool) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSlideXReverseTransitionBuilder) Group(v bool) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VSlideXReverseTransitionBuilder) HideOnLeave(v bool) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VSlideXReverseTransitionBuilder) LeaveAbsolute(v bool) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VSlideXReverseTransitionBuilder) Mode(v string) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VSlideXReverseTransitionBuilder) Origin(v string) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VSlideXReverseTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideXReverseTransitionBuilder) Attr(vs ...interface{}) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideXReverseTransitionBuilder) Children(children ...h.HTMLComponent) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideXReverseTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideXReverseTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideXReverseTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideXReverseTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideXReverseTransitionBuilder) Class(names ...string) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideXReverseTransitionBuilder) ClassIf(name string, add bool) (r *VSlideXReverseTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideXReverseTransitionBuilder) On(name string, value string) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideXReverseTransitionBuilder) Bind(name string, value string) (r *VSlideXReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideXReverseTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
