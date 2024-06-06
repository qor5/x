package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerHeaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerHeader(children ...h.HTMLComponent) (r *VDatePickerHeaderBuilder) {
	r = &VDatePickerHeaderBuilder{
		tag: h.Tag("v-date-picker-header").Children(children...),
	}
	return
}

func (b *VDatePickerHeaderBuilder) AppendIcon(v string) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VDatePickerHeaderBuilder) Color(v string) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerHeaderBuilder) Header(v string) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr("header", v)
	return b
}

func (b *VDatePickerHeaderBuilder) Transition(v string) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VDatePickerHeaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerHeaderBuilder) Attr(vs ...interface{}) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerHeaderBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerHeaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerHeaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerHeaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerHeaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerHeaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerHeaderBuilder) Class(names ...string) (r *VDatePickerHeaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerHeaderBuilder) ClassIf(name string, add bool) (r *VDatePickerHeaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerHeaderBuilder) On(name string, value string) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerHeaderBuilder) Bind(name string, value string) (r *VDatePickerHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerHeaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
