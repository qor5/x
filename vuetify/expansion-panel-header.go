package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpansionPanelHeaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpansionPanelHeader(children ...h.HTMLComponent) (r *VExpansionPanelHeaderBuilder) {
	r = &VExpansionPanelHeaderBuilder{
		tag: h.Tag("v-expansion-panel-header").Children(children...),
	}
	return
}

func (b *VExpansionPanelHeaderBuilder) Color(v string) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VExpansionPanelHeaderBuilder) DisableIconRotate(v bool) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr(":disable-icon-rotate", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelHeaderBuilder) ExpandIcon(v string) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VExpansionPanelHeaderBuilder) HideActions(v bool) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelHeaderBuilder) Ripple(v interface{}) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VExpansionPanelHeaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpansionPanelHeaderBuilder) Attr(vs ...interface{}) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpansionPanelHeaderBuilder) Children(children ...h.HTMLComponent) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpansionPanelHeaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpansionPanelHeaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpansionPanelHeaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpansionPanelHeaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpansionPanelHeaderBuilder) Class(names ...string) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpansionPanelHeaderBuilder) ClassIf(name string, add bool) (r *VExpansionPanelHeaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpansionPanelHeaderBuilder) On(name string, value string) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpansionPanelHeaderBuilder) Bind(name string, value string) (r *VExpansionPanelHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpansionPanelHeaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
