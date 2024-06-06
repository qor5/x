package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VScrollXTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VScrollXTransition(children ...h.HTMLComponent) (r *VScrollXTransitionBuilder) {
	r = &VScrollXTransitionBuilder{
		tag: h.Tag("v-scroll-x-transition").Children(children...),
	}
	return
}

func (b *VScrollXTransitionBuilder) Disabled(v bool) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VScrollXTransitionBuilder) Group(v bool) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VScrollXTransitionBuilder) HideOnLeave(v bool) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VScrollXTransitionBuilder) LeaveAbsolute(v bool) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VScrollXTransitionBuilder) Mode(v string) (r *VScrollXTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VScrollXTransitionBuilder) Origin(v string) (r *VScrollXTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VScrollXTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VScrollXTransitionBuilder) Attr(vs ...interface{}) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VScrollXTransitionBuilder) Children(children ...h.HTMLComponent) (r *VScrollXTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VScrollXTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VScrollXTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VScrollXTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VScrollXTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VScrollXTransitionBuilder) Class(names ...string) (r *VScrollXTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VScrollXTransitionBuilder) ClassIf(name string, add bool) (r *VScrollXTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VScrollXTransitionBuilder) On(name string, value string) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VScrollXTransitionBuilder) Bind(name string, value string) (r *VScrollXTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VScrollXTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
