package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerMonthBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerMonth(children ...h.HTMLComponent) (r *VDatePickerMonthBuilder) {
	r = &VDatePickerMonthBuilder{
		tag: h.Tag("v-date-picker-month").Children(children...),
	}
	return
}

func (b *VDatePickerMonthBuilder) Color(v string) (r *VDatePickerMonthBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerMonthBuilder) HideWeekdays(v bool) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":hide-weekdays", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthBuilder) ShowWeek(v bool) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthBuilder) Transition(v string) (r *VDatePickerMonthBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VDatePickerMonthBuilder) ReverseTransition(v string) (r *VDatePickerMonthBuilder) {
	b.tag.Attr("reverse-transition", v)
	return b
}

func (b *VDatePickerMonthBuilder) Disabled(v bool) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthBuilder) Month(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":month", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) ShowAdjacentMonths(v bool) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":show-adjacent-months", fmt.Sprint(v))
	return b
}

func (b *VDatePickerMonthBuilder) Year(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":year", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) Weekdays(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":weekdays", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) WeeksInMonth(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":weeks-in-month", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) FirstDayOfWeek(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":first-day-of-week", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) AllowedDates(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":allowed-dates", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) DisplayValue(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":display-value", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) ModelValue(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) Max(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) Min(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) Multiple(v interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(":multiple", h.JSONString(v))
	return b
}

func (b *VDatePickerMonthBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerMonthBuilder) Attr(vs ...interface{}) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerMonthBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerMonthBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerMonthBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerMonthBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerMonthBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerMonthBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerMonthBuilder) Class(names ...string) (r *VDatePickerMonthBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerMonthBuilder) ClassIf(name string, add bool) (r *VDatePickerMonthBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerMonthBuilder) On(name string, value string) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerMonthBuilder) Bind(name string, value string) (r *VDatePickerMonthBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerMonthBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
