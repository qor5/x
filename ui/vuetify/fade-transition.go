package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFadeTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFadeTransition(children ...h.HTMLComponent) (r *VFadeTransitionBuilder) {
	r = &VFadeTransitionBuilder{
		tag: h.Tag("v-fade-transition").Children(children...),
	}
	return
}

func (b *VFadeTransitionBuilder) Disabled(v bool) (r *VFadeTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VFadeTransitionBuilder) Group(v bool) (r *VFadeTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VFadeTransitionBuilder) HideOnLeave(v bool) (r *VFadeTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VFadeTransitionBuilder) LeaveAbsolute(v bool) (r *VFadeTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VFadeTransitionBuilder) Mode(v string) (r *VFadeTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VFadeTransitionBuilder) Origin(v string) (r *VFadeTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VFadeTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFadeTransitionBuilder) Attr(vs ...interface{}) (r *VFadeTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFadeTransitionBuilder) Children(children ...h.HTMLComponent) (r *VFadeTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFadeTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFadeTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFadeTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFadeTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFadeTransitionBuilder) Class(names ...string) (r *VFadeTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFadeTransitionBuilder) ClassIf(name string, add bool) (r *VFadeTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFadeTransitionBuilder) On(name string, value string) (r *VFadeTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFadeTransitionBuilder) Bind(name string, value string) (r *VFadeTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFadeTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
