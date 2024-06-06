package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VWindowItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VWindowItem(children ...h.HTMLComponent) (r *VWindowItemBuilder) {
	r = &VWindowItemBuilder{
		tag: h.Tag("v-window-item").Children(children...),
	}
	return
}

func (b *VWindowItemBuilder) ReverseTransition(v interface{}) (r *VWindowItemBuilder) {
	b.tag.Attr(":reverse-transition", h.JSONString(v))
	return b
}

func (b *VWindowItemBuilder) Transition(v interface{}) (r *VWindowItemBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VWindowItemBuilder) Value(v interface{}) (r *VWindowItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VWindowItemBuilder) Disabled(v bool) (r *VWindowItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VWindowItemBuilder) SelectedClass(v string) (r *VWindowItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VWindowItemBuilder) Eager(v bool) (r *VWindowItemBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VWindowItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VWindowItemBuilder) Attr(vs ...interface{}) (r *VWindowItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VWindowItemBuilder) Children(children ...h.HTMLComponent) (r *VWindowItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VWindowItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VWindowItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VWindowItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VWindowItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VWindowItemBuilder) Class(names ...string) (r *VWindowItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VWindowItemBuilder) ClassIf(name string, add bool) (r *VWindowItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VWindowItemBuilder) On(name string, value string) (r *VWindowItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VWindowItemBuilder) Bind(name string, value string) (r *VWindowItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VWindowItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
