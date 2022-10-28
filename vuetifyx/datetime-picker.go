package vuetifyx

import (
	"context"
	"fmt"

	"github.com/goplaid/web"
	h "github.com/theplant/htmlgo"
)

type VXDateTimePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXDateTimePicker() (r *VXDateTimePickerBuilder) {
	r = &VXDateTimePickerBuilder{
		tag: h.Tag("vx-datetimepicker"),
	}
	return
}

type DatePickerProps struct {
}

type TimePickerProps struct {
	Format     string `json:"format"` // 可用的选项是 ampm 和 24hr
	Scrollable bool   `json:"scrollable"`
	UseSeconds bool   `json:"use-seconds"`
	NoTitle    bool   `json:"no-title"`
}

func (b *VXDateTimePickerBuilder) Value(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) FieldName(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VXDateTimePickerBuilder) Label(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":label", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) DialogWidth(v int) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":dialogWidth", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) DateFormat(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":dateFormat", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) TimeFormat(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":timeFormat", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) ClearText(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":clearText", h.JSONString(v))
	return b
}
func (b *VXDateTimePickerBuilder) OkText(v string) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":okText", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) Disabled(v bool) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) DatePickerProps(v DatePickerProps) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":datePickerProps", h.JSONString(v))
	return b
}
func (b *VXDateTimePickerBuilder) TimePickerProps(v TimePickerProps) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":timePickerProps", h.JSONString(v))
	return b
}

func (b *VXDateTimePickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXDateTimePickerBuilder) Attr(vs ...interface{}) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXDateTimePickerBuilder) HideDetails(v bool) (r *VXDateTimePickerBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VXDateTimePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
