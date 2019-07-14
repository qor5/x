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

func (b *VExpansionPanelContentBuilder) Disabled(v bool) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelContentBuilder) ExpandIcon(v string) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr("expand-icon", v)
	return b
}

func (b *VExpansionPanelContentBuilder) HideActions(v bool) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelContentBuilder) Lazy(v bool) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelContentBuilder) Readonly(v bool) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelContentBuilder) Ripple(v bool) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VExpansionPanelContentBuilder) Value(v interface{}) (r *VExpansionPanelContentBuilder) {
	b.tag.Attr(":value", v)
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
