package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSparklineBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSparkline(children ...h.HTMLComponent) (r *VSparklineBuilder) {
	r = &VSparklineBuilder{
		tag: h.Tag("v-sparkline").Children(children...),
	}
	return
}

func (b *VSparklineBuilder) AutoDraw(v bool) (r *VSparklineBuilder) {
	b.tag.Attr(":auto-draw", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) AutoDrawDuration(v int) (r *VSparklineBuilder) {
	b.tag.Attr(":auto-draw-duration", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) AutoDrawEasing(v string) (r *VSparklineBuilder) {
	b.tag.Attr("auto-draw-easing", v)
	return b
}

func (b *VSparklineBuilder) AutoLineWidth(v bool) (r *VSparklineBuilder) {
	b.tag.Attr(":auto-line-width", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) Color(v string) (r *VSparklineBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSparklineBuilder) Fill(v bool) (r *VSparklineBuilder) {
	b.tag.Attr(":fill", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) Gradient(v interface{}) (r *VSparklineBuilder) {
	b.tag.Attr(":gradient", h.JSONString(v))
	return b
}

func (b *VSparklineBuilder) GradientDirection(v string) (r *VSparklineBuilder) {
	b.tag.Attr("gradient-direction", v)
	return b
}

func (b *VSparklineBuilder) Height(v string) (r *VSparklineBuilder) {
	b.tag.Attr("height", v)
	return b
}

func (b *VSparklineBuilder) LabelSize(v int) (r *VSparklineBuilder) {
	b.tag.Attr(":label-size", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) Labels(v interface{}) (r *VSparklineBuilder) {
	b.tag.Attr(":labels", h.JSONString(v))
	return b
}

func (b *VSparklineBuilder) LineWidth(v string) (r *VSparklineBuilder) {
	b.tag.Attr("line-width", v)
	return b
}

func (b *VSparklineBuilder) Padding(v string) (r *VSparklineBuilder) {
	b.tag.Attr("padding", v)
	return b
}

func (b *VSparklineBuilder) ShowLabels(v bool) (r *VSparklineBuilder) {
	b.tag.Attr(":show-labels", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) Smooth(v int) (r *VSparklineBuilder) {
	b.tag.Attr(":smooth", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) Type(v string) (r *VSparklineBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VSparklineBuilder) Value(v interface{}) (r *VSparklineBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSparklineBuilder) Width(v int) (r *VSparklineBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VSparklineBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSparklineBuilder) Attr(vs ...interface{}) (r *VSparklineBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSparklineBuilder) Children(children ...h.HTMLComponent) (r *VSparklineBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSparklineBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSparklineBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSparklineBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSparklineBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSparklineBuilder) Class(names ...string) (r *VSparklineBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSparklineBuilder) ClassIf(name string, add bool) (r *VSparklineBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSparklineBuilder) On(name string, value string) (r *VSparklineBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSparklineBuilder) Bind(name string, value string) (r *VSparklineBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSparklineBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
