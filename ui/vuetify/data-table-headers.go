package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableHeadersBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableHeaders(children ...h.HTMLComponent) (r *VDataTableHeadersBuilder) {
	r = &VDataTableHeadersBuilder{
		tag: h.Tag("v-data-table-headers").Children(children...),
	}
	return
}

func (b *VDataTableHeadersBuilder) Color(v string) (r *VDataTableHeadersBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDataTableHeadersBuilder) Sticky(v bool) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeadersBuilder) DisableSort(v bool) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeadersBuilder) MultiSort(v bool) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":multi-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeadersBuilder) SortAscIcon(v interface{}) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":sort-asc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableHeadersBuilder) SortDescIcon(v interface{}) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":sort-desc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableHeadersBuilder) HeaderProps(v interface{}) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":header-props", h.JSONString(v))
	return b
}

func (b *VDataTableHeadersBuilder) Mobile(v bool) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeadersBuilder) MobileBreakpoint(v interface{}) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VDataTableHeadersBuilder) Loading(v interface{}) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VDataTableHeadersBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableHeadersBuilder) Attr(vs ...interface{}) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableHeadersBuilder) Children(children ...h.HTMLComponent) (r *VDataTableHeadersBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableHeadersBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableHeadersBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableHeadersBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableHeadersBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableHeadersBuilder) Class(names ...string) (r *VDataTableHeadersBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableHeadersBuilder) ClassIf(name string, add bool) (r *VDataTableHeadersBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableHeadersBuilder) On(name string, value string) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableHeadersBuilder) Bind(name string, value string) (r *VDataTableHeadersBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableHeadersBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
