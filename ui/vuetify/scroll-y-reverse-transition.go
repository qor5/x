package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VScrollYReverseTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VScrollYReverseTransition(children ...h.HTMLComponent) (r *VScrollYReverseTransitionBuilder) {
	r = &VScrollYReverseTransitionBuilder{
		tag: h.Tag("v-scroll-y-reverse-transition").Children(children...),
	}
	return
}

func (b *VScrollYReverseTransitionBuilder) Disabled(v bool) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VScrollYReverseTransitionBuilder) Group(v bool) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VScrollYReverseTransitionBuilder) HideOnLeave(v bool) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VScrollYReverseTransitionBuilder) LeaveAbsolute(v bool) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VScrollYReverseTransitionBuilder) Mode(v string) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VScrollYReverseTransitionBuilder) Origin(v string) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VScrollYReverseTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VScrollYReverseTransitionBuilder) Attr(vs ...interface{}) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VScrollYReverseTransitionBuilder) Children(children ...h.HTMLComponent) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VScrollYReverseTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VScrollYReverseTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VScrollYReverseTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VScrollYReverseTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VScrollYReverseTransitionBuilder) Class(names ...string) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VScrollYReverseTransitionBuilder) ClassIf(name string, add bool) (r *VScrollYReverseTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VScrollYReverseTransitionBuilder) On(name string, value string) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VScrollYReverseTransitionBuilder) Bind(name string, value string) (r *VScrollYReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VScrollYReverseTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
