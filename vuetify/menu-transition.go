package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VMenuTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VMenuTransition(children ...h.HTMLComponent) (r *VMenuTransitionBuilder) {
	r = &VMenuTransitionBuilder{
		tag: h.Tag("v-menu-transition").Children(children...),
	}
	return
}

func (b *VMenuTransitionBuilder) Group(v bool) (r *VMenuTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VMenuTransitionBuilder) HideOnLeave(v bool) (r *VMenuTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VMenuTransitionBuilder) LeaveAbsolute(v bool) (r *VMenuTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VMenuTransitionBuilder) Mode(v string) (r *VMenuTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VMenuTransitionBuilder) Origin(v string) (r *VMenuTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VMenuTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VMenuTransitionBuilder) Attr(vs ...interface{}) (r *VMenuTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VMenuTransitionBuilder) Children(children ...h.HTMLComponent) (r *VMenuTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VMenuTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VMenuTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VMenuTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VMenuTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VMenuTransitionBuilder) Class(names ...string) (r *VMenuTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VMenuTransitionBuilder) ClassIf(name string, add bool) (r *VMenuTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VMenuTransitionBuilder) On(name string, value string) (r *VMenuTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VMenuTransitionBuilder) Bind(name string, value string) (r *VMenuTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VMenuTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
