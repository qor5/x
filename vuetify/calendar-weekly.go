package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarWeeklyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarWeekly() (r *VCalendarWeeklyBuilder) {
	r = &VCalendarWeeklyBuilder{
		tag: h.Tag("v-calendar-weekly"),
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

func (b *VCalendarWeeklyBuilder) End(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("end", v)
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

func (b *VCalendarWeeklyBuilder) MinWeeks(v interface{}) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":min-weeks", v)
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

func (b *VCalendarWeeklyBuilder) Start(v string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr("start", v)
	return b
}

func (b *VCalendarWeeklyBuilder) Weekdays(v []string) (r *VCalendarWeeklyBuilder) {
	b.tag.Attr(":weekdays", v)
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
