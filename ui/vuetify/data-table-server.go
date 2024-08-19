package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableServerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableServer(children ...h.HTMLComponent) (r *VDataTableServerBuilder) {
	r = &VDataTableServerBuilder{
		tag: h.Tag("v-data-table-server").Children(children...),
	}
	return
}

func (b *VDataTableServerBuilder) Width(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) HeaderProps(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":header-props", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) CellProps(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":cell-props", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) Mobile(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) Loading(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) Headers(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":headers", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ItemsLength(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":items-length", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) Page(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":page", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ItemsPerPage(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":items-per-page", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) LoadingText(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VDataTableServerBuilder) HideNoData(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) Items(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) NoDataText(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VDataTableServerBuilder) MobileBreakpoint(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) RowProps(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":row-props", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) HideDefaultBody(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":hide-default-body", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) HideDefaultFooter(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":hide-default-footer", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) HideDefaultHeader(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":hide-default-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) Search(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VDataTableServerBuilder) ExpandOnClick(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":expand-on-click", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) ShowExpand(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":show-expand", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) Expanded(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":expanded", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) GroupBy(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":group-by", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ItemValue(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ItemSelectable(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":item-selectable", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ReturnObject(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) ShowSelect(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":show-select", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) SelectStrategy(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ModelValue(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ValueComparator(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) SortBy(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":sort-by", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) MultiSort(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":multi-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) MustSort(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":must-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) CustomKeySort(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":custom-key-sort", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) Color(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDataTableServerBuilder) Sticky(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) DisableSort(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) SortAscIcon(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":sort-asc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) SortDescIcon(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":sort-desc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) FixedHeader(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) FixedFooter(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":fixed-footer", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) Height(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) Hover(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) Density(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) Tag(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VDataTableServerBuilder) Theme(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDataTableServerBuilder) PrevIcon(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) NextIcon(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) FirstIcon(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":first-icon", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) LastIcon(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":last-icon", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ItemsPerPageText(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("items-per-page-text", v)
	return b
}

func (b *VDataTableServerBuilder) PageText(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("page-text", v)
	return b
}

func (b *VDataTableServerBuilder) FirstPageLabel(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("first-page-label", v)
	return b
}

func (b *VDataTableServerBuilder) PrevPageLabel(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("prev-page-label", v)
	return b
}

func (b *VDataTableServerBuilder) NextPageLabel(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("next-page-label", v)
	return b
}

func (b *VDataTableServerBuilder) LastPageLabel(v string) (r *VDataTableServerBuilder) {
	b.tag.Attr("last-page-label", v)
	return b
}

func (b *VDataTableServerBuilder) ItemsPerPageOptions(v interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(":items-per-page-options", h.JSONString(v))
	return b
}

func (b *VDataTableServerBuilder) ShowCurrentPage(v bool) (r *VDataTableServerBuilder) {
	b.tag.Attr(":show-current-page", fmt.Sprint(v))
	return b
}

func (b *VDataTableServerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableServerBuilder) Attr(vs ...interface{}) (r *VDataTableServerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableServerBuilder) Children(children ...h.HTMLComponent) (r *VDataTableServerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableServerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableServerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableServerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableServerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableServerBuilder) Class(names ...string) (r *VDataTableServerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableServerBuilder) ClassIf(name string, add bool) (r *VDataTableServerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableServerBuilder) On(name string, value string) (r *VDataTableServerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableServerBuilder) Bind(name string, value string) (r *VDataTableServerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableServerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
