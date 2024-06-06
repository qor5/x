package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerMonthsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerMonths(children ...h.HTMLComponent) (r *VDatePickerMonthsBuilder) {
	r = &VDatePickerMonthsBuilder{
		tag: h.Tag("v-date-picker-months").Children(children...),
	}
	return
}

func (b *VDatePickerMonthsBuilder) Color(v string) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerMonthsBuilder) Height(v interface{}) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthsBuilder) ModelValue(v int) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthsBuilder) Year(v int) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(":year", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthsBuilder) Min(v interface{}) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthsBuilder) Max(v interface{}) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerMonthsBuilder) Attr(vs ...interface{}) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerMonthsBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerMonthsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerMonthsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerMonthsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerMonthsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerMonthsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerMonthsBuilder) Class(names ...string) (r *VDatePickerMonthsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerMonthsBuilder) ClassIf(name string, add bool) (r *VDatePickerMonthsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerMonthsBuilder) On(name string, value string) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerMonthsBuilder) Bind(name string, value string) (r *VDatePickerMonthsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerMonthsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
