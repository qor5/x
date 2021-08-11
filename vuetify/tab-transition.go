package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabTransition(children ...h.HTMLComponent) (r *VTabTransitionBuilder) {
	r = &VTabTransitionBuilder{
		tag: h.Tag("v-tab-transition").Children(children...),
	}
	return
}

func (b *VTabTransitionBuilder) Group(v bool) (r *VTabTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VTabTransitionBuilder) HideOnLeave(v bool) (r *VTabTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VTabTransitionBuilder) LeaveAbsolute(v bool) (r *VTabTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VTabTransitionBuilder) Mode(v string) (r *VTabTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VTabTransitionBuilder) Origin(v string) (r *VTabTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VTabTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabTransitionBuilder) Attr(vs ...interface{}) (r *VTabTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabTransitionBuilder) Children(children ...h.HTMLComponent) (r *VTabTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabTransitionBuilder) Class(names ...string) (r *VTabTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabTransitionBuilder) ClassIf(name string, add bool) (r *VTabTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabTransitionBuilder) On(name string, value string) (r *VTabTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabTransitionBuilder) Bind(name string, value string) (r *VTabTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
