package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideXTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideXTransition(children ...h.HTMLComponent) (r *VSlideXTransitionBuilder) {
	r = &VSlideXTransitionBuilder{
		tag: h.Tag("v-slide-x-transition").Children(children...),
	}
	return
}

func (b *VSlideXTransitionBuilder) Disabled(v bool) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSlideXTransitionBuilder) Group(v bool) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VSlideXTransitionBuilder) HideOnLeave(v bool) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VSlideXTransitionBuilder) LeaveAbsolute(v bool) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VSlideXTransitionBuilder) Mode(v string) (r *VSlideXTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VSlideXTransitionBuilder) Origin(v string) (r *VSlideXTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VSlideXTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideXTransitionBuilder) Attr(vs ...interface{}) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideXTransitionBuilder) Children(children ...h.HTMLComponent) (r *VSlideXTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideXTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideXTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideXTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideXTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideXTransitionBuilder) Class(names ...string) (r *VSlideXTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideXTransitionBuilder) ClassIf(name string, add bool) (r *VSlideXTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideXTransitionBuilder) On(name string, value string) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideXTransitionBuilder) Bind(name string, value string) (r *VSlideXTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideXTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
