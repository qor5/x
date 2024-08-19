package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDateInputBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDateInput(children ...h.HTMLComponent) (r *VDateInputBuilder) {
	r = &VDateInputBuilder{
		tag: h.Tag("v-date-input").Children(children...),
	}
	return
}

func (b *VDateInputBuilder) Flat(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) HideActions(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Focused(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Reverse(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) ModelValue(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Color(v string) (r *VDateInputBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDateInputBuilder) CancelText(v string) (r *VDateInputBuilder) {
	b.tag.Attr("cancel-text", v)
	return b
}

func (b *VDateInputBuilder) Type(v string) (r *VDateInputBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VDateInputBuilder) OkText(v string) (r *VDateInputBuilder) {
	b.tag.Attr("ok-text", v)
	return b
}

func (b *VDateInputBuilder) Autofocus(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Counter(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":counter", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Prefix(v string) (r *VDateInputBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VDateInputBuilder) Placeholder(v string) (r *VDateInputBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VDateInputBuilder) PersistentPlaceholder(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) PersistentCounter(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":persistent-counter", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Suffix(v string) (r *VDateInputBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VDateInputBuilder) Role(v string) (r *VDateInputBuilder) {
	b.tag.Attr("role", v)
	return b
}

func (b *VDateInputBuilder) Text(v string) (r *VDateInputBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VDateInputBuilder) Id(v string) (r *VDateInputBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VDateInputBuilder) AppendIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) CenterAffix(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) PrependIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) HideSpinButtons(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Hint(v string) (r *VDateInputBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VDateInputBuilder) PersistentHint(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Messages(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Direction(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Density(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Height(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) MaxHeight(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) MaxWidth(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) MinHeight(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) MinWidth(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Width(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Theme(v string) (r *VDateInputBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VDateInputBuilder) Disabled(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Error(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) ErrorMessages(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) MaxErrors(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Name(v string) (r *VDateInputBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VDateInputBuilder) Label(v string) (r *VDateInputBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VDateInputBuilder) Readonly(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Rules(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) ValidateOn(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) ValidationValue(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) HideDetails(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) AppendInnerIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":append-inner-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) BgColor(v string) (r *VDateInputBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VDateInputBuilder) Clearable(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) ClearIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Active(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) BaseColor(v string) (r *VDateInputBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VDateInputBuilder) Dirty(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) PersistentClear(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) PrependInnerIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) SingleLine(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Variant(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Loading(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Rounded(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Tile(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) CounterValue(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) ModelModifiers(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":model-modifiers", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Header(v string) (r *VDateInputBuilder) {
	b.tag.Attr("header", v)
	return b
}

func (b *VDateInputBuilder) NextIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) PrevIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) ModeIcon(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":mode-icon", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) ViewMode(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":view-mode", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Month(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":month", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Year(v int) (r *VDateInputBuilder) {
	b.tag.Attr(":year", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) HideWeekdays(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":hide-weekdays", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) ShowWeek(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":show-week", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Transition(v string) (r *VDateInputBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VDateInputBuilder) ReverseTransition(v string) (r *VDateInputBuilder) {
	b.tag.Attr("reverse-transition", v)
	return b
}

func (b *VDateInputBuilder) ShowAdjacentMonths(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":show-adjacent-months", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Weekdays(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":weekdays", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) WeeksInMonth(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":weeks-in-month", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) FirstDayOfWeek(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":first-day-of-week", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) AllowedDates(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":allowed-dates", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) DisplayValue(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":display-value", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Max(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Min(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Multiple(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":multiple", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Landscape(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":landscape", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Title(v string) (r *VDateInputBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VDateInputBuilder) HideHeader(v bool) (r *VDateInputBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VDateInputBuilder) Border(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Elevation(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Location(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Position(v interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VDateInputBuilder) Tag(v string) (r *VDateInputBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VDateInputBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDateInputBuilder) Attr(vs ...interface{}) (r *VDateInputBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDateInputBuilder) Children(children ...h.HTMLComponent) (r *VDateInputBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDateInputBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDateInputBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDateInputBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDateInputBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDateInputBuilder) Class(names ...string) (r *VDateInputBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDateInputBuilder) ClassIf(name string, add bool) (r *VDateInputBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDateInputBuilder) On(name string, value string) (r *VDateInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDateInputBuilder) Bind(name string, value string) (r *VDateInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDateInputBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
