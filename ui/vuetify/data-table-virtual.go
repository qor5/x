package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDataTableVirtualBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDataTableVirtual(children ...h.HTMLComponent) (r *VDataTableVirtualBuilder) {
	r = &VDataTableVirtualBuilder{
		tag: h.Tag("v-data-table-virtual").Children(children...),
	}
	return
}

func (b *VDataTableVirtualBuilder) Width(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) HeaderProps(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":header-props", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) CellProps(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":cell-props", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) Mobile(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) Headers(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":headers", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) Loading(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) LoadingText(v string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VDataTableVirtualBuilder) HideNoData(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) Items(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) NoDataText(v string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VDataTableVirtualBuilder) MobileBreakpoint(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) RowProps(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":row-props", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) HideDefaultBody(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":hide-default-body", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) HideDefaultFooter(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":hide-default-footer", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) HideDefaultHeader(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":hide-default-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) Search(v string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VDataTableVirtualBuilder) ExpandOnClick(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":expand-on-click", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) ShowExpand(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":show-expand", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) Expanded(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":expanded", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) GroupBy(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":group-by", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) ItemValue(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) ItemSelectable(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":item-selectable", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) ReturnObject(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) ShowSelect(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":show-select", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) SelectStrategy(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) ModelValue(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) ValueComparator(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) SortBy(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":sort-by", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) MultiSort(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":multi-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) MustSort(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":must-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) CustomKeySort(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":custom-key-sort", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) Color(v string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDataTableVirtualBuilder) Sticky(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":sticky", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) DisableSort(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":disable-sort", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) SortAscIcon(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":sort-asc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) SortDescIcon(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":sort-desc-icon", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) FixedHeader(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) FixedFooter(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":fixed-footer", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) Height(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) Hover(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) Density(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) Tag(v string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VDataTableVirtualBuilder) Theme(v string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDataTableVirtualBuilder) ItemHeight(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":item-height", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) FilterMode(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":filter-mode", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) NoFilter(v bool) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VDataTableVirtualBuilder) CustomFilter(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) CustomKeyFilter(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":custom-key-filter", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) FilterKeys(v interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(":filter-keys", h.JSONString(v))
	return b
}

func (b *VDataTableVirtualBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDataTableVirtualBuilder) Attr(vs ...interface{}) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDataTableVirtualBuilder) Children(children ...h.HTMLComponent) (r *VDataTableVirtualBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDataTableVirtualBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDataTableVirtualBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDataTableVirtualBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDataTableVirtualBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDataTableVirtualBuilder) Class(names ...string) (r *VDataTableVirtualBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDataTableVirtualBuilder) ClassIf(name string, add bool) (r *VDataTableVirtualBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDataTableVirtualBuilder) On(name string, value string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDataTableVirtualBuilder) Bind(name string, value string) (r *VDataTableVirtualBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDataTableVirtualBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
