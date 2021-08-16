package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSwitchBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSwitch(children ...h.HTMLComponent) (r *VSwitchBuilder) {
	r = &VSwitchBuilder{
		tag: h.Tag("v-switch").Children(children...),
	}
	return
}

func (b *VSwitchBuilder) AppendIcon(v string) (r *VSwitchBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VSwitchBuilder) BackgroundColor(v string) (r *VSwitchBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VSwitchBuilder) Color(v string) (r *VSwitchBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSwitchBuilder) Dark(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Dense(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Disabled(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Error(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) ErrorCount(v int) (r *VSwitchBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) FalseValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Flat(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) HideDetails(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Hint(v string) (r *VSwitchBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VSwitchBuilder) Id(v string) (r *VSwitchBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSwitchBuilder) InputValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":input-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Inset(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Label(v string) (r *VSwitchBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSwitchBuilder) Light(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Loading(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Messages(v string) (r *VSwitchBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VSwitchBuilder) Multiple(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) PersistentHint(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) PrependIcon(v string) (r *VSwitchBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VSwitchBuilder) Readonly(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Ripple(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Rules(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Success(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) SuccessMessages(v string) (r *VSwitchBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VSwitchBuilder) TrueValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) ValidateOnBlur(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Value(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) ValueComparator(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSwitchBuilder) Attr(vs ...interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSwitchBuilder) Children(children ...h.HTMLComponent) (r *VSwitchBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSwitchBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSwitchBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSwitchBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSwitchBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSwitchBuilder) Class(names ...string) (r *VSwitchBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSwitchBuilder) ClassIf(name string, add bool) (r *VSwitchBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSwitchBuilder) On(name string, value string) (r *VSwitchBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSwitchBuilder) Bind(name string, value string) (r *VSwitchBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSwitchBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
