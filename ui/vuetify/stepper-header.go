package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperHeaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperHeader(children ...h.HTMLComponent) (r *VStepperHeaderBuilder) {
	r = &VStepperHeaderBuilder{
		tag: h.Tag("v-stepper-header").Children(children...),
	}
	return
}

func (b *VStepperHeaderBuilder) Tag(v string) (r *VStepperHeaderBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VStepperHeaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperHeaderBuilder) Attr(vs ...interface{}) (r *VStepperHeaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperHeaderBuilder) Children(children ...h.HTMLComponent) (r *VStepperHeaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperHeaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperHeaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperHeaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperHeaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperHeaderBuilder) Class(names ...string) (r *VStepperHeaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperHeaderBuilder) ClassIf(name string, add bool) (r *VStepperHeaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperHeaderBuilder) On(name string, value string) (r *VStepperHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperHeaderBuilder) Bind(name string, value string) (r *VStepperHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperHeaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
