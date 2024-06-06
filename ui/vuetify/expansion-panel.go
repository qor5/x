package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpansionPanelBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpansionPanel(children ...h.HTMLComponent) (r *VExpansionPanelBuilder) {
	r = &VExpansionPanelBuilder{
		tag: h.Tag("v-expansion-panel").Children(children...),
	}
	return
}

func (b *VExpansionPanelBuilder) Title(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VExpansionPanelBuilder) Text(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VExpansionPanelBuilder) BgColor(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VExpansionPanelBuilder) Elevation(v interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VExpansionPanelBuilder) Value(v interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VExpansionPanelBuilder) Disabled(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) SelectedClass(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VExpansionPanelBuilder) Rounded(v interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VExpansionPanelBuilder) Tile(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Tag(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VExpansionPanelBuilder) Color(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VExpansionPanelBuilder) ExpandIcon(v interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VExpansionPanelBuilder) CollapseIcon(v interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VExpansionPanelBuilder) HideActions(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Focusable(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Static(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":static", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Ripple(v interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VExpansionPanelBuilder) Readonly(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Eager(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpansionPanelBuilder) Attr(vs ...interface{}) (r *VExpansionPanelBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpansionPanelBuilder) Children(children ...h.HTMLComponent) (r *VExpansionPanelBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpansionPanelBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpansionPanelBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpansionPanelBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpansionPanelBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpansionPanelBuilder) Class(names ...string) (r *VExpansionPanelBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpansionPanelBuilder) ClassIf(name string, add bool) (r *VExpansionPanelBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpansionPanelBuilder) On(name string, value string) (r *VExpansionPanelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpansionPanelBuilder) Bind(name string, value string) (r *VExpansionPanelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpansionPanelBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
