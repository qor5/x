package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarIntervalEventBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarIntervalEvent(children ...h.HTMLComponent) (r *VCalendarIntervalEventBuilder) {
	r = &VCalendarIntervalEventBuilder{
		tag: h.Tag("v-calendar-interval-event").Children(children...),
	}
	return
}

func (b *VCalendarIntervalEventBuilder) AllDay(v bool) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(":all-day", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalEventBuilder) Interval(v interface{}) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(":interval", h.JSONString(v))
	return b
}

func (b *VCalendarIntervalEventBuilder) IntervalDivisions(v int) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(":interval-divisions", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalEventBuilder) IntervalDuration(v int) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(":interval-duration", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalEventBuilder) IntervalHeight(v int) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(":interval-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalEventBuilder) Event(v interface{}) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(":event", h.JSONString(v))
	return b
}

func (b *VCalendarIntervalEventBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarIntervalEventBuilder) Attr(vs ...interface{}) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarIntervalEventBuilder) Children(children ...h.HTMLComponent) (r *VCalendarIntervalEventBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarIntervalEventBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarIntervalEventBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarIntervalEventBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarIntervalEventBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarIntervalEventBuilder) Class(names ...string) (r *VCalendarIntervalEventBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarIntervalEventBuilder) ClassIf(name string, add bool) (r *VCalendarIntervalEventBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarIntervalEventBuilder) On(name string, value string) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarIntervalEventBuilder) Bind(name string, value string) (r *VCalendarIntervalEventBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarIntervalEventBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
