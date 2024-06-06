package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpansionPanelTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpansionPanelTitle(children ...h.HTMLComponent) (r *VExpansionPanelTitleBuilder) {
	r = &VExpansionPanelTitleBuilder{
		tag: h.Tag("v-expansion-panel-title").Children(children...),
	}
	return
}

func (b *VExpansionPanelTitleBuilder) Color(v string) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VExpansionPanelTitleBuilder) ExpandIcon(v interface{}) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) CollapseIcon(v interface{}) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) HideActions(v bool) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) Focusable(v bool) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) Static(v bool) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":static", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) Ripple(v interface{}) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) Readonly(v bool) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpansionPanelTitleBuilder) Attr(vs ...interface{}) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpansionPanelTitleBuilder) Children(children ...h.HTMLComponent) (r *VExpansionPanelTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpansionPanelTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpansionPanelTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpansionPanelTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpansionPanelTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpansionPanelTitleBuilder) Class(names ...string) (r *VExpansionPanelTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpansionPanelTitleBuilder) ClassIf(name string, add bool) (r *VExpansionPanelTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpansionPanelTitleBuilder) On(name string, value string) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpansionPanelTitleBuilder) Bind(name string, value string) (r *VExpansionPanelTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpansionPanelTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
