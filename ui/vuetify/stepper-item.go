package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperItem(children ...h.HTMLComponent) (r *VStepperItemBuilder) {
	r = &VStepperItemBuilder{
		tag: h.Tag("v-stepper-item").Children(children...),
	}
	return
}

func (b *VStepperItemBuilder) Icon(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("icon", v)
	return b
}

func (b *VStepperItemBuilder) Title(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VStepperItemBuilder) Subtitle(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("subtitle", v)
	return b
}

func (b *VStepperItemBuilder) Color(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperItemBuilder) Complete(v bool) (r *VStepperItemBuilder) {
	b.tag.Attr(":complete", fmt.Sprint(v))
	return b
}

func (b *VStepperItemBuilder) CompleteIcon(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("complete-icon", v)
	return b
}

func (b *VStepperItemBuilder) Editable(v bool) (r *VStepperItemBuilder) {
	b.tag.Attr(":editable", fmt.Sprint(v))
	return b
}

func (b *VStepperItemBuilder) EditIcon(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("edit-icon", v)
	return b
}

func (b *VStepperItemBuilder) Error(v bool) (r *VStepperItemBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VStepperItemBuilder) ErrorIcon(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("error-icon", v)
	return b
}

func (b *VStepperItemBuilder) Ripple(v interface{}) (r *VStepperItemBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VStepperItemBuilder) Value(v interface{}) (r *VStepperItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VStepperItemBuilder) Rules(v interface{}) (r *VStepperItemBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VStepperItemBuilder) Disabled(v bool) (r *VStepperItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VStepperItemBuilder) SelectedClass(v string) (r *VStepperItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VStepperItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VStepperItemBuilder) Attr(vs ...interface{}) (r *VStepperItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VStepperItemBuilder) Children(children ...h.HTMLComponent) (r *VStepperItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VStepperItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VStepperItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VStepperItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VStepperItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VStepperItemBuilder) Class(names ...string) (r *VStepperItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperItemBuilder) ClassIf(name string, add bool) (r *VStepperItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperItemBuilder) On(name string, value string) (r *VStepperItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperItemBuilder) Bind(name string, value string) (r *VStepperItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
