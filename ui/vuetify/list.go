package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListBuilder struct {
	tag *h.HTMLTagBuilder
}

func VList(children ...h.HTMLComponent) (r *VListBuilder) {
	r = &VListBuilder{
		tag: h.Tag("v-list").Children(children...),
	}
	return
}

func (b *VListBuilder) BaseColor(v string) (r *VListBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VListBuilder) ActiveColor(v string) (r *VListBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VListBuilder) ActiveClass(v string) (r *VListBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListBuilder) BgColor(v string) (r *VListBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VListBuilder) Disabled(v bool) (r *VListBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) ExpandIcon(v string) (r *VListBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VListBuilder) CollapseIcon(v string) (r *VListBuilder) {
	b.tag.Attr("collapse-icon", v)
	return b
}

func (b *VListBuilder) Lines(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":lines", h.JSONString(v))
	return b
}

func (b *VListBuilder) Slim(v bool) (r *VListBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Nav(v bool) (r *VListBuilder) {
	b.tag.Attr(":nav", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Activatable(v bool) (r *VListBuilder) {
	b.tag.Attr(":activatable", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Selectable(v bool) (r *VListBuilder) {
	b.tag.Attr(":selectable", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Opened(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":opened", h.JSONString(v))
	return b
}

func (b *VListBuilder) Activated(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":activated", h.JSONString(v))
	return b
}

func (b *VListBuilder) Selected(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":selected", h.JSONString(v))
	return b
}

func (b *VListBuilder) Mandatory(v bool) (r *VListBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) ActiveStrategy(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":active-strategy", h.JSONString(v))
	return b
}

func (b *VListBuilder) SelectStrategy(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":select-strategy", h.JSONString(v))
	return b
}

func (b *VListBuilder) OpenStrategy(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":open-strategy", h.JSONString(v))
	return b
}

func (b *VListBuilder) Border(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VListBuilder) Density(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VListBuilder) Height(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VListBuilder) MaxHeight(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VListBuilder) MaxWidth(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VListBuilder) MinHeight(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VListBuilder) MinWidth(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VListBuilder) Width(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VListBuilder) Elevation(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VListBuilder) ItemType(v string) (r *VListBuilder) {
	b.tag.Attr("item-type", v)
	return b
}

func (b *VListBuilder) Items(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VListBuilder) ItemTitle(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":item-title", h.JSONString(v))
	return b
}

func (b *VListBuilder) ItemValue(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}

func (b *VListBuilder) ItemChildren(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":item-children", h.JSONString(v))
	return b
}

func (b *VListBuilder) ItemProps(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":item-props", h.JSONString(v))
	return b
}

func (b *VListBuilder) ReturnObject(v bool) (r *VListBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) ValueComparator(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VListBuilder) Rounded(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VListBuilder) Tile(v bool) (r *VListBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Tag(v string) (r *VListBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListBuilder) Theme(v string) (r *VListBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VListBuilder) Color(v string) (r *VListBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListBuilder) Variant(v interface{}) (r *VListBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VListBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListBuilder) Attr(vs ...interface{}) (r *VListBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListBuilder) Children(children ...h.HTMLComponent) (r *VListBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListBuilder) Class(names ...string) (r *VListBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListBuilder) ClassIf(name string, add bool) (r *VListBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListBuilder) On(name string, value string) (r *VListBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListBuilder) Bind(name string, value string) (r *VListBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
