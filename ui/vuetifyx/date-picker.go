package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXDatePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXDatePicker() (r *VXDatePickerBuilder) {
	r = &VXDatePickerBuilder{
		tag: h.Tag("vx-datepicker"),
	}
	return
}

func (b *VXDatePickerBuilder) Value(v string) (r *VXDatePickerBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) Label(v string) (r *VXDatePickerBuilder) {
	b.tag.Attr(":label", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) DialogWidth(v int) (r *VXDatePickerBuilder) {
	b.tag.Attr(":dialogWidth", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) DateFormat(v string) (r *VXDatePickerBuilder) {
	b.tag.Attr(":dateFormat", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) TimeFormat(v string) (r *VXDatePickerBuilder) {
	b.tag.Attr(":timeFormat", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) ClearText(v string) (r *VXDatePickerBuilder) {
	b.tag.Attr(":clearText", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) OkText(v string) (r *VXDatePickerBuilder) {
	b.tag.Attr(":okText", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) Disabled(v bool) (r *VXDatePickerBuilder) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) DatePickerProps(v DatePickerProps) (r *VXDatePickerBuilder) {
	b.tag.Attr(":datePickerProps", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) TimePickerProps(v TimePickerProps) (r *VXDatePickerBuilder) {
	b.tag.Attr(":timePickerProps", h.JSONString(v))
	return b
}

func (b *VXDatePickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXDatePickerBuilder) Attr(vs ...interface{}) (r *VXDatePickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXDatePickerBuilder) HideDetails(v bool) (r *VXDatePickerBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXDatePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
