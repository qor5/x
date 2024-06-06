package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideYTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideYTransition(children ...h.HTMLComponent) (r *VSlideYTransitionBuilder) {
	r = &VSlideYTransitionBuilder{
		tag: h.Tag("v-slide-y-transition").Children(children...),
	}
	return
}

func (b *VSlideYTransitionBuilder) Disabled(v bool) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSlideYTransitionBuilder) Group(v bool) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VSlideYTransitionBuilder) HideOnLeave(v bool) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VSlideYTransitionBuilder) LeaveAbsolute(v bool) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VSlideYTransitionBuilder) Mode(v string) (r *VSlideYTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VSlideYTransitionBuilder) Origin(v string) (r *VSlideYTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VSlideYTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideYTransitionBuilder) Attr(vs ...interface{}) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideYTransitionBuilder) Children(children ...h.HTMLComponent) (r *VSlideYTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideYTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideYTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideYTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideYTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideYTransitionBuilder) Class(names ...string) (r *VSlideYTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideYTransitionBuilder) ClassIf(name string, add bool) (r *VSlideYTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideYTransitionBuilder) On(name string, value string) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideYTransitionBuilder) Bind(name string, value string) (r *VSlideYTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideYTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
