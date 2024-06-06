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

func (b *VSwitchBuilder) Label(v string) (r *VSwitchBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSwitchBuilder) Indeterminate(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Inset(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Flat(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Loading(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Type(v string) (r *VSwitchBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VSwitchBuilder) Id(v string) (r *VSwitchBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSwitchBuilder) AppendIcon(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) CenterAffix(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) PrependIcon(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) HideSpinButtons(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Hint(v string) (r *VSwitchBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VSwitchBuilder) PersistentHint(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Messages(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Direction(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Density(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) MaxWidth(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) MinWidth(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Width(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Theme(v string) (r *VSwitchBuilder) {
	b.tag.Attr("theme", v)
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

func (b *VSwitchBuilder) MaxErrors(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Name(v string) (r *VSwitchBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VSwitchBuilder) Readonly(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) Rules(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) ModelValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) ValidateOn(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) ValidationValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Focused(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) HideDetails(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) BaseColor(v string) (r *VSwitchBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VSwitchBuilder) TrueValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) FalseValue(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Value(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Color(v string) (r *VSwitchBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSwitchBuilder) DefaultsTarget(v string) (r *VSwitchBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VSwitchBuilder) Inline(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VSwitchBuilder) FalseIcon(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) TrueIcon(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Ripple(v interface{}) (r *VSwitchBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VSwitchBuilder) Multiple(v bool) (r *VSwitchBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
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
