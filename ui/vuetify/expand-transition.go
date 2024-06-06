package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpandTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpandTransition(children ...h.HTMLComponent) (r *VExpandTransitionBuilder) {
	r = &VExpandTransitionBuilder{
		tag: h.Tag("v-expand-transition").Children(children...),
	}
	return
}

func (b *VExpandTransitionBuilder) Disabled(v bool) (r *VExpandTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpandTransitionBuilder) Group(v bool) (r *VExpandTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VExpandTransitionBuilder) Mode(v interface{}) (r *VExpandTransitionBuilder) {
	b.tag.Attr(":mode", h.JSONString(v))
	return b
}

func (b *VExpandTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpandTransitionBuilder) Attr(vs ...interface{}) (r *VExpandTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpandTransitionBuilder) Children(children ...h.HTMLComponent) (r *VExpandTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpandTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpandTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpandTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpandTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpandTransitionBuilder) Class(names ...string) (r *VExpandTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpandTransitionBuilder) ClassIf(name string, add bool) (r *VExpandTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpandTransitionBuilder) On(name string, value string) (r *VExpandTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpandTransitionBuilder) Bind(name string, value string) (r *VExpandTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpandTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
