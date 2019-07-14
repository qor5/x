package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendar() (r *VCalendarBuilder) {
	r = &VCalendarBuilder{
		tag: h.Tag("v-calendar"),
	}
	return
}

func (b *VCalendarBuilder) Color(v string) (r *VCalendarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCalendarBuilder) Dark(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) End(v string) (r *VCalendarBuilder) {
	b.tag.Attr("end", v)
	return b
}

func (b *VCalendarBuilder) FirstInterval(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":first-interval", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) HideHeader(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) IntervalCount(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-count", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) IntervalHeight(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) IntervalMinutes(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-minutes", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Light(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Locale(v string) (r *VCalendarBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VCalendarBuilder) MaxDays(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":max-days", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) MinWeeks(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":min-weeks", v)
	return b
}

func (b *VCalendarBuilder) Now(v string) (r *VCalendarBuilder) {
	b.tag.Attr("now", v)
	return b
}

func (b *VCalendarBuilder) ShortIntervals(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":short-intervals", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) ShortMonths(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":short-months", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) ShortWeekdays(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":short-weekdays", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) ShowMonthOnFirst(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":show-month-on-first", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Start(v string) (r *VCalendarBuilder) {
	b.tag.Attr("start", v)
	return b
}

func (b *VCalendarBuilder) Type(v string) (r *VCalendarBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VCalendarBuilder) Value(v string) (r *VCalendarBuilder) {
	b.tag.Attr("value", v)
	return b
}

func (b *VCalendarBuilder) Weekdays(v []string) (r *VCalendarBuilder) {
	b.tag.Attr(":weekdays", v)
	return b
}

func (b *VCalendarBuilder) Class(names ...string) (r *VCalendarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarBuilder) ClassIf(name string, add bool) (r *VCalendarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarBuilder) On(name string, value string) (r *VCalendarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarBuilder) Bind(name string, value string) (r *VCalendarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
