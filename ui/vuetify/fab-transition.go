package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFabTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFabTransition(children ...h.HTMLComponent) (r *VFabTransitionBuilder) {
	r = &VFabTransitionBuilder{
		tag: h.Tag("v-fab-transition").Children(children...),
	}
	return
}

func (b *VFabTransitionBuilder) Disabled(v bool) (r *VFabTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VFabTransitionBuilder) Group(v bool) (r *VFabTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VFabTransitionBuilder) HideOnLeave(v bool) (r *VFabTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VFabTransitionBuilder) LeaveAbsolute(v bool) (r *VFabTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VFabTransitionBuilder) Mode(v string) (r *VFabTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VFabTransitionBuilder) Origin(v string) (r *VFabTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VFabTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFabTransitionBuilder) Attr(vs ...interface{}) (r *VFabTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFabTransitionBuilder) Children(children ...h.HTMLComponent) (r *VFabTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFabTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFabTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFabTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFabTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFabTransitionBuilder) Class(names ...string) (r *VFabTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFabTransitionBuilder) ClassIf(name string, add bool) (r *VFabTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFabTransitionBuilder) On(name string, value string) (r *VFabTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFabTransitionBuilder) Bind(name string, value string) (r *VFabTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFabTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
