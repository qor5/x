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

func (b *VStepperBuilder) Color(v string) (r *VStepperBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperBuilder) Dark(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Elevation(v int) (r *VStepperBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Flat(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Height(v int) (r *VStepperBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Light(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) MaxHeight(v int) (r *VStepperBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) MaxWidth(v int) (r *VStepperBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) MinHeight(v int) (r *VStepperBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) MinWidth(v int) (r *VStepperBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) NonLinear(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":non-linear", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Outlined(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Rounded(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Shaped(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Tag(v string) (r *VStepperBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VStepperBuilder) Tile(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Value(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Vertical(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Width(v int) (r *VStepperBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperBuilder) Attr(vs ...interface{}) (r *VStepperBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperBuilder) Children(children ...h.HTMLComponent) (r *VStepperBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperBuilder) {
	b.tag.PrependChildren(children...)
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
