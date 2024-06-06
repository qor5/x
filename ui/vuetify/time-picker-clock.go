package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTimePickerClockBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTimePickerClock(children ...h.HTMLComponent) (r *VTimePickerClockBuilder) {
	r = &VTimePickerClockBuilder{
		tag: h.Tag("v-time-picker-clock").Children(children...),
	}
	return
}

func (b *VTimePickerClockBuilder) Ampm(v bool) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":ampm", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Color(v string) (r *VTimePickerClockBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTimePickerClockBuilder) Disabled(v bool) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) DisplayedValue(v interface{}) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":displayed-value", h.JSONString(v))
	return b
}

func (b *VTimePickerClockBuilder) Double(v bool) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":double", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Format(v interface{}) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":format", h.JSONString(v))
	return b
}

func (b *VTimePickerClockBuilder) Max(v int) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Min(v int) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":min", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Scrollable(v bool) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Readonly(v bool) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Rotate(v int) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":rotate", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) Step(v int) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":step", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) ModelValue(v int) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VTimePickerClockBuilder) AllowedValues(v interface{}) (r *VTimePickerClockBuilder) {
	b.tag.Attr(":allowed-values", h.JSONString(v))
	return b
}

func (b *VTimePickerClockBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTimePickerClockBuilder) Attr(vs ...interface{}) (r *VTimePickerClockBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTimePickerClockBuilder) Children(children ...h.HTMLComponent) (r *VTimePickerClockBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTimePickerClockBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTimePickerClockBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTimePickerClockBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTimePickerClockBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTimePickerClockBuilder) Class(names ...string) (r *VTimePickerClockBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTimePickerClockBuilder) ClassIf(name string, add bool) (r *VTimePickerClockBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTimePickerClockBuilder) On(name string, value string) (r *VTimePickerClockBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTimePickerClockBuilder) Bind(name string, value string) (r *VTimePickerClockBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTimePickerClockBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
