package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperVerticalBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperVertical(children ...h.HTMLComponent) (r *VStepperVerticalBuilder) {
	r = &VStepperVerticalBuilder{
		tag: h.Tag("v-stepper-vertical").Children(children...),
	}
	return
}

func (b *VStepperVerticalBuilder) Flat(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) PrevText(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("prev-text", v)
	return b
}

func (b *VStepperVerticalBuilder) NextText(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("next-text", v)
	return b
}

func (b *VStepperVerticalBuilder) AltLabels(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":alt-labels", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) BgColor(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VStepperVerticalBuilder) CompleteIcon(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("complete-icon", v)
	return b
}

func (b *VStepperVerticalBuilder) EditIcon(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("edit-icon", v)
	return b
}

func (b *VStepperVerticalBuilder) Editable(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":editable", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) ErrorIcon(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("error-icon", v)
	return b
}

func (b *VStepperVerticalBuilder) HideActions(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Items(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) ItemTitle(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("item-title", v)
	return b
}

func (b *VStepperVerticalBuilder) ItemValue(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VStepperVerticalBuilder) NonLinear(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":non-linear", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Mobile(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) MobileBreakpoint(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) ModelValue(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) Multiple(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Max(v int) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) SelectedClass(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VStepperVerticalBuilder) Disabled(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Mandatory(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) Elevation(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) Rounded(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) Tile(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Tag(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VStepperVerticalBuilder) Color(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperVerticalBuilder) ExpandIcon(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) CollapseIcon(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) Focusable(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Ripple(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) Readonly(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Eager(v bool) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalBuilder) Theme(v string) (r *VStepperVerticalBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VStepperVerticalBuilder) Variant(v interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VStepperVerticalBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperVerticalBuilder) Attr(vs ...interface{}) (r *VStepperVerticalBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperVerticalBuilder) Children(children ...h.HTMLComponent) (r *VStepperVerticalBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperVerticalBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperVerticalBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperVerticalBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperVerticalBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperVerticalBuilder) Class(names ...string) (r *VStepperVerticalBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperVerticalBuilder) ClassIf(name string, add bool) (r *VStepperVerticalBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperVerticalBuilder) On(name string, value string) (r *VStepperVerticalBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperVerticalBuilder) Bind(name string, value string) (r *VStepperVerticalBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperVerticalBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
