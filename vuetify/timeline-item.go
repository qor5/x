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

func (b *VTimelineItemBuilder) Color(v string) (r *VTimelineItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTimelineItemBuilder) Dark(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
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

func (b *VTimelineItemBuilder) Icon(v string) (r *VTimelineItemBuilder) {
	b.tag.Attr("icon", v)
	return b
}

func (b *VTimelineItemBuilder) IconColor(v string) (r *VTimelineItemBuilder) {
	b.tag.Attr("icon-color", v)
	return b
}

func (b *VTimelineItemBuilder) Large(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) Left(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) Light(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) Right(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VTimelineItemBuilder) Small(v bool) (r *VTimelineItemBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
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
