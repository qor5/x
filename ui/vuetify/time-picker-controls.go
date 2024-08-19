package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTimePickerControlsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTimePickerControls(children ...h.HTMLComponent) (r *VTimePickerControlsBuilder) {
	r = &VTimePickerControlsBuilder{
		tag: h.Tag("v-time-picker-controls").Children(children...),
	}
	return
}

func (b *VTimePickerControlsBuilder) Ampm(v bool) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":ampm", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) AmpmInTitle(v bool) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":ampm-in-title", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) AmpmReadonly(v bool) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":ampm-readonly", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Color(v string) (r *VTimePickerControlsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTimePickerControlsBuilder) Disabled(v bool) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Hour(v int) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":hour", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Minute(v int) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":minute", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Second(v int) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":second", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Period(v string) (r *VTimePickerControlsBuilder) {
	b.tag.Attr("period", v)
	return b
}

func (b *VTimePickerControlsBuilder) Readonly(v bool) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) UseSeconds(v bool) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":use-seconds", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Selecting(v int) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":selecting", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) Value(v int) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VTimePickerControlsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTimePickerControlsBuilder) Attr(vs ...interface{}) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTimePickerControlsBuilder) Children(children ...h.HTMLComponent) (r *VTimePickerControlsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTimePickerControlsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTimePickerControlsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTimePickerControlsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTimePickerControlsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTimePickerControlsBuilder) Class(names ...string) (r *VTimePickerControlsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTimePickerControlsBuilder) ClassIf(name string, add bool) (r *VTimePickerControlsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTimePickerControlsBuilder) On(name string, value string) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTimePickerControlsBuilder) Bind(name string, value string) (r *VTimePickerControlsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTimePickerControlsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
