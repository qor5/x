package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXPaginationBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXPagination(children ...h.HTMLComponent) (r *VXPaginationBuilder) {
	r = &VXPaginationBuilder{
		tag: h.Tag("vx-pagination").Children(children...),
	}
	return
}

func (b *VXPaginationBuilder) Length(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":length", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) ActiveColor(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VXPaginationBuilder) Start(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":start", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) ModelValue(v int) (r *VXPaginationBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VXPaginationBuilder) Disabled(v bool) (r *VXPaginationBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXPaginationBuilder) TotalVisible(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":total-visible", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) FirstIcon(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":first-icon", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) PrevIcon(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) NextIcon(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) LastIcon(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":last-icon", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) AriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("aria-label", v)
	return b
}

func (b *VXPaginationBuilder) PageAriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("page-aria-label", v)
	return b
}

func (b *VXPaginationBuilder) CurrentPageAriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("current-page-aria-label", v)
	return b
}

func (b *VXPaginationBuilder) FirstAriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("first-aria-label", v)
	return b
}

func (b *VXPaginationBuilder) PreviousAriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("previous-aria-label", v)
	return b
}

func (b *VXPaginationBuilder) NextAriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("next-aria-label", v)
	return b
}

func (b *VXPaginationBuilder) LastAriaLabel(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("last-aria-label", v)
	return b
}

func (b *VXPaginationBuilder) Ellipsis(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("ellipsis", v)
	return b
}

func (b *VXPaginationBuilder) ShowFirstLastPage(v bool) (r *VXPaginationBuilder) {
	b.tag.Attr(":show-first-last-page", fmt.Sprint(v))
	return b
}

func (b *VXPaginationBuilder) Border(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) Density(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) Elevation(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) Rounded(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) Tile(v bool) (r *VXPaginationBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXPaginationBuilder) Size(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) Tag(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXPaginationBuilder) Theme(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXPaginationBuilder) Color(v string) (r *VXPaginationBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXPaginationBuilder) Variant(v interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXPaginationBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXPaginationBuilder) Attr(vs ...interface{}) (r *VXPaginationBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXPaginationBuilder) Children(children ...h.HTMLComponent) (r *VXPaginationBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXPaginationBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXPaginationBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXPaginationBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXPaginationBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXPaginationBuilder) Class(names ...string) (r *VXPaginationBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXPaginationBuilder) ClassIf(name string, add bool) (r *VXPaginationBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXPaginationBuilder) On(name string, value string) (r *VXPaginationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXPaginationBuilder) Bind(name string, value string) (r *VXPaginationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXPaginationBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
