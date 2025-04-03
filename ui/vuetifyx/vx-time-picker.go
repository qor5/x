package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXTimePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXTimepicker(children ...h.HTMLComponent) (r *VXTimePickerBuilder) {
	r = &VXTimePickerBuilder{
		tag: h.Tag("vx-time-picker").Children(children...),
	}
	return
}

func (b *VXTimePickerBuilder) Label(v string) (r *VXTimePickerBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXTimePickerBuilder) Name(v string) (r *VXTimePickerBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VXTimePickerBuilder) Id(v string) (r *VXTimePickerBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXTimePickerBuilder) Format(v string) (r *VXTimePickerBuilder) {
	b.tag.Attr("format", v)
	return b
}

func (b *VXTimePickerBuilder) Placeholder(v string) (r *VXTimePickerBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VXTimePickerBuilder) Width(v int) (r *VXTimePickerBuilder) {
	b.tag.Attr("width", h.JSONString(v))
	return b
}

func (b *VXTimePickerBuilder) Disabled(v bool) (r *VXTimePickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) Required(v bool) (r *VXTimePickerBuilder) {
	b.tag.Attr(":required", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) HideDetails(v bool) (r *VXTimePickerBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) Clearable(v bool) (r *VXTimePickerBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) TimePickerProps(v interface{}) (r *VXTimePickerBuilder) {
	b.tag.Attr(":time-picker-props", h.JSONString(v))
	return b
}

func (b *VXTimePickerBuilder) ModelValue(v interface{}) (r *VXTimePickerBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXTimePickerBuilder) HideAppendInner(v bool) (r *VXTimePickerBuilder) {
	b.tag.Attr(":hide-append-inner", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) NeedConfirm(v bool) (r *VXTimePickerBuilder) {
	b.tag.Attr(":need-confirm", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) Attr(vs ...interface{}) (r *VXTimePickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTimePickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXTimePickerBuilder) Children(children ...h.HTMLComponent) (r *VXTimePickerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXTimePickerBuilder) Class(names ...string) (r *VXTimePickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXTimePickerBuilder) Tips(v string) (r *VXTimePickerBuilder) {
	b.tag.Attr("tips", fmt.Sprint(v))
	return b
}

func (b *VXTimePickerBuilder) On(name string, value string) (r *VXTimePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXTimePickerBuilder) Bind(name string, value string) (r *VXTimePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXTimePickerBuilder) ErrorMessages(errMsgs ...string) (r *VXTimePickerBuilder) {
	b.tag.Attr(":error-messages", errMsgs)
	return b
}

func (b *VXTimePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
