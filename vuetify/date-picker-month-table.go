package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerMonthTableBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerMonthTable() (r *VDatePickerMonthTableBuilder) {
	r = &VDatePickerMonthTableBuilder{
		tag: h.Tag("v-date-picker-month-table"),
	}
	return
}

func (b *VDatePickerMonthTableBuilder) Color(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Current(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("current", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Dark(v bool) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthTableBuilder) Disabled(v bool) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthTableBuilder) EventColor(v []string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":event-color", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Events(v []string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":events", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Light(v bool) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthTableBuilder) Locale(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Max(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("max", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Min(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("min", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Readonly(v bool) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthTableBuilder) Scrollable(v bool) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthTableBuilder) TableDate(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("table-date", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) Value(v string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr("value", v)
	return b
}

func (b *VDatePickerMonthTableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerMonthTableBuilder) Attr(vs ...interface{}) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerMonthTableBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerMonthTableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerMonthTableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerMonthTableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerMonthTableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerMonthTableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerMonthTableBuilder) Class(names ...string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerMonthTableBuilder) ClassIf(name string, add bool) (r *VDatePickerMonthTableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerMonthTableBuilder) On(name string, value string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerMonthTableBuilder) Bind(name string, value string) (r *VDatePickerMonthTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerMonthTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
