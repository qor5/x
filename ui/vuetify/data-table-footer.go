package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableFooterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableFooter(children ...h.HTMLComponent) (r *VDataTableFooterBuilder) {
	r = &VDataTableFooterBuilder{
		tag: h.Tag("v-data-table-footer").Children(children...),
	}
	return
}

func (b *VDataTableFooterBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableFooterBuilder) Attr(vs ...interface{}) (r *VDataTableFooterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableFooterBuilder) Children(children ...h.HTMLComponent) (r *VDataTableFooterBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableFooterBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableFooterBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableFooterBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableFooterBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableFooterBuilder) Class(names ...string) (r *VDataTableFooterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableFooterBuilder) ClassIf(name string, add bool) (r *VDataTableFooterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableFooterBuilder) On(name string, value string) (r *VDataTableFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableFooterBuilder) Bind(name string, value string) (r *VDataTableFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableFooterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
