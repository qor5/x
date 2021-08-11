package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePicker(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	r = &VDatePickerBuilder{
		tag: h.Tag("v-date-picker").Children(children...),
	}
	return
}

func (b *VDatePickerBuilder) ActivePicker(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("active-picker", v)
	return b
}

func (b *VDatePickerBuilder) AllowedDates(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":allowed-dates", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Color(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerBuilder) Dark(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) DayFormat(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":day-format", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Disabled(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Elevation(v int) (r *VDatePickerBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) EventColor(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":event-color", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Events(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":events", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) FirstDayOfWeek(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("first-day-of-week", v)
	return b
}

func (b *VDatePickerBuilder) Flat(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) FullWidth(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) HeaderColor(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("header-color", v)
	return b
}

func (b *VDatePickerBuilder) HeaderDateFormat(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":header-date-format", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Landscape(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":landscape", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Light(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Locale(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VDatePickerBuilder) LocaleFirstDayOfYear(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("locale-first-day-of-year", v)
	return b
}

func (b *VDatePickerBuilder) Max(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("max", v)
	return b
}

func (b *VDatePickerBuilder) Min(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("min", v)
	return b
}

func (b *VDatePickerBuilder) MonthFormat(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":month-format", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Multiple(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) NextIcon(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VDatePickerBuilder) NextMonthAriaLabel(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("next-month-aria-label", v)
	return b
}

func (b *VDatePickerBuilder) NextYearAriaLabel(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("next-year-aria-label", v)
	return b
}

func (b *VDatePickerBuilder) NoTitle(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":no-title", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) PickerDate(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("picker-date", v)
	return b
}

func (b *VDatePickerBuilder) PrevIcon(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VDatePickerBuilder) PrevMonthAriaLabel(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("prev-month-aria-label", v)
	return b
}

func (b *VDatePickerBuilder) PrevYearAriaLabel(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("prev-year-aria-label", v)
	return b
}

func (b *VDatePickerBuilder) Range(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":range", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Reactive(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":reactive", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Readonly(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Scrollable(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) SelectedItemsText(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("selected-items-text", v)
	return b
}

func (b *VDatePickerBuilder) ShowAdjacentMonths(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":show-adjacent-months", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) ShowCurrent(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":show-current", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) ShowWeek(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) TitleDateFormat(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":title-date-format", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Type(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VDatePickerBuilder) Value(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("value", v)
	return b
}

func (b *VDatePickerBuilder) WeekdayFormat(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":weekday-format", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Width(v int) (r *VDatePickerBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) YearFormat(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":year-format", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) YearIcon(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("year-icon", v)
	return b
}

func (b *VDatePickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerBuilder) Attr(vs ...interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerBuilder) Class(names ...string) (r *VDatePickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerBuilder) ClassIf(name string, add bool) (r *VDatePickerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerBuilder) On(name string, value string) (r *VDatePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerBuilder) Bind(name string, value string) (r *VDatePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
