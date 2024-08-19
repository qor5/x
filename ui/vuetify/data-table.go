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

func (b *VDataTableBuilder) Width(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) HeaderProps(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":header-props", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) CellProps(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":cell-props", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Mobile(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Loading(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Headers(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":headers", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Page(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":page", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ItemsPerPage(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":items-per-page", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) LoadingText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VDataTableBuilder) HideNoData(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Items(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) NoDataText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VDataTableBuilder) MobileBreakpoint(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) RowProps(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":row-props", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) HideDefaultBody(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":hide-default-body", fmt.Sprint(v))
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

func (b *VDataTableBuilder) Search(v string) (r *VDataTableBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VDataTableBuilder) ExpandOnClick(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":expand-on-click", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ShowExpand(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":show-expand", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Expanded(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":expanded", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) GroupBy(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":group-by", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ItemValue(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ItemSelectable(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":item-selectable", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ReturnObject(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) ShowSelect(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":show-select", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) SelectStrategy(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ModelValue(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ValueComparator(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) SortBy(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":sort-by", h.JSONString(v))
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

func (b *VDataTableBuilder) CustomKeySort(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":custom-key-sort", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Color(v string) (r *VDataTableBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDataTableBuilder) Sticky(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) DisableSort(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) SortAscIcon(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":sort-asc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) SortDescIcon(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":sort-desc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) FixedHeader(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) FixedFooter(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":fixed-footer", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Height(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Hover(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) Density(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) Tag(v string) (r *VDataTableBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VDataTableBuilder) Theme(v string) (r *VDataTableBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDataTableBuilder) FilterMode(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":filter-mode", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) NoFilter(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VDataTableBuilder) CustomFilter(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) CustomKeyFilter(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":custom-key-filter", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) FilterKeys(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":filter-keys", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) PrevIcon(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) NextIcon(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) FirstIcon(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":first-icon", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) LastIcon(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":last-icon", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ItemsPerPageText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("items-per-page-text", v)
	return b
}

func (b *VDataTableBuilder) PageText(v string) (r *VDataTableBuilder) {
	b.tag.Attr("page-text", v)
	return b
}

func (b *VDataTableBuilder) FirstPageLabel(v string) (r *VDataTableBuilder) {
	b.tag.Attr("first-page-label", v)
	return b
}

func (b *VDataTableBuilder) PrevPageLabel(v string) (r *VDataTableBuilder) {
	b.tag.Attr("prev-page-label", v)
	return b
}

func (b *VDataTableBuilder) NextPageLabel(v string) (r *VDataTableBuilder) {
	b.tag.Attr("next-page-label", v)
	return b
}

func (b *VDataTableBuilder) LastPageLabel(v string) (r *VDataTableBuilder) {
	b.tag.Attr("last-page-label", v)
	return b
}

func (b *VDataTableBuilder) ItemsPerPageOptions(v interface{}) (r *VDataTableBuilder) {
	b.tag.Attr(":items-per-page-options", h.JSONString(v))
	return b
}

func (b *VDataTableBuilder) ShowCurrentPage(v bool) (r *VDataTableBuilder) {
	b.tag.Attr(":show-current-page", fmt.Sprint(v))
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
