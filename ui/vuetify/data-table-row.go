package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableRowBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableRow(children ...h.HTMLComponent) (r *VDataTableRowBuilder) {
	r = &VDataTableRowBuilder{
		tag: h.Tag("v-data-table-row").Children(children...),
	}
	return
}

func (b *VDataTableRowBuilder) CellProps(v interface{}) (r *VDataTableRowBuilder) {
	b.tag.Attr(":cell-props", h.JSONString(v))
	return b
}

func (b *VDataTableRowBuilder) Mobile(v bool) (r *VDataTableRowBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableRowBuilder) Index(v int) (r *VDataTableRowBuilder) {
	b.tag.Attr(":index", fmt.Sprint(v))
	return b
}

func (b *VDataTableRowBuilder) MobileBreakpoint(v interface{}) (r *VDataTableRowBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VDataTableRowBuilder) Item(v interface{}) (r *VDataTableRowBuilder) {
	b.tag.Attr(":item", h.JSONString(v))
	return b
}

func (b *VDataTableRowBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableRowBuilder) Attr(vs ...interface{}) (r *VDataTableRowBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableRowBuilder) Children(children ...h.HTMLComponent) (r *VDataTableRowBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableRowBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableRowBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableRowBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableRowBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableRowBuilder) Class(names ...string) (r *VDataTableRowBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableRowBuilder) ClassIf(name string, add bool) (r *VDataTableRowBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableRowBuilder) On(name string, value string) (r *VDataTableRowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableRowBuilder) Bind(name string, value string) (r *VDataTableRowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableRowBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
