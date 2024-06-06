package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VScrollXReverseTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VScrollXReverseTransition(children ...h.HTMLComponent) (r *VScrollXReverseTransitionBuilder) {
	r = &VScrollXReverseTransitionBuilder{
		tag: h.Tag("v-scroll-x-reverse-transition").Children(children...),
	}
	return
}

func (b *VScrollXReverseTransitionBuilder) Disabled(v bool) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VScrollXReverseTransitionBuilder) Group(v bool) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VScrollXReverseTransitionBuilder) HideOnLeave(v bool) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VScrollXReverseTransitionBuilder) LeaveAbsolute(v bool) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VScrollXReverseTransitionBuilder) Mode(v string) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VScrollXReverseTransitionBuilder) Origin(v string) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VScrollXReverseTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VScrollXReverseTransitionBuilder) Attr(vs ...interface{}) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VScrollXReverseTransitionBuilder) Children(children ...h.HTMLComponent) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VScrollXReverseTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VScrollXReverseTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VScrollXReverseTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VScrollXReverseTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VScrollXReverseTransitionBuilder) Class(names ...string) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VScrollXReverseTransitionBuilder) ClassIf(name string, add bool) (r *VScrollXReverseTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VScrollXReverseTransitionBuilder) On(name string, value string) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VScrollXReverseTransitionBuilder) Bind(name string, value string) (r *VScrollXReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VScrollXReverseTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
