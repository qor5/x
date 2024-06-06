package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTimelineItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTimelineItem(children ...h.HTMLComponent) (r *VTimelineItemBuilder) {
	r = &VTimelineItemBuilder{
		tag: h.Tag("v-timeline-item").Children(children...),
	}
	return
}

func (b *VTimelineItemBuilder) Icon(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) DotColor(v string) (r *VTimelineItemBuilder) {
	b.tag.Attr("dot-color", v)
	return b
}

func (b *VTimelineItemBuilder) FillDot(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":fill-dot", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) HideDot(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":hide-dot", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) HideOpposite(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":hide-opposite", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) IconColor(v string) (r *VTimelineItemBuilder) {
	b.tag.Attr("icon-color", v)
	return b
}

func (b *VTimelineItemBuilder) LineInset(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":line-inset", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) Height(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) MaxHeight(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) MaxWidth(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) MinHeight(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) MinWidth(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) Width(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) Elevation(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) Rounded(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) Tile(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) Size(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) Tag(v string) (r *VTimelineItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTimelineItemBuilder) Density(v interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTimelineItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTimelineItemBuilder) Attr(vs ...interface{}) (r *VTimelineItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTimelineItemBuilder) Children(children ...h.HTMLComponent) (r *VTimelineItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTimelineItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTimelineItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTimelineItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTimelineItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTimelineItemBuilder) Class(names ...string) (r *VTimelineItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTimelineItemBuilder) ClassIf(name string, add bool) (r *VTimelineItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTimelineItemBuilder) On(name string, value string) (r *VTimelineItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTimelineItemBuilder) Bind(name string, value string) (r *VTimelineItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTimelineItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
