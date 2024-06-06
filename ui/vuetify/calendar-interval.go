package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarIntervalBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarInterval(children ...h.HTMLComponent) (r *VCalendarIntervalBuilder) {
	r = &VCalendarIntervalBuilder{
		tag: h.Tag("v-calendar-interval").Children(children...),
	}
	return
}

func (b *VCalendarIntervalBuilder) Index(v int) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":index", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalBuilder) Day(v interface{}) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":day", h.JSONString(v))
	return b
}

func (b *VCalendarIntervalBuilder) DayIndex(v int) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":day-index", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalBuilder) Events(v interface{}) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":events", h.JSONString(v))
	return b
}

func (b *VCalendarIntervalBuilder) IntervalDivisions(v int) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":interval-divisions", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalBuilder) IntervalDuration(v int) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":interval-duration", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalBuilder) IntervalHeight(v int) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":interval-height", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalBuilder) IntervalFormat(v interface{}) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":interval-format", h.JSONString(v))
	return b
}

func (b *VCalendarIntervalBuilder) IntervalStart(v int) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(":interval-start", fmt.Sprint(v))
	return b
}

func (b *VCalendarIntervalBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarIntervalBuilder) Attr(vs ...interface{}) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarIntervalBuilder) Children(children ...h.HTMLComponent) (r *VCalendarIntervalBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarIntervalBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarIntervalBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarIntervalBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarIntervalBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarIntervalBuilder) Class(names ...string) (r *VCalendarIntervalBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarIntervalBuilder) ClassIf(name string, add bool) (r *VCalendarIntervalBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarIntervalBuilder) On(name string, value string) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarIntervalBuilder) Bind(name string, value string) (r *VCalendarIntervalBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarIntervalBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
