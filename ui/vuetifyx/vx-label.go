package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXLabelBuilder struct {
	tag           *h.HTMLTagBuilder
}

func VXLabel(children ...h.HTMLComponent) (r *VXLabelBuilder) {
	r = &VXLabelBuilder{
		tag: h.Tag("vx-label").Children(children...),
	}
	return
}

func (b *VXLabelBuilder) ToggleLabel(v bool) (r *VXLabelBuilder) {
	b.tag.Attr(":toggle-label", fmt.Sprint((v)))
	return b
}

func (b *VXLabelBuilder) Class(v string) (r *VXLabelBuilder) {
	b.tag.Attr("class", v)
	return b
}

func (b *VXLabelBuilder) ToggleIconSize(v string) (r *VXLabelBuilder) {
	b.tag.Attr("toggle-icon-size", v)
	return b
}

func (b *VXLabelBuilder) Tooltip(v interface{}) (r *VXLabelBuilder) {
	b.tag.Attr("tooltip", fmt.Sprint(v))
	return b
}

func (b *VXLabelBuilder) RequiredSymbol(v bool) (r *VXLabelBuilder) {
	b.tag.Attr(":required-symbol", fmt.Sprint(v))
	return b
}

func (b *VXLabelBuilder) TooltipIconColor(v string) (r *VXLabelBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXLabelBuilder) Icon(v string) (r *VXLabelBuilder) {
	b.tag.Attr("icon", v)
	return b
}

func (b *VXLabelBuilder) IconSize(v string) (r *VXLabelBuilder) {
	b.tag.Attr("iconSize", v)
	return b
}

func (b *VXLabelBuilder) Slot(name string, children ...h.HTMLComponent) *VXLabelBuilder {
	slotTemplate := h.Tag("template").Attr("#"+name).Children(children...)
	b.tag.Children(slotTemplate)
	return b
}

func (b *VXLabelBuilder) Attr(vs ...interface{}) (r *VXLabelBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXLabelBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
