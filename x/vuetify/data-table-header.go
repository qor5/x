package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableHeaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableHeader() (r *VDataTableHeaderBuilder) {
	r = &VDataTableHeaderBuilder{
		tag: h.Tag("v-data-table-header"),
	}
	return
}

func (b *VDataTableHeaderBuilder) Mobile(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableHeaderBuilder) Attr(vs ...interface{}) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableHeaderBuilder) Children(children ...h.HTMLComponent) (r *VDataTableHeaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableHeaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableHeaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableHeaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableHeaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableHeaderBuilder) Class(names ...string) (r *VDataTableHeaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableHeaderBuilder) ClassIf(name string, add bool) (r *VDataTableHeaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableHeaderBuilder) On(name string, value string) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableHeaderBuilder) Bind(name string, value string) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableHeaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
