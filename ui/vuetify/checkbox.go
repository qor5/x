package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCheckboxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCheckbox(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	r = &VCheckboxBuilder{
		tag: h.Tag("v-checkbox").Children(children...),
	}
	return
}

func (b *VCheckboxBuilder) Label(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VCheckboxBuilder) Id(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VCheckboxBuilder) AppendIcon(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) CenterAffix(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Type(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VCheckboxBuilder) PrependIcon(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) HideSpinButtons(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Hint(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VCheckboxBuilder) PersistentHint(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Messages(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Direction(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Density(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) MaxWidth(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) MinWidth(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Width(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Theme(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VCheckboxBuilder) Disabled(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Error(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) MaxErrors(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Name(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VCheckboxBuilder) Readonly(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Rules(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) ModelValue(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) ValidateOn(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) ValidationValue(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Focused(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) HideDetails(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Indeterminate(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) IndeterminateIcon(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":indeterminate-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) BaseColor(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VCheckboxBuilder) TrueValue(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) FalseValue(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Value(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Color(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCheckboxBuilder) DefaultsTarget(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VCheckboxBuilder) FalseIcon(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) TrueIcon(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Ripple(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Multiple(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) ValueComparator(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCheckboxBuilder) Attr(vs ...interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCheckboxBuilder) Children(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCheckboxBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCheckboxBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCheckboxBuilder) Class(names ...string) (r *VCheckboxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCheckboxBuilder) ClassIf(name string, add bool) (r *VCheckboxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCheckboxBuilder) On(name string, value string) (r *VCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCheckboxBuilder) Bind(name string, value string) (r *VCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCheckboxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
