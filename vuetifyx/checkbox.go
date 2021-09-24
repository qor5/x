package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXCheckboxBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VXCheckboxBuilder) AppendIcon(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VXCheckboxBuilder) BackgroundColor(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VXCheckboxBuilder) Color(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXCheckboxBuilder) Dark(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Dense(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
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

func (b *VXCheckboxBuilder) ErrorCount(v int) (r *VXCheckboxBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) FalseValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) HideDetails(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Hint(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VXCheckboxBuilder) Id(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXCheckboxBuilder) Indeterminate(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) IndeterminateIcon(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("indeterminate-icon", v)
	return b
}

func (b *VXCheckboxBuilder) InputValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":input-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Label(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXCheckboxBuilder) Light(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Messages(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VXCheckboxBuilder) Multiple(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) OffIcon(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("off-icon", v)
	return b
}

func (b *VXCheckboxBuilder) OnIcon(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("on-icon", v)
	return b
}

func (b *VXCheckboxBuilder) PersistentHint(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) PrependIcon(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VXCheckboxBuilder) Readonly(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Ripple(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Rules(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) Success(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) SuccessMessages(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VXCheckboxBuilder) TrueValue(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) ValidateOnBlur(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VXCheckboxBuilder) Value(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXCheckboxBuilder) ValueComparator(v interface{}) (r *VXCheckboxBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
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

func (b *VXCheckboxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
