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

func (b *VCalendarBuilder) HideHeader(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) HideWeekNumber(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":hide-week-number", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Disabled(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Month(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":month", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) ShowAdjacentMonths(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":show-adjacent-months", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Year(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":year", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) Weekdays(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":weekdays", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) WeeksInMonth(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":weeks-in-month", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) FirstDayOfWeek(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":first-day-of-week", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) AllowedDates(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":allowed-dates", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) DisplayValue(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":display-value", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) ModelValue(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) Max(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) Min(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) HideDayHeader(v bool) (r *VCalendarBuilder) {
	b.tag.Attr(":hide-day-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Intervals(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":intervals", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Day(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":day", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) DayIndex(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":day-index", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) Events(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":events", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) IntervalDivisions(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-divisions", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) IntervalDuration(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-duration", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) IntervalHeight(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) IntervalFormat(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-format", h.JSONString(v))
	return b
}

func (b *VCalendarBuilder) IntervalStart(v int) (r *VCalendarBuilder) {
	b.tag.Attr(":interval-start", fmt.Sprint(v))
	return b
}

func (b *VCalendarBuilder) NextIcon(v string) (r *VCalendarBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VCalendarBuilder) PrevIcon(v string) (r *VCalendarBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VCalendarBuilder) Title(v string) (r *VCalendarBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VCalendarBuilder) Text(v string) (r *VCalendarBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VCalendarBuilder) ViewMode(v interface{}) (r *VCalendarBuilder) {
	b.tag.Attr(":view-mode", h.JSONString(v))
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
