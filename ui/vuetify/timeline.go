package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTimelineBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTimeline(children ...h.HTMLComponent) (r *VTimelineBuilder) {
	r = &VTimelineBuilder{
		tag: h.Tag("v-timeline").Children(children...),
	}
	return
}

func (b *VTimelineBuilder) Justify(v string) (r *VTimelineBuilder) {
	b.tag.Attr("justify", v)
	return b
}

func (b *VTimelineBuilder) LineThickness(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":line-thickness", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) LineColor(v string) (r *VTimelineBuilder) {
	b.tag.Attr("line-color", v)
	return b
}

func (b *VTimelineBuilder) DotColor(v string) (r *VTimelineBuilder) {
	b.tag.Attr("dot-color", v)
	return b
}

func (b *VTimelineBuilder) FillDot(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":fill-dot", fmt.Sprint(v))
	return b
}

func (b *VTimelineBuilder) HideOpposite(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":hide-opposite", fmt.Sprint(v))
	return b
}

func (b *VTimelineBuilder) IconColor(v string) (r *VTimelineBuilder) {
	b.tag.Attr("icon-color", v)
	return b
}

func (b *VTimelineBuilder) LineInset(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":line-inset", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) Size(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) Tag(v string) (r *VTimelineBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTimelineBuilder) Density(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) Theme(v string) (r *VTimelineBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTimelineBuilder) Align(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":align", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) Direction(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) Side(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":side", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) TruncateLine(v interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(":truncate-line", h.JSONString(v))
	return b
}

func (b *VTimelineBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTimelineBuilder) Attr(vs ...interface{}) (r *VTimelineBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTimelineBuilder) Children(children ...h.HTMLComponent) (r *VTimelineBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTimelineBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTimelineBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTimelineBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTimelineBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTimelineBuilder) Class(names ...string) (r *VTimelineBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTimelineBuilder) ClassIf(name string, add bool) (r *VTimelineBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTimelineBuilder) On(name string, value string) (r *VTimelineBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTimelineBuilder) Bind(name string, value string) (r *VTimelineBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTimelineBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
