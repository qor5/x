package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepper(children ...h.HTMLComponent) (r *VStepperBuilder) {
	r = &VStepperBuilder{
		tag: h.Tag("v-stepper").Children(children...),
	}
	return
}

func (b *VStepperBuilder) Flat(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) AltLabels(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":alt-labels", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) BgColor(v string) (r *VStepperBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VStepperBuilder) CompleteIcon(v string) (r *VStepperBuilder) {
	b.tag.Attr("complete-icon", v)
	return b
}

func (b *VStepperBuilder) EditIcon(v string) (r *VStepperBuilder) {
	b.tag.Attr("edit-icon", v)
	return b
}

func (b *VStepperBuilder) Editable(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":editable", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) ErrorIcon(v string) (r *VStepperBuilder) {
	b.tag.Attr("error-icon", v)
	return b
}

func (b *VStepperBuilder) HideActions(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Items(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) ItemTitle(v string) (r *VStepperBuilder) {
	b.tag.Attr("item-title", v)
	return b
}

func (b *VStepperBuilder) ItemValue(v string) (r *VStepperBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VStepperBuilder) NonLinear(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":non-linear", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Mobile(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) MobileBreakpoint(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) ModelValue(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Multiple(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Max(v int) (r *VStepperBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) SelectedClass(v string) (r *VStepperBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VStepperBuilder) Disabled(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Mandatory(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Color(v string) (r *VStepperBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperBuilder) Border(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Height(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) MaxHeight(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) MaxWidth(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) MinHeight(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) MinWidth(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Width(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Elevation(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Location(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Position(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Rounded(v interface{}) (r *VStepperBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VStepperBuilder) Tile(v bool) (r *VStepperBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VStepperBuilder) Tag(v string) (r *VStepperBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VStepperBuilder) Theme(v string) (r *VStepperBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VStepperBuilder) PrevText(v string) (r *VStepperBuilder) {
	b.tag.Attr("prev-text", v)
	return b
}

func (b *VStepperBuilder) NextText(v string) (r *VStepperBuilder) {
	b.tag.Attr("next-text", v)
	return b
}

func (b *VStepperBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperBuilder) Attr(vs ...interface{}) (r *VStepperBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperBuilder) Children(children ...h.HTMLComponent) (r *VStepperBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperBuilder) Class(names ...string) (r *VStepperBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperBuilder) ClassIf(name string, add bool) (r *VStepperBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperBuilder) On(name string, value string) (r *VStepperBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperBuilder) Bind(name string, value string) (r *VStepperBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
