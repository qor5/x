package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpansionPanelsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpansionPanels(children ...h.HTMLComponent) (r *VExpansionPanelsBuilder) {
	r = &VExpansionPanelsBuilder{
		tag: h.Tag("v-expansion-panels").Children(children...),
	}
	return
}

func (b *VExpansionPanelsBuilder) Flat(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) ModelValue(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) Multiple(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Max(v int) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) SelectedClass(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VExpansionPanelsBuilder) Disabled(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Mandatory(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) BgColor(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VExpansionPanelsBuilder) Elevation(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) Rounded(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) Tile(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Tag(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VExpansionPanelsBuilder) Color(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VExpansionPanelsBuilder) ExpandIcon(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) CollapseIcon(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) HideActions(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Focusable(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Static(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":static", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Ripple(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) Readonly(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Eager(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Theme(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VExpansionPanelsBuilder) Variant(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VExpansionPanelsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpansionPanelsBuilder) Attr(vs ...interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpansionPanelsBuilder) Children(children ...h.HTMLComponent) (r *VExpansionPanelsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpansionPanelsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpansionPanelsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpansionPanelsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpansionPanelsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpansionPanelsBuilder) Class(names ...string) (r *VExpansionPanelsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpansionPanelsBuilder) ClassIf(name string, add bool) (r *VExpansionPanelsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpansionPanelsBuilder) On(name string, value string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpansionPanelsBuilder) Bind(name string, value string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpansionPanelsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
