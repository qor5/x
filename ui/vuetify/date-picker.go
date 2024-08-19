package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePicker(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	r = &VDatePickerBuilder{
		tag: h.Tag("v-date-picker").Children(children...),
	}
	return
}

func (b *VDatePickerBuilder) Header(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("header", v)
	return b
}

func (b *VDatePickerBuilder) Title(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VDatePickerBuilder) Active(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":active", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Disabled(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) NextIcon(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) PrevIcon(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) ModeIcon(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":mode-icon", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Text(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VDatePickerBuilder) ViewMode(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":view-mode", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Month(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":month", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Year(v int) (r *VDatePickerBuilder) {
	b.tag.Attr(":year", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Color(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerBuilder) HideWeekdays(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":hide-weekdays", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) ShowWeek(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Transition(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VDatePickerBuilder) ReverseTransition(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("reverse-transition", v)
	return b
}

func (b *VDatePickerBuilder) ShowAdjacentMonths(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":show-adjacent-months", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Weekdays(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":weekdays", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) WeeksInMonth(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":weeks-in-month", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) FirstDayOfWeek(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":first-day-of-week", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) AllowedDates(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":allowed-dates", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) DisplayValue(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":display-value", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) ModelValue(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Max(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Min(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Multiple(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":multiple", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Height(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) BgColor(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VDatePickerBuilder) Landscape(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":landscape", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) HideHeader(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Border(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) MaxHeight(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) MaxWidth(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) MinHeight(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) MinWidth(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Width(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Elevation(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Location(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Position(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Rounded(v interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VDatePickerBuilder) Tile(v bool) (r *VDatePickerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VDatePickerBuilder) Tag(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VDatePickerBuilder) Theme(v string) (r *VDatePickerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDatePickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerBuilder) Attr(vs ...interface{}) (r *VDatePickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerBuilder) Class(names ...string) (r *VDatePickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerBuilder) ClassIf(name string, add bool) (r *VDatePickerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerBuilder) On(name string, value string) (r *VDatePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerBuilder) Bind(name string, value string) (r *VDatePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
