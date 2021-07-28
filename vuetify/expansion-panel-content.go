package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VExpansionPanelContentBuilder struct {
	tag *h.HTMLTagBuilder
}

func VExpansionPanelContent(children ...h.HTMLComponent) (r *VExpansionPanelContentBuilder) {
	r = &VExpansionPanelContentBuilder{
		tag: h.Tag("v-expansion-panel-content").Children(children...),
	}
	return
}

func (b *VExpansionPanelContentBuilder) Color(v string) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VExpansionPanelContentBuilder) Eager(v bool) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelContentBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VExpansionPanelContentBuilder) Attr(vs ...interface{}) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VExpansionPanelContentBuilder) Children(children ...h.HTMLComponent) (r *VExpansionPanelContentBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VExpansionPanelContentBuilder) AppendChildren(children ...h.HTMLComponent) (r *VExpansionPanelContentBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VExpansionPanelContentBuilder) PrependChildren(children ...h.HTMLComponent) (r *VExpansionPanelContentBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VExpansionPanelContentBuilder) Class(names ...string) (r *VExpansionPanelContentBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VExpansionPanelContentBuilder) ClassIf(name string, add bool) (r *VExpansionPanelContentBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VExpansionPanelContentBuilder) On(name string, value string) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VExpansionPanelContentBuilder) Bind(name string, value string) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VExpansionPanelContentBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
