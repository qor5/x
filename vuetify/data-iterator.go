package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataIteratorBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataIterator(children ...h.HTMLComponent) (r *VDataIteratorBuilder) {
	r = &VDataIteratorBuilder{
		tag: h.Tag("v-data-iterator").Children(children...),
	}
	return
}

func (b *VDataIteratorBuilder) CheckboxColor(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("checkbox-color", v)
	return b
}

func (b *VDataIteratorBuilder) CustomFilter(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) CustomGroup(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":custom-group", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) CustomSort(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":custom-sort", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) Dark(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) DisableFiltering(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":disable-filtering", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) DisablePagination(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":disable-pagination", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) DisableSort(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Expanded(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":expanded", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) FooterProps(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":footer-props", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) GroupBy(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("group-by", v)
	return b
}

func (b *VDataIteratorBuilder) GroupDesc(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":group-desc", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) HideDefaultFooter(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":hide-default-footer", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) ItemKey(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("item-key", v)
	return b
}

func (b *VDataIteratorBuilder) Items(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ItemsPerPage(v int) (r *VDataIteratorBuilder) {
	b.tag.Attr(":items-per-page", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Light(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Loading(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) LoadingText(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VDataIteratorBuilder) Locale(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VDataIteratorBuilder) MobileBreakpoint(v int) (r *VDataIteratorBuilder) {
	b.tag.Attr(":mobile-breakpoint", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) MultiSort(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":multi-sort", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) MustSort(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":must-sort", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) NoDataText(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VDataIteratorBuilder) NoResultsText(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("no-results-text", v)
	return b
}

func (b *VDataIteratorBuilder) Options(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) Page(v int) (r *VDataIteratorBuilder) {
	b.tag.Attr(":page", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Search(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VDataIteratorBuilder) SelectableKey(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("selectable-key", v)
	return b
}

func (b *VDataIteratorBuilder) ServerItemsLength(v int) (r *VDataIteratorBuilder) {
	b.tag.Attr(":server-items-length", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) SingleExpand(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":single-expand", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) SingleSelect(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":single-select", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) SortBy(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("sort-by", v)
	return b
}

func (b *VDataIteratorBuilder) SortDesc(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":sort-desc", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Value(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataIteratorBuilder) Attr(vs ...interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataIteratorBuilder) Children(children ...h.HTMLComponent) (r *VDataIteratorBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataIteratorBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataIteratorBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataIteratorBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataIteratorBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataIteratorBuilder) Class(names ...string) (r *VDataIteratorBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataIteratorBuilder) ClassIf(name string, add bool) (r *VDataIteratorBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataIteratorBuilder) On(name string, value string) (r *VDataIteratorBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataIteratorBuilder) Bind(name string, value string) (r *VDataIteratorBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataIteratorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
