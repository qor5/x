package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VNumberInputBuilder struct {
	tag *h.HTMLTagBuilder
}

func VNumberInput(children ...h.HTMLComponent) (r *VNumberInputBuilder) {
	r = &VNumberInputBuilder{
		tag: h.Tag("v-number-input").Children(children...),
	}
	return
}

func (b *VNumberInputBuilder) Label(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VNumberInputBuilder) Counter(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":counter", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Flat(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) ControlVariant(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":control-variant", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Inset(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) HideInput(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":hide-input", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) ModelValue(v int) (r *VNumberInputBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Min(v int) (r *VNumberInputBuilder) {
	b.tag.Attr(":min", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Type(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VNumberInputBuilder) Max(v int) (r *VNumberInputBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Step(v int) (r *VNumberInputBuilder) {
	b.tag.Attr(":step", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Autofocus(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Prefix(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VNumberInputBuilder) Placeholder(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VNumberInputBuilder) PersistentPlaceholder(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) PersistentCounter(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":persistent-counter", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Suffix(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VNumberInputBuilder) Role(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("role", v)
	return b
}

func (b *VNumberInputBuilder) Id(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VNumberInputBuilder) AppendIcon(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) CenterAffix(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) PrependIcon(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) HideSpinButtons(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Hint(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VNumberInputBuilder) PersistentHint(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Messages(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Direction(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Reverse(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Density(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) MaxWidth(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) MinWidth(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Width(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Theme(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VNumberInputBuilder) Disabled(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Error(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) ErrorMessages(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) MaxErrors(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Name(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VNumberInputBuilder) Readonly(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Rules(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) ValidateOn(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) ValidationValue(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Focused(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) HideDetails(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) BgColor(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VNumberInputBuilder) Clearable(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) ClearIcon(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Active(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Color(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VNumberInputBuilder) BaseColor(v string) (r *VNumberInputBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VNumberInputBuilder) Dirty(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) PersistentClear(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) SingleLine(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) Variant(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Loading(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Rounded(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) Tile(v bool) (r *VNumberInputBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VNumberInputBuilder) CounterValue(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) ModelModifiers(v interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(":model-modifiers", h.JSONString(v))
	return b
}

func (b *VNumberInputBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VNumberInputBuilder) Attr(vs ...interface{}) (r *VNumberInputBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VNumberInputBuilder) Children(children ...h.HTMLComponent) (r *VNumberInputBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VNumberInputBuilder) AppendChildren(children ...h.HTMLComponent) (r *VNumberInputBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VNumberInputBuilder) PrependChildren(children ...h.HTMLComponent) (r *VNumberInputBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VNumberInputBuilder) Class(names ...string) (r *VNumberInputBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VNumberInputBuilder) ClassIf(name string, add bool) (r *VNumberInputBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VNumberInputBuilder) On(name string, value string) (r *VNumberInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VNumberInputBuilder) Bind(name string, value string) (r *VNumberInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VNumberInputBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
