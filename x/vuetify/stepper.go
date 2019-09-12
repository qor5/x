package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepper(children ...h.HTMLComponent) (r *VStepperBuilder) {
	r = &VStepperBuilder{
		tag: h.Tag("v-stepper").Children(children...),
	}
	return
}

func (b *VStepperBuilder) AltLabels(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":alt-labels", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Dark(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Light(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) NonLinear(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":non-linear", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Value(v int) (r *VStepperBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Vertical(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Class(names ...string) (r *VStepperBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperBuilder) ClassIf(name string, add bool) (r *VStepperBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperBuilder) On(name string, value string) (r *VStepperBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperBuilder) Bind(name string, value string) (r *VStepperBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
