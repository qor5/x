package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendar(children ...h.HTMLComponent) (r *VCalendarBuilder) {
	r = &VCalendarBuilder{
		tag: h.Tag("v-calendar").Children(children...),
	}
	return
}

func (b *VCalendarBuilder) Categories(v []string) (r *VCalendarBuilder) {
	b.tag.Attr(":categories", v)
	return b
}

func (b *VCalendarBuilder) CategoryDays(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":category-days", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) CategoryForInvalid(v string) (r *VCalendarBuilder) {
	b.tag.Attr("category-for-invalid", v)
	return b
}

func (b *VCalendarBuilder) CategoryHideDynamic(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":category-hide-dynamic", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) CategoryShowAll(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":category-show-all", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) CategoryText(v string) (r *VCalendarBuilder) {
	b.tag.Attr("category-text", v)
	return b
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

func (b *VCalendarBuilder) EventCategory(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-category", v)
	return b
}

func (b *VCalendarBuilder) EventColor(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-color", v)
	return b
}

func (b *VCalendarBuilder) EventEnd(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-end", v)
	return b
}

func (b *VCalendarBuilder) EventHeight(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":event-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) EventMarginBottom(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":event-margin-bottom", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) EventMore(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":event-more", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) EventMoreText(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-more-text", v)
	return b
}

func (b *VCalendarBuilder) EventName(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-name", v)
	return b
}

func (b *VCalendarBuilder) EventOverlapMode(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-overlap-mode", v)
	return b
}

func (b *VCalendarBuilder) EventOverlapThreshold(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-overlap-threshold", v)
	return b
}

func (b *VCalendarBuilder) EventRipple(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":event-ripple", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) EventStart(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-start", v)
	return b
}

func (b *VCalendarBuilder) EventTextColor(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-text-color", v)
	return b
}

func (b *VCalendarBuilder) EventTimed(v string) (r *VCalendarBuilder) {
	b.tag.Attr("event-timed", v)
	return b
}

func (b *VCalendarBuilder) Events(v []string) (r *VCalendarBuilder) {
	b.tag.Attr(":events", v)
	return b
}

func (b *VCalendarBuilder) FirstInterval(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":first-interval", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) FirstTime(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":first-time", fmt.Sprint(v))
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

func (b *VCalendarBuilder) IntervalWidth(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-width", fmt.Sprint(v))
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

func (b *VCalendarBuilder) LocaleFirstDayOfYear(v string) (r *VCalendarBuilder) {
	b.tag.Attr("locale-first-day-of-year", v)
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

func (b *VCalendarBuilder) ShowWeek(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
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

func (b *VCalendarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarBuilder) Attr(vs ...interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarBuilder) Children(children ...h.HTMLComponent) (r *VCalendarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarBuilder) {
	b.tag.PrependChildren(children...)
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
