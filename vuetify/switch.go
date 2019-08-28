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
		tag: h.Tag("vw-switch").Children(children...),
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

func (b *VSwitchBuilder) ErrorMessages(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) FalseValue(v string) (r *VSwitchBuilder) {
	b.tag.Attr("false-value", v)
	return b
}

func (b *VSwitchBuilder) Height(v int) (r *VSwitchBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
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

func (b *VSwitchBuilder) InputValue(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":input-value", fmt.Sprint(v))
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

func (b *VSwitchBuilder) Ripple(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Rules(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr("rules", v)
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

func (b *VSwitchBuilder) TrueValue(v string) (r *VSwitchBuilder) {
	b.tag.Attr("true-value", v)
	return b
}

func (b *VSwitchBuilder) ValidateOnBlur(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Value(v bool) (r *VSwitchBuilder) {
	b.tag.Attr("value", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
