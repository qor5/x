package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRadioGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRadioGroup(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	r = &VRadioGroupBuilder{
		tag: h.Tag("v-radio-group").Children(children...),
	}
	return
}

func (b *VRadioGroupBuilder) Label(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRadioGroupBuilder) Height(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Type(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VRadioGroupBuilder) Id(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VRadioGroupBuilder) AppendIcon(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) CenterAffix(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) PrependIcon(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) HideSpinButtons(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Hint(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VRadioGroupBuilder) PersistentHint(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Messages(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Direction(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Density(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) MaxWidth(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) MinWidth(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Width(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Theme(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VRadioGroupBuilder) Disabled(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Error(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) MaxErrors(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Name(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VRadioGroupBuilder) Readonly(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Rules(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) ModelValue(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) ValidateOn(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) ValidationValue(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Focused(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) HideDetails(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Color(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRadioGroupBuilder) DefaultsTarget(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VRadioGroupBuilder) Inline(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) FalseIcon(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) TrueIcon(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Ripple(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) ValueComparator(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VRadioGroupBuilder) Attr(vs ...interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VRadioGroupBuilder) Children(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRadioGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRadioGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRadioGroupBuilder) Class(names ...string) (r *VRadioGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRadioGroupBuilder) ClassIf(name string, add bool) (r *VRadioGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRadioGroupBuilder) On(name string, value string) (r *VRadioGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRadioGroupBuilder) Bind(name string, value string) (r *VRadioGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRadioGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
