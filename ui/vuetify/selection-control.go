package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSelectionControlBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSelectionControl(children ...h.HTMLComponent) (r *VSelectionControlBuilder) {
	r = &VSelectionControlBuilder{
		tag: h.Tag("v-selection-control").Children(children...),
	}
	return
}

func (b *VSelectionControlBuilder) Label(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VSelectionControlBuilder) BaseColor(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VSelectionControlBuilder) TrueValue(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":true-value", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) FalseValue(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":false-value", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Value(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Color(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSelectionControlBuilder) Disabled(v bool) (r *VSelectionControlBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlBuilder) DefaultsTarget(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VSelectionControlBuilder) Error(v bool) (r *VSelectionControlBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlBuilder) Id(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSelectionControlBuilder) Inline(v bool) (r *VSelectionControlBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlBuilder) FalseIcon(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) TrueIcon(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Ripple(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Multiple(v bool) (r *VSelectionControlBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlBuilder) Name(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VSelectionControlBuilder) Readonly(v bool) (r *VSelectionControlBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlBuilder) ModelValue(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Type(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VSelectionControlBuilder) ValueComparator(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Density(v interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VSelectionControlBuilder) Theme(v string) (r *VSelectionControlBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSelectionControlBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSelectionControlBuilder) Attr(vs ...interface{}) (r *VSelectionControlBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSelectionControlBuilder) Children(children ...h.HTMLComponent) (r *VSelectionControlBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSelectionControlBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSelectionControlBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSelectionControlBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSelectionControlBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSelectionControlBuilder) Class(names ...string) (r *VSelectionControlBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSelectionControlBuilder) ClassIf(name string, add bool) (r *VSelectionControlBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSelectionControlBuilder) On(name string, value string) (r *VSelectionControlBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSelectionControlBuilder) Bind(name string, value string) (r *VSelectionControlBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSelectionControlBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
