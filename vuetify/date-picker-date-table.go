package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerDateTableBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerDateTable() (r *VDatePickerDateTableBuilder) {
	r = &VDatePickerDateTableBuilder{
		tag: h.Tag("v-date-picker-date-table"),
	}
	return
}

func (b *VDatePickerDateTableBuilder) Color(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Current(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("current", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Dark(v bool) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDatePickerDateTableBuilder) Disabled(v bool) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDatePickerDateTableBuilder) EventColor(v []string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":event-color", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Events(v []string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":events", v)
	return b
}

func (b *VDatePickerDateTableBuilder) FirstDayOfWeek(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("first-day-of-week", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Light(v bool) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDatePickerDateTableBuilder) Locale(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Max(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("max", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Min(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("min", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Readonly(v bool) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VDatePickerDateTableBuilder) Scrollable(v bool) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VDatePickerDateTableBuilder) ShowWeek(v bool) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VDatePickerDateTableBuilder) TableDate(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("table-date", v)
	return b
}

func (b *VDatePickerDateTableBuilder) Value(v string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr("value", v)
	return b
}

func (b *VDatePickerDateTableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerDateTableBuilder) Attr(vs ...interface{}) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerDateTableBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerDateTableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerDateTableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerDateTableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerDateTableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerDateTableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerDateTableBuilder) Class(names ...string) (r *VDatePickerDateTableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerDateTableBuilder) ClassIf(name string, add bool) (r *VDatePickerDateTableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerDateTableBuilder) On(name string, value string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerDateTableBuilder) Bind(name string, value string) (r *VDatePickerDateTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerDateTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
