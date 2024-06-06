package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCheckboxBtnBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCheckboxBtn(children ...h.HTMLComponent) (r *VCheckboxBtnBuilder) {
	r = &VCheckboxBtnBuilder{
		tag: h.Tag("v-checkbox-btn").Children(children...),
	}
	return
}

func (b *VCheckboxBtnBuilder) Label(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VCheckboxBtnBuilder) Indeterminate(v bool) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBtnBuilder) IndeterminateIcon(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":indeterminate-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Type(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VCheckboxBtnBuilder) BaseColor(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VCheckboxBtnBuilder) TrueValue(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) FalseValue(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Value(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Color(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCheckboxBtnBuilder) Disabled(v bool) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBtnBuilder) DefaultsTarget(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VCheckboxBtnBuilder) Error(v bool) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBtnBuilder) Id(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VCheckboxBtnBuilder) Inline(v bool) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBtnBuilder) FalseIcon(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) TrueIcon(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Ripple(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Multiple(v bool) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBtnBuilder) Name(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VCheckboxBtnBuilder) Readonly(v bool) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBtnBuilder) ModelValue(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) ValueComparator(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Density(v interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VCheckboxBtnBuilder) Theme(v string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VCheckboxBtnBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCheckboxBtnBuilder) Attr(vs ...interface{}) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCheckboxBtnBuilder) Children(children ...h.HTMLComponent) (r *VCheckboxBtnBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCheckboxBtnBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCheckboxBtnBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCheckboxBtnBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCheckboxBtnBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCheckboxBtnBuilder) Class(names ...string) (r *VCheckboxBtnBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCheckboxBtnBuilder) ClassIf(name string, add bool) (r *VCheckboxBtnBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCheckboxBtnBuilder) On(name string, value string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCheckboxBtnBuilder) Bind(name string, value string) (r *VCheckboxBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCheckboxBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
