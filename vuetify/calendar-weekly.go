package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarWeeklyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarWeekly(children ...h.HTMLComponent) (r *VCalendarWeeklyBuilder) {
	r = &VCalendarWeeklyBuilder{
		tag: h.Tag("v-calendar-weekly").Children(children...),
	}
	return
}

func (b *VCalendarWeeklyBuilder) Color(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCalendarWeeklyBuilder) Dark(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) DayFormat(v interface{}) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":day-format", h.JSONString(v))
	return b
}

func (b *VCalendarWeeklyBuilder) End(v int) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":end", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) HideHeader(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) Light(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) Locale(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VCalendarWeeklyBuilder) LocaleFirstDayOfYear(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("locale-first-day-of-year", v)
	return b
}

func (b *VCalendarWeeklyBuilder) MinWeeks(v interface{}) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":min-weeks", h.JSONString(v))
	return b
}

func (b *VCalendarWeeklyBuilder) MonthFormat(v interface{}) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":month-format", h.JSONString(v))
	return b
}

func (b *VCalendarWeeklyBuilder) Now(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("now", v)
	return b
}

func (b *VCalendarWeeklyBuilder) ShortMonths(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":short-months", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) ShortWeekdays(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":short-weekdays", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) ShowMonthOnFirst(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":show-month-on-first", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) ShowWeek(v bool) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) Start(v int) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":start", fmt.Sprint(v))
	return b
}

func (b *VCalendarWeeklyBuilder) WeekdayFormat(v interface{}) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":weekday-format", h.JSONString(v))
	return b
}

func (b *VCalendarWeeklyBuilder) Weekdays(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("weekdays", v)
	return b
}

func (b *VCalendarWeeklyBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarWeeklyBuilder) Attr(vs ...interface{}) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarWeeklyBuilder) Children(children ...h.HTMLComponent) (r *VCalendarWeeklyBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarWeeklyBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarWeeklyBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarWeeklyBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarWeeklyBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarWeeklyBuilder) Class(names ...string) (r *VCalendarWeeklyBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarWeeklyBuilder) ClassIf(name string, add bool) (r *VCalendarWeeklyBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarWeeklyBuilder) On(name string, value string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarWeeklyBuilder) Bind(name string, value string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarWeeklyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
