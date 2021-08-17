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

func (b *VTextareaBuilder) AppendIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VTextareaBuilder) AppendOuterIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("append-outer-icon", v)
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

func (b *VTextareaBuilder) BackgroundColor(v string) (r *VTextareaBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VTextareaBuilder) ClearIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VTextareaBuilder) Clearable(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Color(v string) (r *VTextareaBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTextareaBuilder) Counter(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) CounterValue(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":counter-value", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Dark(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Dense(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
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

func (b *VTextareaBuilder) ErrorCount(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Filled(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":filled", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Flat(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) FullWidth(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Height(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) HideDetails(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Hint(v string) (r *VTextareaBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VTextareaBuilder) Id(v string) (r *VTextareaBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VTextareaBuilder) Label(v string) (r *VTextareaBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VTextareaBuilder) Light(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) LoaderHeight(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":loader-height", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Loading(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Messages(v string) (r *VTextareaBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VTextareaBuilder) NoResize(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":no-resize", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Outlined(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PersistentHint(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) PersistentPlaceholder(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":persistent-placeholder", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Placeholder(v string) (r *VTextareaBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VTextareaBuilder) Prefix(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VTextareaBuilder) PrependIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VTextareaBuilder) PrependInnerIcon(v string) (r *VTextareaBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VTextareaBuilder) Readonly(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Reverse(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rounded(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) RowHeight(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":row-height", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rows(v int) (r *VTextareaBuilder) {
	b.tag.Attr(":rows", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Rules(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) Shaped(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) SingleLine(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Solo(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) SoloInverted(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Success(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) SuccessMessages(v string) (r *VTextareaBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VTextareaBuilder) Suffix(v string) (r *VTextareaBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VTextareaBuilder) Type(v string) (r *VTextareaBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VTextareaBuilder) ValidateOnBlur(v bool) (r *VTextareaBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VTextareaBuilder) Value(v interface{}) (r *VTextareaBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
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
