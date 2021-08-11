package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataFooterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataFooter(children ...h.HTMLComponent) (r *VDataFooterBuilder) {
	r = &VDataFooterBuilder{
		tag: h.Tag("v-data-footer").Children(children...),
	}
	return
}

func (b *VDataFooterBuilder) DisableItemsPerPage(v bool) (r *VDataFooterBuilder) {
	b.tag.Attr(":disable-items-per-page", fmt.Sprint(v))
	return b
}

func (b *VDataFooterBuilder) DisablePagination(v bool) (r *VDataFooterBuilder) {
	b.tag.Attr(":disable-pagination", fmt.Sprint(v))
	return b
}

func (b *VDataFooterBuilder) FirstIcon(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("first-icon", v)
	return b
}

func (b *VDataFooterBuilder) ItemsPerPageAllText(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("items-per-page-all-text", v)
	return b
}

func (b *VDataFooterBuilder) ItemsPerPageOptions(v interface{}) (r *VDataFooterBuilder) {
	b.tag.Attr(":items-per-page-options", h.JSONString(v))
	return b
}

func (b *VDataFooterBuilder) ItemsPerPageText(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("items-per-page-text", v)
	return b
}

func (b *VDataFooterBuilder) LastIcon(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("last-icon", v)
	return b
}

func (b *VDataFooterBuilder) NextIcon(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VDataFooterBuilder) Options(v interface{}) (r *VDataFooterBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VDataFooterBuilder) PageText(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("page-text", v)
	return b
}

func (b *VDataFooterBuilder) Pagination(v interface{}) (r *VDataFooterBuilder) {
	b.tag.Attr(":pagination", h.JSONString(v))
	return b
}

func (b *VDataFooterBuilder) PrevIcon(v string) (r *VDataFooterBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VDataFooterBuilder) ShowCurrentPage(v bool) (r *VDataFooterBuilder) {
	b.tag.Attr(":show-current-page", fmt.Sprint(v))
	return b
}

func (b *VDataFooterBuilder) ShowFirstLastPage(v bool) (r *VDataFooterBuilder) {
	b.tag.Attr(":show-first-last-page", fmt.Sprint(v))
	return b
}

func (b *VDataFooterBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataFooterBuilder) Attr(vs ...interface{}) (r *VDataFooterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataFooterBuilder) Children(children ...h.HTMLComponent) (r *VDataFooterBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataFooterBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataFooterBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataFooterBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataFooterBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataFooterBuilder) Class(names ...string) (r *VDataFooterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataFooterBuilder) ClassIf(name string, add bool) (r *VDataFooterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataFooterBuilder) On(name string, value string) (r *VDataFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataFooterBuilder) Bind(name string, value string) (r *VDataFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataFooterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
