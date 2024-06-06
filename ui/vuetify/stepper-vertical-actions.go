package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperVerticalActionsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperVerticalActions(children ...h.HTMLComponent) (r *VStepperVerticalActionsBuilder) {
	r = &VStepperVerticalActionsBuilder{
		tag: h.Tag("v-stepper-vertical-actions").Children(children...),
	}
	return
}

func (b *VStepperVerticalActionsBuilder) Color(v string) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperVerticalActionsBuilder) Disabled(v interface{}) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VStepperVerticalActionsBuilder) PrevText(v string) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr("prev-text", v)
	return b
}

func (b *VStepperVerticalActionsBuilder) NextText(v string) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr("next-text", v)
	return b
}

func (b *VStepperVerticalActionsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperVerticalActionsBuilder) Attr(vs ...interface{}) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperVerticalActionsBuilder) Children(children ...h.HTMLComponent) (r *VStepperVerticalActionsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperVerticalActionsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperVerticalActionsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperVerticalActionsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperVerticalActionsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperVerticalActionsBuilder) Class(names ...string) (r *VStepperVerticalActionsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperVerticalActionsBuilder) ClassIf(name string, add bool) (r *VStepperVerticalActionsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperVerticalActionsBuilder) On(name string, value string) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperVerticalActionsBuilder) Bind(name string, value string) (r *VStepperVerticalActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperVerticalActionsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
