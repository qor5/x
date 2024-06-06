package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VScaleTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VScaleTransition(children ...h.HTMLComponent) (r *VScaleTransitionBuilder) {
	r = &VScaleTransitionBuilder{
		tag: h.Tag("v-scale-transition").Children(children...),
	}
	return
}

func (b *VScaleTransitionBuilder) Disabled(v bool) (r *VScaleTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VScaleTransitionBuilder) Group(v bool) (r *VScaleTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VScaleTransitionBuilder) HideOnLeave(v bool) (r *VScaleTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VScaleTransitionBuilder) LeaveAbsolute(v bool) (r *VScaleTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VScaleTransitionBuilder) Mode(v string) (r *VScaleTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VScaleTransitionBuilder) Origin(v string) (r *VScaleTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VScaleTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VScaleTransitionBuilder) Attr(vs ...interface{}) (r *VScaleTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VScaleTransitionBuilder) Children(children ...h.HTMLComponent) (r *VScaleTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VScaleTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VScaleTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VScaleTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VScaleTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VScaleTransitionBuilder) Class(names ...string) (r *VScaleTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VScaleTransitionBuilder) ClassIf(name string, add bool) (r *VScaleTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VScaleTransitionBuilder) On(name string, value string) (r *VScaleTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VScaleTransitionBuilder) Bind(name string, value string) (r *VScaleTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VScaleTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
