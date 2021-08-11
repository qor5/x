package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTable(children ...h.HTMLComponent) (r *VDataTableBuilder) {
	r = &VDataTableBuilder{
		tag: h.Tag("v-data-table").Children(children...),
	}
	return
}

func (b *VDataTableBuilder) CalculateWidths(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":calculate-widths", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Caption(v string) (r *VDataTableBuilder) {
	b.tag.Attr("caption", v)
	return b
}

func (b *VDataTableBuilder) CheckboxColor(v string) (r *VDataTableBuilder) {
	b.tag.Attr("checkbox-color", v)
	return b
}

func (b *VDataTableBuilder) CustomFilter(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) CustomGroup(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":custom-group", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) CustomSort(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":custom-sort", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Dark(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Dense(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) DisableFiltering(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":disable-filtering", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) DisablePagination(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":disable-pagination", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) DisableSort(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ExpandIcon(v string) (r *VDataTableBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VDataTableBuilder) Expanded(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":expanded", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) FixedHeader(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) FooterProps(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":footer-props", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) GroupBy(v string) (r *VDataTableBuilder) {
	b.tag.Attr("group-by", v)
	return b
}

func (b *VDataTableBuilder) GroupDesc(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":group-desc", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) HeaderProps(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":header-props", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Headers(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":headers", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) HeadersLength(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":headers-length", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Height(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) HideDefaultFooter(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":hide-default-footer", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) HideDefaultHeader(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":hide-default-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ItemClass(v string) (r *VDataTableBuilder) {
	b.tag.Attr("item-class", v)
	return b
}

func (b *VDataTableBuilder) ItemKey(v string) (r *VDataTableBuilder) {
	b.tag.Attr("item-key", v)
	return b
}

func (b *VDataTableBuilder) Items(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ItemsPerPage(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":items-per-page", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Light(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) LoaderHeight(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Loading(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) LoadingText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VDataTableBuilder) Locale(v string) (r *VDataTableBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VDataTableBuilder) MobileBreakpoint(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":mobile-breakpoint", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) MultiSort(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":multi-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) MustSort(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":must-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) NoDataText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VDataTableBuilder) NoResultsText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("no-results-text", v)
	return b
}

func (b *VDataTableBuilder) Options(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Page(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":page", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Search(v string) (r *VDataTableBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VDataTableBuilder) SelectableKey(v string) (r *VDataTableBuilder) {
	b.tag.Attr("selectable-key", v)
	return b
}

func (b *VDataTableBuilder) ServerItemsLength(v int) (r *VDataTableBuilder) {
	b.tag.Attr(":server-items-length", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ShowExpand(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":show-expand", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ShowGroupBy(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":show-group-by", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ShowSelect(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":show-select", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) SingleExpand(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":single-expand", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) SingleSelect(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":single-select", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) SortBy(v string) (r *VDataTableBuilder) {
	b.tag.Attr("sort-by", v)
	return b
}

func (b *VDataTableBuilder) SortDesc(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":sort-desc", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Value(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableBuilder) Attr(vs ...interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableBuilder) Children(children ...h.HTMLComponent) (r *VDataTableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableBuilder) Class(names ...string) (r *VDataTableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableBuilder) ClassIf(name string, add bool) (r *VDataTableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableBuilder) On(name string, value string) (r *VDataTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableBuilder) Bind(name string, value string) (r *VDataTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
