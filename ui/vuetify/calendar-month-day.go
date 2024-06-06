package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarMonthDayBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarMonthDay(children ...h.HTMLComponent) (r *VCalendarMonthDayBuilder) {
	r = &VCalendarMonthDayBuilder{
		tag: h.Tag("v-calendar-month-day").Children(children...),
	}
	return
}

func (b *VCalendarMonthDayBuilder) Title(v interface{}) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(":title", h.JSONString(v))
	return b
}

func (b *VCalendarMonthDayBuilder) Active(v bool) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthDayBuilder) Color(v string) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCalendarMonthDayBuilder) Day(v interface{}) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(":day", h.JSONString(v))
	return b
}

func (b *VCalendarMonthDayBuilder) Disabled(v bool) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthDayBuilder) Events(v interface{}) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(":events", h.JSONString(v))
	return b
}

func (b *VCalendarMonthDayBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarMonthDayBuilder) Attr(vs ...interface{}) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarMonthDayBuilder) Children(children ...h.HTMLComponent) (r *VCalendarMonthDayBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarMonthDayBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarMonthDayBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarMonthDayBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarMonthDayBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarMonthDayBuilder) Class(names ...string) (r *VCalendarMonthDayBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarMonthDayBuilder) ClassIf(name string, add bool) (r *VCalendarMonthDayBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarMonthDayBuilder) On(name string, value string) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarMonthDayBuilder) Bind(name string, value string) (r *VCalendarMonthDayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarMonthDayBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
