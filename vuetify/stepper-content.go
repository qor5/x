package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperContentBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperContent(children ...h.HTMLComponent) (r *VStepperContentBuilder) {
	r = &VStepperContentBuilder{
		tag: h.Tag("v-stepper-content").Children(children...),
	}
	return
}

func (b *VStepperContentBuilder) Step(v int) (r *VStepperContentBuilder) {
	b.tag.Attr(":step", fmt.Sprint(v))
	return b
}

func (b *VStepperContentBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperContentBuilder) Attr(vs ...interface{}) (r *VStepperContentBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperContentBuilder) Children(children ...h.HTMLComponent) (r *VStepperContentBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperContentBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperContentBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperContentBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperContentBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperContentBuilder) Class(names ...string) (r *VStepperContentBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperContentBuilder) ClassIf(name string, add bool) (r *VStepperContentBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperContentBuilder) On(name string, value string) (r *VStepperContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperContentBuilder) Bind(name string, value string) (r *VStepperContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperContentBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
