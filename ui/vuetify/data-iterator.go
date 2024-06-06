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

func (b *VDataIteratorBuilder) Search(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VDataIteratorBuilder) Loading(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Items(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ItemValue(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ItemSelectable(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":item-selectable", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ReturnObject(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) ShowSelect(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":show-select", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) SelectStrategy(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) Page(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":page", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ModelValue(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ValueComparator(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) SortBy(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":sort-by", h.JSONString(v))
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

func (b *VDataIteratorBuilder) CustomKeySort(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":custom-key-sort", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ItemsPerPage(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":items-per-page", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) ExpandOnClick(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":expand-on-click", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) ShowExpand(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":show-expand", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) Expanded(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":expanded", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) GroupBy(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":group-by", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) FilterMode(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":filter-mode", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) NoFilter(v bool) (r *VDataIteratorBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VDataIteratorBuilder) CustomFilter(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) CustomKeyFilter(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":custom-key-filter", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) FilterKeys(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":filter-keys", h.JSONString(v))
	return b
}

func (b *VDataIteratorBuilder) Tag(v string) (r *VDataIteratorBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VDataIteratorBuilder) Transition(v interface{}) (r *VDataIteratorBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
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
