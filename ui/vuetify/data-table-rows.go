package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableRowsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableRows(children ...h.HTMLComponent) (r *VDataTableRowsBuilder) {
	r = &VDataTableRowsBuilder{
		tag: h.Tag("v-data-table-rows").Children(children...),
	}
	return
}

func (b *VDataTableRowsBuilder) CellProps(v interface{}) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":cell-props", h.JSONString(v))
	return b
}

func (b *VDataTableRowsBuilder) Mobile(v bool) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableRowsBuilder) Loading(v interface{}) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VDataTableRowsBuilder) LoadingText(v string) (r *VDataTableRowsBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VDataTableRowsBuilder) HideNoData(v bool) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VDataTableRowsBuilder) Items(v interface{}) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataTableRowsBuilder) NoDataText(v string) (r *VDataTableRowsBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VDataTableRowsBuilder) MobileBreakpoint(v interface{}) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VDataTableRowsBuilder) RowProps(v interface{}) (r *VDataTableRowsBuilder) {
	b.tag.Attr(":row-props", h.JSONString(v))
	return b
}

func (b *VDataTableRowsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableRowsBuilder) Attr(vs ...interface{}) (r *VDataTableRowsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableRowsBuilder) Children(children ...h.HTMLComponent) (r *VDataTableRowsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableRowsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableRowsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableRowsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableRowsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableRowsBuilder) Class(names ...string) (r *VDataTableRowsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableRowsBuilder) ClassIf(name string, add bool) (r *VDataTableRowsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableRowsBuilder) On(name string, value string) (r *VDataTableRowsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableRowsBuilder) Bind(name string, value string) (r *VDataTableRowsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableRowsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
