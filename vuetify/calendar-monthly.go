package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarMonthlyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarMonthly() (r *VCalendarMonthlyBuilder) {
	r = &VCalendarMonthlyBuilder{
		tag: h.Tag("v-calendar-monthly"),
	}
	return
}

func (b *VCalendarMonthlyBuilder) Color(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCalendarMonthlyBuilder) Dark(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) End(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("end", v)
	return b
}

func (b *VCalendarMonthlyBuilder) HideHeader(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Light(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Locale(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("locale", v)
	return b
}

func (b *VCalendarMonthlyBuilder) MinWeeks(v interface{}) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":min-weeks", v)
	return b
}

func (b *VCalendarMonthlyBuilder) Now(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("now", v)
	return b
}

func (b *VCalendarMonthlyBuilder) ShortMonths(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":short-months", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) ShortWeekdays(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":short-weekdays", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) ShowMonthOnFirst(v bool) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":show-month-on-first", fmt.Sprint(v))
	return b
}

func (b *VCalendarMonthlyBuilder) Start(v string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr("start", v)
	return b
}

func (b *VCalendarMonthlyBuilder) Weekdays(v []string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(":weekdays", v)
	return b
}

func (b *VCalendarMonthlyBuilder) Class(names ...string) (r *VCalendarMonthlyBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarMonthlyBuilder) ClassIf(name string, add bool) (r *VCalendarMonthlyBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarMonthlyBuilder) On(name string, value string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarMonthlyBuilder) Bind(name string, value string) (r *VCalendarMonthlyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarMonthlyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
