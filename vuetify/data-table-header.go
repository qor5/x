package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableHeaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableHeader(children ...h.HTMLComponent) (r *VDataTableHeaderBuilder) {
	r = &VDataTableHeaderBuilder{
		tag: h.Tag("v-data-table-header").Children(children...),
	}
	return
}

func (b *VDataTableHeaderBuilder) CheckboxColor(v string) (r *VDataTableHeaderBuilder) {
	b.tag.Attr("checkbox-color", v)
	return b
}

func (b *VDataTableHeaderBuilder) DisableSort(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) EveryItem(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":every-item", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) Headers(v interface{}) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":headers", h.JSONString(v))
	return b
}

func (b *VDataTableHeaderBuilder) Mobile(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) Options(v interface{}) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VDataTableHeaderBuilder) ShowGroupBy(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":show-group-by", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) SingleSelect(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":single-select", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) SomeItems(v bool) (r *VDataTableHeaderBuilder) {
	b.tag.Attr(":some-items", fmt.Sprint(v))
	return b
}

func (b *VDataTableHeaderBuilder) SortByText(v string) (r *VDataTableHeaderBuilder) {
	b.tag.Attr("sort-by-text", v)
	return b
}

func (b *VDataTableHeaderBuilder) SortIcon(v string) (r *VDataTableHeaderBuilder) {
	b.tag.Attr("sort-icon", v)
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
