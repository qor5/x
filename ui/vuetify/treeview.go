package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTreeviewBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTreeview(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	r = &VTreeviewBuilder{
		tag: h.Tag("v-treeview").Children(children...),
	}
	return
}

func (b *VTreeviewBuilder) OpenAll(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":open-all", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Search(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("search", v)
	return b
}

func (b *VTreeviewBuilder) FilterMode(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":filter-mode", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) NoFilter(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) CustomFilter(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":custom-filter", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) CustomKeyFilter(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":custom-key-filter", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) FilterKeys(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":filter-keys", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) LoadingIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("loading-icon", v)
	return b
}

func (b *VTreeviewBuilder) Selectable(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":selectable", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) LoadChildren(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":load-children", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Items(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) SelectStrategy(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) BaseColor(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VTreeviewBuilder) ActiveColor(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VTreeviewBuilder) ActiveClass(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTreeviewBuilder) BgColor(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VTreeviewBuilder) Disabled(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) ExpandIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VTreeviewBuilder) CollapseIcon(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("collapse-icon", v)
	return b
}

func (b *VTreeviewBuilder) Lines(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":lines", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Slim(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Activatable(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":activatable", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Opened(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":opened", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Activated(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":activated", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Selected(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":selected", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Mandatory(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) ActiveStrategy(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":active-strategy", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) OpenStrategy(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":open-strategy", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Border(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Density(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Height(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) MaxHeight(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) MaxWidth(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) MinHeight(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) MinWidth(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Width(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Elevation(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) ItemType(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("item-type", v)
	return b
}

func (b *VTreeviewBuilder) ItemTitle(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":item-title", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) ItemValue(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) ItemChildren(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":item-children", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) ItemProps(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":item-props", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) ReturnObject(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) ValueComparator(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Rounded(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) Tile(v bool) (r *VTreeviewBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTreeviewBuilder) Tag(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTreeviewBuilder) Theme(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTreeviewBuilder) Color(v string) (r *VTreeviewBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTreeviewBuilder) Variant(v interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VTreeviewBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTreeviewBuilder) Attr(vs ...interface{}) (r *VTreeviewBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTreeviewBuilder) Children(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTreeviewBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTreeviewBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTreeviewBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTreeviewBuilder) Class(names ...string) (r *VTreeviewBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTreeviewBuilder) ClassIf(name string, add bool) (r *VTreeviewBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTreeviewBuilder) On(name string, value string) (r *VTreeviewBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTreeviewBuilder) Bind(name string, value string) (r *VTreeviewBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTreeviewBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
