package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperActionsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperActions(children ...h.HTMLComponent) (r *VStepperActionsBuilder) {
	r = &VStepperActionsBuilder{
		tag: h.Tag("v-stepper-actions").Children(children...),
	}
	return
}

func (b *VStepperActionsBuilder) Color(v string) (r *VStepperActionsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperActionsBuilder) Disabled(v interface{}) (r *VStepperActionsBuilder) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VStepperActionsBuilder) PrevText(v string) (r *VStepperActionsBuilder) {
	b.tag.Attr("prev-text", v)
	return b
}

func (b *VStepperActionsBuilder) NextText(v string) (r *VStepperActionsBuilder) {
	b.tag.Attr("next-text", v)
	return b
}

func (b *VStepperActionsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperActionsBuilder) Attr(vs ...interface{}) (r *VStepperActionsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperActionsBuilder) Children(children ...h.HTMLComponent) (r *VStepperActionsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperActionsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperActionsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperActionsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperActionsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperActionsBuilder) Class(names ...string) (r *VStepperActionsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperActionsBuilder) ClassIf(name string, add bool) (r *VStepperActionsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperActionsBuilder) On(name string, value string) (r *VStepperActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperActionsBuilder) Bind(name string, value string) (r *VStepperActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperActionsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
