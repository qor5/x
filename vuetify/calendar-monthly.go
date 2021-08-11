package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarMonthlyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarMonthly(children ...h.HTMLComponent) (r *VCalendarMonthlyBuilder) {
	r = &VCalendarMonthlyBuilder{
		tag: h.Tag("v-calendar-monthly").Children(children...),
	}
	return
}

func (b *VCalendarMonthlyBuilder) Color(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCalendarMonthlyBuilder) Dark(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) DayFormat(v interface{}) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":day-format", h.JSONString(v))
	return b
}

func (b *VCalendarMonthlyBuilder) End(v int) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":end", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) HideHeader(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Light(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Locale(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VCalendarMonthlyBuilder) LocaleFirstDayOfYear(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("locale-first-day-of-year", v)
	return b
}

func (b *VCalendarMonthlyBuilder) MinWeeks(v interface{}) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":min-weeks", h.JSONString(v))
	return b
}

func (b *VCalendarMonthlyBuilder) MonthFormat(v interface{}) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":month-format", h.JSONString(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Now(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("now", v)
	return b
}

func (b *VCalendarMonthlyBuilder) ShortMonths(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":short-months", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) ShortWeekdays(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":short-weekdays", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) ShowMonthOnFirst(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":show-month-on-first", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) ShowWeek(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Start(v int) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":start", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) WeekdayFormat(v interface{}) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":weekday-format", h.JSONString(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Weekdays(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("weekdays", v)
	return b
}

func (b *VCalendarMonthlyBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarMonthlyBuilder) Attr(vs ...interface{}) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarMonthlyBuilder) Children(children ...h.HTMLComponent) (r *VCalendarMonthlyBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarMonthlyBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarMonthlyBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarMonthlyBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarMonthlyBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarMonthlyBuilder) Class(names ...string) (r *VCalendarMonthlyBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarMonthlyBuilder) ClassIf(name string, add bool) (r *VCalendarMonthlyBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarMonthlyBuilder) On(name string, value string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarMonthlyBuilder) Bind(name string, value string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarMonthlyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
