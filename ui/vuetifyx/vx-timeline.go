package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXTimelineBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXTimeline(children ...h.HTMLComponent) (r *VXTimelineBuilder) {
	r = &VXTimelineBuilder{
		tag: h.Tag("vx-timeline").Children(children...),
	}
	return
}

func (b *VXTimelineBuilder) Justify(v string) (r *VXTimelineBuilder) {
	b.tag.Attr("justify", v)
	return b
}

func (b *VXTimelineBuilder) LineThickness(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":line-thickness", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) LineColor(v string) (r *VXTimelineBuilder) {
	b.tag.Attr("line-color", v)
	return b
}

func (b *VXTimelineBuilder) DotColor(v string) (r *VXTimelineBuilder) {
	b.tag.Attr("dot-color", v)
	return b
}

func (b *VXTimelineBuilder) FillDot(v bool) (r *VXTimelineBuilder) {
	b.tag.Attr(":fill-dot", fmt.Sprint(v))
	return b
}

func (b *VXTimelineBuilder) HideOpposite(v bool) (r *VXTimelineBuilder) {
	b.tag.Attr(":hide-opposite", fmt.Sprint(v))
	return b
}

func (b *VXTimelineBuilder) AnimateOnScroll(v bool) (r *VXTimelineBuilder) {
	b.tag.Attr(":animate-on-scroll", fmt.Sprint(v))
	return b
}

func (b *VXTimelineBuilder) Parallax(v bool) (r *VXTimelineBuilder) {
	b.tag.Attr(":parallax", fmt.Sprint(v))
	return b
}

func (b *VXTimelineBuilder) IconColor(v string) (r *VXTimelineBuilder) {
	b.tag.Attr("icon-color", v)
	return b
}

func (b *VXTimelineBuilder) LineInset(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":line-inset", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) Size(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) Tag(v string) (r *VXTimelineBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXTimelineBuilder) Density(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) Theme(v string) (r *VXTimelineBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXTimelineBuilder) Align(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":align", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) Direction(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) Side(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":side", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) TruncateLine(v interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(":truncate-line", h.JSONString(v))
	return b
}

func (b *VXTimelineBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXTimelineBuilder) Attr(vs ...interface{}) (r *VXTimelineBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTimelineBuilder) Children(children ...h.HTMLComponent) (r *VXTimelineBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXTimelineBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXTimelineBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXTimelineBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXTimelineBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXTimelineBuilder) Class(names ...string) (r *VXTimelineBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXTimelineBuilder) ClassIf(name string, add bool) (r *VXTimelineBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXTimelineBuilder) Sinuous(v bool) (r *VXTimelineBuilder) {
	b.tag.Attr(":sinuous", fmt.Sprint(v))
	return b
}

func (b *VXTimelineBuilder) On(name string, value string) (r *VXTimelineBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXTimelineBuilder) Bind(name string, value string) (r *VXTimelineBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXTimelineBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
