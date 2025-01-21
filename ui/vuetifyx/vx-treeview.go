package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXTreeviewBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXTreeview(children ...h.HTMLComponent) (r *VXTreeviewBuilder) {
	r = &VXTreeviewBuilder{
		tag: h.Tag("vx-treeview").Children(children...),
	}
	return
}

func (b *VXTreeviewBuilder) OpenAll(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":open-all", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) Search(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VXTreeviewBuilder) FilterMode(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":filter-mode", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) NoFilter(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) CustomFilter(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) CustomKeyFilter(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":custom-key-filter", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) FilterKeys(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":filter-keys", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) LoadingIcon(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("loading-icon", v)
	return b
}

func (b *VXTreeviewBuilder) Selectable(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":selectable", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) LoadChildren(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":load-children", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Items(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) SelectStrategy(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) BaseColor(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VXTreeviewBuilder) ActiveColor(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VXTreeviewBuilder) ActiveClass(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VXTreeviewBuilder) BgColor(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VXTreeviewBuilder) Disabled(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) ExpandIcon(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VXTreeviewBuilder) CollapseIcon(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("collapse-icon", v)
	return b
}

func (b *VXTreeviewBuilder) Lines(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":lines", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Slim(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) Activatable(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":activatable", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) Opened(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":opened", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Activated(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":activated", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Selected(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":selected", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Mandatory(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) ActiveStrategy(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":active-strategy", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) OpenStrategy(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":open-strategy", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Border(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Density(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Height(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) MaxHeight(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) MaxWidth(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) MinHeight(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) MinWidth(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Width(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Elevation(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) ItemType(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("item-type", v)
	return b
}

func (b *VXTreeviewBuilder) ItemTitle(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":item-title", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) ItemValue(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) ItemChildren(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":item-children", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) ItemProps(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":item-props", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) ReturnObject(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) ValueComparator(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Rounded(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) Tile(v bool) (r *VXTreeviewBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXTreeviewBuilder) Tag(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXTreeviewBuilder) Theme(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXTreeviewBuilder) Color(v string) (r *VXTreeviewBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXTreeviewBuilder) Variant(v interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXTreeviewBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXTreeviewBuilder) Attr(vs ...interface{}) (r *VXTreeviewBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTreeviewBuilder) Children(children ...h.HTMLComponent) (r *VXTreeviewBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXTreeviewBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXTreeviewBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXTreeviewBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXTreeviewBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXTreeviewBuilder) Class(names ...string) (r *VXTreeviewBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXTreeviewBuilder) ClassIf(name string, add bool) (r *VXTreeviewBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXTreeviewBuilder) On(name string, value string) (r *VXTreeviewBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXTreeviewBuilder) Bind(name string, value string) (r *VXTreeviewBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXTreeviewBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
