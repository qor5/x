package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSlideGroupItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSlideGroupItem(children ...h.HTMLComponent) (r *VSlideGroupItemBuilder) {
	r = &VSlideGroupItemBuilder{
		tag: h.Tag("v-slide-group-item").Children(children...),
	}
	return
}

func (b *VSlideGroupItemBuilder) Value(v interface{}) (r *VSlideGroupItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSlideGroupItemBuilder) Disabled(v bool) (r *VSlideGroupItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSlideGroupItemBuilder) SelectedClass(v string) (r *VSlideGroupItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VSlideGroupItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSlideGroupItemBuilder) Attr(vs ...interface{}) (r *VSlideGroupItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSlideGroupItemBuilder) Children(children ...h.HTMLComponent) (r *VSlideGroupItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSlideGroupItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSlideGroupItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSlideGroupItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSlideGroupItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSlideGroupItemBuilder) Class(names ...string) (r *VSlideGroupItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSlideGroupItemBuilder) ClassIf(name string, add bool) (r *VSlideGroupItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSlideGroupItemBuilder) On(name string, value string) (r *VSlideGroupItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSlideGroupItemBuilder) Bind(name string, value string) (r *VSlideGroupItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSlideGroupItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
