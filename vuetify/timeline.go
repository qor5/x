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

func (b *VTimelineBuilder) AlignTop(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":align-top", fmt.Sprint(v))
	return b
}

func (b *VTimelineBuilder) Dark(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTimelineBuilder) Dense(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VTimelineBuilder) Light(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTimelineBuilder) Reverse(v bool) (r *VTimelineBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
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
