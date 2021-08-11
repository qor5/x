package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarDailyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarDaily(children ...h.HTMLComponent) (r *VCalendarDailyBuilder) {
	r = &VCalendarDailyBuilder{
		tag: h.Tag("v-calendar-daily").Children(children...),
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

func (b *VCalendarDailyBuilder) DayFormat(v interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":day-format", h.JSONString(v))
	return b
}

func (b *VCalendarDailyBuilder) End(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":end", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) FirstInterval(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":first-interval", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) FirstTime(v interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":first-time", h.JSONString(v))
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

func (b *VCalendarDailyBuilder) IntervalFormat(v interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":interval-format", h.JSONString(v))
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

func (b *VCalendarDailyBuilder) IntervalStyle(v interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":interval-style", h.JSONString(v))
	return b
}

func (b *VCalendarDailyBuilder) IntervalWidth(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":interval-width", fmt.Sprint(v))
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

func (b *VCalendarDailyBuilder) ShowIntervalLabel(v interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":show-interval-label", h.JSONString(v))
	return b
}

func (b *VCalendarDailyBuilder) Start(v int) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":start", fmt.Sprint(v))
	return b
}

func (b *VCalendarDailyBuilder) WeekdayFormat(v interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(":weekday-format", h.JSONString(v))
	return b
}

func (b *VCalendarDailyBuilder) Weekdays(v string) (r *VCalendarDailyBuilder) {
	b.tag.Attr("weekdays", v)
	return b
}

func (b *VCalendarDailyBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarDailyBuilder) Attr(vs ...interface{}) (r *VCalendarDailyBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarDailyBuilder) Children(children ...h.HTMLComponent) (r *VCalendarDailyBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarDailyBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarDailyBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarDailyBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarDailyBuilder) {
	b.tag.PrependChildren(children...)
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
