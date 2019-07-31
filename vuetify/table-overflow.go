package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTableOverflowBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTableOverflow() (r *VTableOverflowBuilder) {
	r = &VTableOverflowBuilder{
		tag: h.Tag("v-table-overflow"),
	}
	return
}

func (b *VTableOverflowBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTableOverflowBuilder) Attr(vs ...interface{}) (r *VTableOverflowBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTableOverflowBuilder) Children(children ...h.HTMLComponent) (r *VTableOverflowBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTableOverflowBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTableOverflowBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTableOverflowBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTableOverflowBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTableOverflowBuilder) Class(names ...string) (r *VTableOverflowBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTableOverflowBuilder) ClassIf(name string, add bool) (r *VTableOverflowBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTableOverflowBuilder) On(name string, value string) (r *VTableOverflowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTableOverflowBuilder) Bind(name string, value string) (r *VTableOverflowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTableOverflowBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
