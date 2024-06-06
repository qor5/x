package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpansionPanelTextBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpansionPanelText(children ...h.HTMLComponent) (r *VExpansionPanelTextBuilder) {
	r = &VExpansionPanelTextBuilder{
		tag: h.Tag("v-expansion-panel-text").Children(children...),
	}
	return
}

func (b *VExpansionPanelTextBuilder) Eager(v bool) (r *VExpansionPanelTextBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelTextBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpansionPanelTextBuilder) Attr(vs ...interface{}) (r *VExpansionPanelTextBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpansionPanelTextBuilder) Children(children ...h.HTMLComponent) (r *VExpansionPanelTextBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpansionPanelTextBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpansionPanelTextBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpansionPanelTextBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpansionPanelTextBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpansionPanelTextBuilder) Class(names ...string) (r *VExpansionPanelTextBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpansionPanelTextBuilder) ClassIf(name string, add bool) (r *VExpansionPanelTextBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpansionPanelTextBuilder) On(name string, value string) (r *VExpansionPanelTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpansionPanelTextBuilder) Bind(name string, value string) (r *VExpansionPanelTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpansionPanelTextBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
