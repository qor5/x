package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperItemsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperItems(children ...h.HTMLComponent) (r *VStepperItemsBuilder) {
	r = &VStepperItemsBuilder{
		tag: h.Tag("v-stepper-items").Children(children...),
	}
	return
}

func (b *VStepperItemsBuilder) Class(names ...string) (r *VStepperItemsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperItemsBuilder) ClassIf(name string, add bool) (r *VStepperItemsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperItemsBuilder) On(name string, value string) (r *VStepperItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperItemsBuilder) Bind(name string, value string) (r *VStepperItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperItemsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
