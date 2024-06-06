package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperWindowItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperWindowItem(children ...h.HTMLComponent) (r *VStepperWindowItemBuilder) {
	r = &VStepperWindowItemBuilder{
		tag: h.Tag("v-stepper-window-item").Children(children...),
	}
	return
}

func (b *VStepperWindowItemBuilder) ReverseTransition(v interface{}) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(":reverse-transition", h.JSONString(v))
	return b
}

func (b *VStepperWindowItemBuilder) Transition(v interface{}) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VStepperWindowItemBuilder) Value(v interface{}) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VStepperWindowItemBuilder) Disabled(v bool) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VStepperWindowItemBuilder) SelectedClass(v string) (r *VStepperWindowItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VStepperWindowItemBuilder) Eager(v bool) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VStepperWindowItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperWindowItemBuilder) Attr(vs ...interface{}) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperWindowItemBuilder) Children(children ...h.HTMLComponent) (r *VStepperWindowItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperWindowItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperWindowItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperWindowItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperWindowItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperWindowItemBuilder) Class(names ...string) (r *VStepperWindowItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperWindowItemBuilder) ClassIf(name string, add bool) (r *VStepperWindowItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperWindowItemBuilder) On(name string, value string) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperWindowItemBuilder) Bind(name string, value string) (r *VStepperWindowItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperWindowItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
