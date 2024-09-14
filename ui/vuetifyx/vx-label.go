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

func (b *VXLabelBuilder) Tooltip(v interface{}) (r *VXLabelBuilder) {
	b.tag.Attr("tooltip", fmt.Sprint(v))
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

func (b *VXLabelBuilder) Attr(vs ...interface{}) (r *VXLabelBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXLabelBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
