package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRadioBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRadio(children ...h.HTMLComponent) (r *VRadioBuilder) {
	r = &VRadioBuilder{
		tag: h.Tag("v-radio").Children(children...),
	}
	return
}

func (b *VRadioBuilder) Label(v string) (r *VRadioBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRadioBuilder) BaseColor(v string) (r *VRadioBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VRadioBuilder) TrueValue(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) FalseValue(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Value(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Color(v string) (r *VRadioBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRadioBuilder) Disabled(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) DefaultsTarget(v string) (r *VRadioBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VRadioBuilder) Error(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Id(v string) (r *VRadioBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VRadioBuilder) Inline(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) FalseIcon(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) TrueIcon(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Ripple(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Multiple(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Name(v string) (r *VRadioBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VRadioBuilder) Readonly(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) ModelValue(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Type(v string) (r *VRadioBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VRadioBuilder) ValueComparator(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Density(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Theme(v string) (r *VRadioBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VRadioBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VRadioBuilder) Attr(vs ...interface{}) (r *VRadioBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VRadioBuilder) Children(children ...h.HTMLComponent) (r *VRadioBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRadioBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRadioBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRadioBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRadioBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRadioBuilder) Class(names ...string) (r *VRadioBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRadioBuilder) ClassIf(name string, add bool) (r *VRadioBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRadioBuilder) On(name string, value string) (r *VRadioBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRadioBuilder) Bind(name string, value string) (r *VRadioBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRadioBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
