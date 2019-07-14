package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarDailyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarDaily() (r *VCalendarDailyBuilder) {
	r = &VCalendarDailyBuilder{
		tag: h.Tag("v-calendar-daily"),
	}
	return
}

func (b *VCalendarDailyBuilder) Color(v string) (r *VCalendarDailyBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCalendarDailyBuilder) Dark(v bool) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) End(v string) (r *VCalendarDailyBuilder) {
	b.tag.Attr("end", v)
	return b
}

func (b *VCalendarDailyBuilder) FirstInterval(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":first-interval", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) HideHeader(v bool) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) IntervalCount(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":interval-count", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) IntervalHeight(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":interval-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) IntervalMinutes(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":interval-minutes", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) Light(v bool) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) Locale(v string) (r *VCalendarDailyBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VCalendarDailyBuilder) MaxDays(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":max-days", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) Now(v string) (r *VCalendarDailyBuilder) {
	b.tag.Attr("now", v)
	return b
}

func (b *VCalendarDailyBuilder) ShortIntervals(v bool) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":short-intervals", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) ShortWeekdays(v bool) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":short-weekdays", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) Start(v string) (r *VCalendarDailyBuilder) {
	b.tag.Attr("start", v)
	return b
}

func (b *VCalendarDailyBuilder) Weekdays(v []string) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":weekdays", v)
	return b
}

func (b *VCalendarDailyBuilder) Class(names ...string) (r *VCalendarDailyBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarDailyBuilder) ClassIf(name string, add bool) (r *VCalendarDailyBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarDailyBuilder) On(name string, value string) (r *VCalendarDailyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarDailyBuilder) Bind(name string, value string) (r *VCalendarDailyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarDailyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
