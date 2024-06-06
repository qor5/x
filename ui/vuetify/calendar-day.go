package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarDayBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarDay(children ...h.HTMLComponent) (r *VCalendarDayBuilder) {
	r = &VCalendarDayBuilder{
		tag: h.Tag("v-calendar-day").Children(children...),
	}
	return
}

func (b *VCalendarDayBuilder) HideDayHeader(v bool) (r *VCalendarDayBuilder) {
	b.tag.Attr(":hide-day-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) Intervals(v int) (r *VCalendarDayBuilder) {
	b.tag.Attr(":intervals", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) Day(v interface{}) (r *VCalendarDayBuilder) {
	b.tag.Attr(":day", h.JSONString(v))
	return b
}

func (b *VCalendarDayBuilder) DayIndex(v int) (r *VCalendarDayBuilder) {
	b.tag.Attr(":day-index", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) Events(v interface{}) (r *VCalendarDayBuilder) {
	b.tag.Attr(":events", h.JSONString(v))
	return b
}

func (b *VCalendarDayBuilder) IntervalDivisions(v int) (r *VCalendarDayBuilder) {
	b.tag.Attr(":interval-divisions", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) IntervalDuration(v int) (r *VCalendarDayBuilder) {
	b.tag.Attr(":interval-duration", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) IntervalHeight(v int) (r *VCalendarDayBuilder) {
	b.tag.Attr(":interval-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) IntervalFormat(v interface{}) (r *VCalendarDayBuilder) {
	b.tag.Attr(":interval-format", h.JSONString(v))
	return b
}

func (b *VCalendarDayBuilder) IntervalStart(v int) (r *VCalendarDayBuilder) {
	b.tag.Attr(":interval-start", fmt.Sprint(v))
	return b
}

func (b *VCalendarDayBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarDayBuilder) Attr(vs ...interface{}) (r *VCalendarDayBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarDayBuilder) Children(children ...h.HTMLComponent) (r *VCalendarDayBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarDayBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarDayBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarDayBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarDayBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarDayBuilder) Class(names ...string) (r *VCalendarDayBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarDayBuilder) ClassIf(name string, add bool) (r *VCalendarDayBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarDayBuilder) On(name string, value string) (r *VCalendarDayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarDayBuilder) Bind(name string, value string) (r *VCalendarDayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarDayBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
