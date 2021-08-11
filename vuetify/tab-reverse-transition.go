package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabReverseTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabReverseTransition(children ...h.HTMLComponent) (r *VTabReverseTransitionBuilder) {
	r = &VTabReverseTransitionBuilder{
		tag: h.Tag("v-tab-reverse-transition").Children(children...),
	}
	return
}

func (b *VTabReverseTransitionBuilder) Group(v bool) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VTabReverseTransitionBuilder) HideOnLeave(v bool) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VTabReverseTransitionBuilder) LeaveAbsolute(v bool) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VTabReverseTransitionBuilder) Mode(v string) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VTabReverseTransitionBuilder) Origin(v string) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VTabReverseTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabReverseTransitionBuilder) Attr(vs ...interface{}) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabReverseTransitionBuilder) Children(children ...h.HTMLComponent) (r *VTabReverseTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabReverseTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabReverseTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabReverseTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabReverseTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabReverseTransitionBuilder) Class(names ...string) (r *VTabReverseTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabReverseTransitionBuilder) ClassIf(name string, add bool) (r *VTabReverseTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabReverseTransitionBuilder) On(name string, value string) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabReverseTransitionBuilder) Bind(name string, value string) (r *VTabReverseTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabReverseTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
