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

func (b *VTextFieldBuilder) AppendIcon(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VTextFieldBuilder) AppendOuterIcon(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VTextFieldBuilder) Autofocus(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) BackgroundColor(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VTextFieldBuilder) ClearIcon(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VTextFieldBuilder) Clearable(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Color(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTextFieldBuilder) Counter(v int) (r *VTextFieldBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) CounterValue(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Dark(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Dense(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
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

func (b *VTextFieldBuilder) ErrorCount(v int) (r *VTextFieldBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Filled(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Flat(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) FullWidth(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Height(v int) (r *VTextFieldBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) HideDetails(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Hint(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VTextFieldBuilder) Id(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VTextFieldBuilder) Label(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VTextFieldBuilder) Light(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) LoaderHeight(v int) (r *VTextFieldBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Loading(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Messages(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VTextFieldBuilder) Outlined(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) PersistentHint(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) PersistentPlaceholder(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Placeholder(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VTextFieldBuilder) Prefix(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VTextFieldBuilder) PrependIcon(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VTextFieldBuilder) PrependInnerIcon(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VTextFieldBuilder) Readonly(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Reverse(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Rounded(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Rules(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) Shaped(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) SingleLine(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Solo(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) SoloInverted(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Success(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) SuccessMessages(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VTextFieldBuilder) Suffix(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VTextFieldBuilder) Type(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VTextFieldBuilder) ValidateOnBlur(v bool) (r *VTextFieldBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VTextFieldBuilder) Value(v interface{}) (r *VTextFieldBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
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
