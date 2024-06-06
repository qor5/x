package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperVerticalItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperVerticalItem(children ...h.HTMLComponent) (r *VStepperVerticalItemBuilder) {
	r = &VStepperVerticalItemBuilder{
		tag: h.Tag("v-stepper-vertical-item").Children(children...),
	}
	return
}

func (b *VStepperVerticalItemBuilder) Icon(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("icon", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Subtitle(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("subtitle", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Title(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Text(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VStepperVerticalItemBuilder) HideActions(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":hide-actions", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Color(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Complete(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":complete", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) CompleteIcon(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("complete-icon", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Editable(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":editable", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) EditIcon(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("edit-icon", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Error(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) ErrorIcon(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("error-icon", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Ripple(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Value(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Rules(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) BgColor(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Elevation(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Disabled(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) SelectedClass(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VStepperVerticalItemBuilder) Rounded(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Tile(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Tag(v string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VStepperVerticalItemBuilder) ExpandIcon(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) CollapseIcon(v interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Focusable(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":focusable", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Static(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":static", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Readonly(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) Eager(v bool) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VStepperVerticalItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperVerticalItemBuilder) Attr(vs ...interface{}) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperVerticalItemBuilder) Children(children ...h.HTMLComponent) (r *VStepperVerticalItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperVerticalItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperVerticalItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperVerticalItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperVerticalItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperVerticalItemBuilder) Class(names ...string) (r *VStepperVerticalItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperVerticalItemBuilder) ClassIf(name string, add bool) (r *VStepperVerticalItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperVerticalItemBuilder) On(name string, value string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperVerticalItemBuilder) Bind(name string, value string) (r *VStepperVerticalItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperVerticalItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
