package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VScrollYTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VScrollYTransition(children ...h.HTMLComponent) (r *VScrollYTransitionBuilder) {
	r = &VScrollYTransitionBuilder{
		tag: h.Tag("v-scroll-y-transition").Children(children...),
	}
	return
}

func (b *VScrollYTransitionBuilder) Disabled(v bool) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VScrollYTransitionBuilder) Group(v bool) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VScrollYTransitionBuilder) HideOnLeave(v bool) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VScrollYTransitionBuilder) LeaveAbsolute(v bool) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VScrollYTransitionBuilder) Mode(v string) (r *VScrollYTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VScrollYTransitionBuilder) Origin(v string) (r *VScrollYTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VScrollYTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VScrollYTransitionBuilder) Attr(vs ...interface{}) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VScrollYTransitionBuilder) Children(children ...h.HTMLComponent) (r *VScrollYTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VScrollYTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VScrollYTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VScrollYTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VScrollYTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VScrollYTransitionBuilder) Class(names ...string) (r *VScrollYTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VScrollYTransitionBuilder) ClassIf(name string, add bool) (r *VScrollYTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VScrollYTransitionBuilder) On(name string, value string) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VScrollYTransitionBuilder) Bind(name string, value string) (r *VScrollYTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VScrollYTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
