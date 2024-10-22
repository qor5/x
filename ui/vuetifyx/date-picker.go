package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXDatepickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXDatePicker() (r *VXDatepickerBuilder) {
	r = &VXDatepickerBuilder{
		tag: h.Tag("vx-datepicker"),
	}
	return
}

func (b *VXDatepickerBuilder) Value(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) Label(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":label", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) DialogWidth(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":dialogWidth", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) DateFormat(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":dateFormat", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) TimeFormat(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":timeFormat", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) ClearText(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":clearText", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) OkText(v string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":okText", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) Disabled(v bool) (r *VXDatepickerBuilder) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) ErrorMessages(vs ...string) (r *VXDatepickerBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(vs))
	return b
}

func (b *VXDatepickerBuilder) DatePickerProps(v DatePickerProps) (r *VXDatepickerBuilder) {
	b.tag.Attr(":datePickerProps", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) TimePickerProps(v TimePickerProps) (r *VXDatepickerBuilder) {
	b.tag.Attr(":timePickerProps", h.JSONString(v))
	return b
}

func (b *VXDatepickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXDatepickerBuilder) Attr(vs ...interface{}) (r *VXDatepickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXDatepickerBuilder) HideDetails(v bool) (r *VXDatepickerBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXDatepickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
