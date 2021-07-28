package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideItem(children ...h.HTMLComponent) (r *VSlideItemBuilder) {
	r = &VSlideItemBuilder{
		tag: h.Tag("v-slide-item").Children(children...),
	}
	return
}

func (b *VSlideItemBuilder) ActiveClass(v string) (r *VSlideItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VSlideItemBuilder) Disabled(v bool) (r *VSlideItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSlideItemBuilder) Value(v interface{}) (r *VSlideItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSlideItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideItemBuilder) Attr(vs ...interface{}) (r *VSlideItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideItemBuilder) Children(children ...h.HTMLComponent) (r *VSlideItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideItemBuilder) Class(names ...string) (r *VSlideItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideItemBuilder) ClassIf(name string, add bool) (r *VSlideItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideItemBuilder) On(name string, value string) (r *VSlideItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideItemBuilder) Bind(name string, value string) (r *VSlideItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
