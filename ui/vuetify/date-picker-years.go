package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerYearsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerYears(children ...h.HTMLComponent) (r *VDatePickerYearsBuilder) {
	r = &VDatePickerYearsBuilder{
		tag: h.Tag("v-date-picker-years").Children(children...),
	}
	return
}

func (b *VDatePickerYearsBuilder) Color(v string) (r *VDatePickerYearsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VDatePickerYearsBuilder) Height(v interface{}) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VDatePickerYearsBuilder) ModelValue(v int) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VDatePickerYearsBuilder) Min(v interface{}) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(":min", h.JSONString(v))
	return b
}

func (b *VDatePickerYearsBuilder) Max(v interface{}) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VDatePickerYearsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerYearsBuilder) Attr(vs ...interface{}) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerYearsBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerYearsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerYearsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerYearsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerYearsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerYearsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerYearsBuilder) Class(names ...string) (r *VDatePickerYearsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerYearsBuilder) ClassIf(name string, add bool) (r *VDatePickerYearsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerYearsBuilder) On(name string, value string) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerYearsBuilder) Bind(name string, value string) (r *VDatePickerYearsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerYearsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
