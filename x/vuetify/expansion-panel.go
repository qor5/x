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

func (b *VExpansionPanelBuilder) ActiveClass(v string) (r *VExpansionPanelBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VExpansionPanelBuilder) Disabled(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Readonly(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
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
