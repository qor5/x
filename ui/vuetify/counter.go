package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCounterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCounter(children ...h.HTMLComponent) (r *VCounterBuilder) {
	r = &VCounterBuilder{
		tag: h.Tag("v-counter").Children(children...),
	}
	return
}

func (b *VCounterBuilder) Active(v bool) (r *VCounterBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VCounterBuilder) Disabled(v bool) (r *VCounterBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCounterBuilder) Max(v interface{}) (r *VCounterBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VCounterBuilder) Value(v interface{}) (r *VCounterBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VCounterBuilder) Transition(v interface{}) (r *VCounterBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VCounterBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCounterBuilder) Attr(vs ...interface{}) (r *VCounterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCounterBuilder) Children(children ...h.HTMLComponent) (r *VCounterBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCounterBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCounterBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCounterBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCounterBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCounterBuilder) Class(names ...string) (r *VCounterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCounterBuilder) ClassIf(name string, add bool) (r *VCounterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCounterBuilder) On(name string, value string) (r *VCounterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCounterBuilder) Bind(name string, value string) (r *VCounterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCounterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
