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

func (b *VExpansionPanelsBuilder) Accordion(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":accordion", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) ActiveClass(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VExpansionPanelsBuilder) Dark(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Disabled(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Flat(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Focusable(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Hover(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Inset(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Light(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Mandatory(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Max(v int) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Multiple(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Popout(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":popout", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Readonly(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Tag(v string) (r *VExpansionPanelsBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VExpansionPanelsBuilder) Tile(v bool) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelsBuilder) Value(v interface{}) (r *VExpansionPanelsBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
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
