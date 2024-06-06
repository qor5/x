package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VPaginationBuilder struct {
	tag *h.HTMLTagBuilder
}

func VPagination(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	r = &VPaginationBuilder{
		tag: h.Tag("v-pagination").Children(children...),
	}
	return
}

func (b *VPaginationBuilder) Length(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":length", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) ActiveColor(v string) (r *VPaginationBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VPaginationBuilder) Start(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":start", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) ModelValue(v int) (r *VPaginationBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Disabled(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) TotalVisible(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":total-visible", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) FirstIcon(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":first-icon", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) PrevIcon(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) NextIcon(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) LastIcon(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":last-icon", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) AriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("aria-label", v)
	return b
}

func (b *VPaginationBuilder) PageAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("page-aria-label", v)
	return b
}

func (b *VPaginationBuilder) CurrentPageAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("current-page-aria-label", v)
	return b
}

func (b *VPaginationBuilder) FirstAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("first-aria-label", v)
	return b
}

func (b *VPaginationBuilder) PreviousAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("previous-aria-label", v)
	return b
}

func (b *VPaginationBuilder) NextAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("next-aria-label", v)
	return b
}

func (b *VPaginationBuilder) LastAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("last-aria-label", v)
	return b
}

func (b *VPaginationBuilder) Ellipsis(v string) (r *VPaginationBuilder) {
	b.tag.Attr("ellipsis", v)
	return b
}

func (b *VPaginationBuilder) ShowFirstLastPage(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":show-first-last-page", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Border(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) Density(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) Elevation(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) Rounded(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) Tile(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Size(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) Tag(v string) (r *VPaginationBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VPaginationBuilder) Theme(v string) (r *VPaginationBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VPaginationBuilder) Color(v string) (r *VPaginationBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VPaginationBuilder) Variant(v interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VPaginationBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VPaginationBuilder) Attr(vs ...interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VPaginationBuilder) Children(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VPaginationBuilder) AppendChildren(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VPaginationBuilder) PrependChildren(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VPaginationBuilder) Class(names ...string) (r *VPaginationBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VPaginationBuilder) ClassIf(name string, add bool) (r *VPaginationBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VPaginationBuilder) On(name string, value string) (r *VPaginationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VPaginationBuilder) Bind(name string, value string) (r *VPaginationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VPaginationBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
