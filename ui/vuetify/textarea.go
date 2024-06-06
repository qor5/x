package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTextareaBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTextarea(children ...h.HTMLComponent) (r *VTextareaBuilder) {
	r = &VTextareaBuilder{
		tag: h.Tag("v-textarea").Children(children...),
	}
	return
}

func (b *VTextareaBuilder) Label(v string) (r *VTextareaBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VTextareaBuilder) Counter(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":counter", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Flat(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) AutoGrow(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":auto-grow", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Autofocus(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Prefix(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VTextareaBuilder) Placeholder(v string) (r *VTextareaBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VTextareaBuilder) PersistentPlaceholder(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PersistentCounter(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-counter", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) NoResize(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":no-resize", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rows(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":rows", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) MaxRows(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":max-rows", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Suffix(v string) (r *VTextareaBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VTextareaBuilder) Id(v string) (r *VTextareaBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VTextareaBuilder) AppendIcon(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) CenterAffix(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PrependIcon(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) HideSpinButtons(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Hint(v string) (r *VTextareaBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VTextareaBuilder) PersistentHint(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Messages(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Direction(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Reverse(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Density(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) MaxWidth(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) MinWidth(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Width(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Theme(v string) (r *VTextareaBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTextareaBuilder) Disabled(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Error(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) MaxErrors(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Name(v string) (r *VTextareaBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VTextareaBuilder) Readonly(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rules(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) ModelValue(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) ValidateOn(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) ValidationValue(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Focused(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) HideDetails(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) AppendInnerIcon(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":append-inner-icon", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) BgColor(v string) (r *VTextareaBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VTextareaBuilder) Clearable(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) ClearIcon(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Active(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Color(v string) (r *VTextareaBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTextareaBuilder) BaseColor(v string) (r *VTextareaBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VTextareaBuilder) Dirty(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PersistentClear(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PrependInnerIcon(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) SingleLine(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Variant(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Loading(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Rounded(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Tile(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) CounterValue(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) ModelModifiers(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":model-modifiers", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTextareaBuilder) Attr(vs ...interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTextareaBuilder) Children(children ...h.HTMLComponent) (r *VTextareaBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTextareaBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTextareaBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTextareaBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTextareaBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTextareaBuilder) Class(names ...string) (r *VTextareaBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTextareaBuilder) ClassIf(name string, add bool) (r *VTextareaBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTextareaBuilder) On(name string, value string) (r *VTextareaBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTextareaBuilder) Bind(name string, value string) (r *VTextareaBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTextareaBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
