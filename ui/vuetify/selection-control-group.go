package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSelectionControlGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSelectionControlGroup(children ...h.HTMLComponent) (r *VSelectionControlGroupBuilder) {
	r = &VSelectionControlGroupBuilder{
		tag: h.Tag("v-selection-control-group").Children(children...),
	}
	return
}

func (b *VSelectionControlGroupBuilder) Color(v string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSelectionControlGroupBuilder) Disabled(v bool) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlGroupBuilder) DefaultsTarget(v string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr("defaults-target", v)
	return b
}

func (b *VSelectionControlGroupBuilder) Error(v bool) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Id(v string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSelectionControlGroupBuilder) Inline(v bool) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlGroupBuilder) FalseIcon(v interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":false-icon", h.JSONString(v))
	return b
}

func (b *VSelectionControlGroupBuilder) TrueIcon(v interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":true-icon", h.JSONString(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Ripple(v interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Multiple(v bool) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Name(v string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VSelectionControlGroupBuilder) Readonly(v bool) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VSelectionControlGroupBuilder) ModelValue(v interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Type(v string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VSelectionControlGroupBuilder) ValueComparator(v interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":value-comparator", h.JSONString(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Density(v interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VSelectionControlGroupBuilder) Theme(v string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSelectionControlGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSelectionControlGroupBuilder) Attr(vs ...interface{}) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSelectionControlGroupBuilder) Children(children ...h.HTMLComponent) (r *VSelectionControlGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSelectionControlGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSelectionControlGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSelectionControlGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSelectionControlGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSelectionControlGroupBuilder) Class(names ...string) (r *VSelectionControlGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSelectionControlGroupBuilder) ClassIf(name string, add bool) (r *VSelectionControlGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSelectionControlGroupBuilder) On(name string, value string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSelectionControlGroupBuilder) Bind(name string, value string) (r *VSelectionControlGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSelectionControlGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
