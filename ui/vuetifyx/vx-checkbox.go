package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/x/v3/ui/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXCheckboxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXCheckbox(children ...h.HTMLComponent) (r *VXCheckboxBuilder) {
	r = &VXCheckboxBuilder{
		tag: h.Tag("vx-checkbox").Children(children...),
	}
	return
}

func (b *VXCheckboxBuilder) Title(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VXCheckboxBuilder) Label(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXCheckboxBuilder) Id(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXCheckboxBuilder) AppendIcon(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) CenterAffix(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Type(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXCheckboxBuilder) PrependIcon(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) HideSpinButtons(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Hint(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VXCheckboxBuilder) PersistentHint(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Messages(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Direction(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Density(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) MaxWidth(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) MinWidth(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Width(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Theme(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXCheckboxBuilder) Disabled(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Error(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) MaxErrors(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Name(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VXCheckboxBuilder) Readonly(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Rules(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) ModelValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) ValidateOn(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) ValidationValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Focused(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) HideDetails(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Indeterminate(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) IndeterminateIcon(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":indeterminate-icon", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) BaseColor(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VXCheckboxBuilder) TrueValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) FalseValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Color(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXCheckboxBuilder) DefaultsTarget(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VXCheckboxBuilder) FalseIcon(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) TrueIcon(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Ripple(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Multiple(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) ValueComparator(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) TrueLabel(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr(":true-label", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) FalseLabel(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr(":false-label", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) TrueIconColor(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("true-icon-color", v)
	return b
}

func (b *VXCheckboxBuilder) FalseIconColor(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("false-icon-color", v)
	return b
}

func (b *VXCheckboxBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXCheckboxBuilder) Attr(vs ...interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXCheckboxBuilder) Children(children ...h.HTMLComponent) (r *VXCheckboxBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXCheckboxBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXCheckboxBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXCheckboxBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXCheckboxBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXCheckboxBuilder) Class(names ...string) (r *VXCheckboxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXCheckboxBuilder) ClassIf(name string, add bool) (r *VXCheckboxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXCheckboxBuilder) On(name string, value string) (r *VXCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXCheckboxBuilder) Bind(name string, value string) (r *VXCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXCheckboxBuilder) ErrorMessages(v ...string) (r *VXCheckboxBuilder) {
	vuetify.SetErrorMessages(b.tag, v)
	return b
}

func (b *VXCheckboxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
