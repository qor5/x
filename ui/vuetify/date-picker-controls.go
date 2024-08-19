package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDatePickerControlsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDatePickerControls(children ...h.HTMLComponent) (r *VDatePickerControlsBuilder) {
	r = &VDatePickerControlsBuilder{
		tag: h.Tag("v-date-picker-controls").Children(children...),
	}
	return
}

func (b *VDatePickerControlsBuilder) Active(v interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(":active", h.JSONString(v))
	return b
}

func (b *VDatePickerControlsBuilder) Disabled(v interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VDatePickerControlsBuilder) NextIcon(v interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VDatePickerControlsBuilder) PrevIcon(v interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VDatePickerControlsBuilder) ModeIcon(v interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(":mode-icon", h.JSONString(v))
	return b
}

func (b *VDatePickerControlsBuilder) Text(v string) (r *VDatePickerControlsBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VDatePickerControlsBuilder) ViewMode(v interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(":view-mode", h.JSONString(v))
	return b
}

func (b *VDatePickerControlsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDatePickerControlsBuilder) Attr(vs ...interface{}) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDatePickerControlsBuilder) Children(children ...h.HTMLComponent) (r *VDatePickerControlsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDatePickerControlsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDatePickerControlsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDatePickerControlsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDatePickerControlsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDatePickerControlsBuilder) Class(names ...string) (r *VDatePickerControlsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDatePickerControlsBuilder) ClassIf(name string, add bool) (r *VDatePickerControlsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDatePickerControlsBuilder) On(name string, value string) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDatePickerControlsBuilder) Bind(name string, value string) (r *VDatePickerControlsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDatePickerControlsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
