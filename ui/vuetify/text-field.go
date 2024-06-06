package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTextFieldBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTextField(children ...h.HTMLComponent) (r *VTextFieldBuilder) {
	r = &VTextFieldBuilder{
		tag: h.Tag("v-text-field").Children(children...),
	}
	return
}

func (b *VTextFieldBuilder) Label(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VTextFieldBuilder) Counter(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":counter", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Flat(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Autofocus(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Prefix(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VTextFieldBuilder) Placeholder(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VTextFieldBuilder) PersistentPlaceholder(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) PersistentCounter(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":persistent-counter", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Suffix(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VTextFieldBuilder) Role(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("role", v)
	return b
}

func (b *VTextFieldBuilder) Type(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VTextFieldBuilder) Id(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VTextFieldBuilder) AppendIcon(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) CenterAffix(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) PrependIcon(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) HideSpinButtons(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Hint(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VTextFieldBuilder) PersistentHint(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Messages(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Direction(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Reverse(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Density(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) MaxWidth(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) MinWidth(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Width(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Theme(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTextFieldBuilder) Disabled(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Error(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) MaxErrors(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Name(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VTextFieldBuilder) Readonly(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Rules(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) ModelValue(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) ValidateOn(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) ValidationValue(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Focused(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) HideDetails(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) AppendInnerIcon(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":append-inner-icon", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) BgColor(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VTextFieldBuilder) Clearable(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) ClearIcon(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Active(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Color(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTextFieldBuilder) BaseColor(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VTextFieldBuilder) Dirty(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) PersistentClear(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) PrependInnerIcon(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) SingleLine(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Variant(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Loading(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Rounded(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Tile(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) CounterValue(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) ModelModifiers(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":model-modifiers", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTextFieldBuilder) Attr(vs ...interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTextFieldBuilder) Children(children ...h.HTMLComponent) (r *VTextFieldBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTextFieldBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTextFieldBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTextFieldBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTextFieldBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTextFieldBuilder) Class(names ...string) (r *VTextFieldBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTextFieldBuilder) ClassIf(name string, add bool) (r *VTextFieldBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTextFieldBuilder) On(name string, value string) (r *VTextFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTextFieldBuilder) Bind(name string, value string) (r *VTextFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTextFieldBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
