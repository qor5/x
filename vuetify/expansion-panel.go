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

func (b *VExpansionPanelBuilder) Dark(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Disabled(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Expand(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":expand", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Focusable(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Inset(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Light(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Popout(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":popout", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Readonly(v bool) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelBuilder) Value(v int) (r *VExpansionPanelBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
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
