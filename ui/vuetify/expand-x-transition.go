package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpandXTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpandXTransition(children ...h.HTMLComponent) (r *VExpandXTransitionBuilder) {
	r = &VExpandXTransitionBuilder{
		tag: h.Tag("v-expand-x-transition").Children(children...),
	}
	return
}

func (b *VExpandXTransitionBuilder) Disabled(v bool) (r *VExpandXTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpandXTransitionBuilder) Group(v bool) (r *VExpandXTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VExpandXTransitionBuilder) Mode(v interface{}) (r *VExpandXTransitionBuilder) {
	b.tag.Attr(":mode", h.JSONString(v))
	return b
}

func (b *VExpandXTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpandXTransitionBuilder) Attr(vs ...interface{}) (r *VExpandXTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpandXTransitionBuilder) Children(children ...h.HTMLComponent) (r *VExpandXTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpandXTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpandXTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpandXTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpandXTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpandXTransitionBuilder) Class(names ...string) (r *VExpandXTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpandXTransitionBuilder) ClassIf(name string, add bool) (r *VExpandXTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpandXTransitionBuilder) On(name string, value string) (r *VExpandXTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpandXTransitionBuilder) Bind(name string, value string) (r *VExpandXTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpandXTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
