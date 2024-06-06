package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperWindowBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperWindow(children ...h.HTMLComponent) (r *VStepperWindowBuilder) {
	r = &VStepperWindowBuilder{
		tag: h.Tag("v-stepper-window").Children(children...),
	}
	return
}

func (b *VStepperWindowBuilder) Reverse(v bool) (r *VStepperWindowBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VStepperWindowBuilder) Direction(v interface{}) (r *VStepperWindowBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VStepperWindowBuilder) ModelValue(v interface{}) (r *VStepperWindowBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VStepperWindowBuilder) Disabled(v bool) (r *VStepperWindowBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VStepperWindowBuilder) SelectedClass(v string) (r *VStepperWindowBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VStepperWindowBuilder) Tag(v string) (r *VStepperWindowBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VStepperWindowBuilder) Theme(v string) (r *VStepperWindowBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VStepperWindowBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperWindowBuilder) Attr(vs ...interface{}) (r *VStepperWindowBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperWindowBuilder) Children(children ...h.HTMLComponent) (r *VStepperWindowBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperWindowBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperWindowBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperWindowBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperWindowBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperWindowBuilder) Class(names ...string) (r *VStepperWindowBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperWindowBuilder) ClassIf(name string, add bool) (r *VStepperWindowBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperWindowBuilder) On(name string, value string) (r *VStepperWindowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperWindowBuilder) Bind(name string, value string) (r *VStepperWindowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperWindowBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
