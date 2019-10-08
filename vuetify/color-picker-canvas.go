package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VColorPickerCanvasBuilder struct {
	tag *h.HTMLTagBuilder
}

func VColorPickerCanvas() (r *VColorPickerCanvasBuilder) {
	r = &VColorPickerCanvasBuilder{
		tag: h.Tag("v-color-picker-canvas"),
	}
	return
}

func (b *VColorPickerCanvasBuilder) Color(v interface{}) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(":color", v)
	return b
}

func (b *VColorPickerCanvasBuilder) Disabled(v bool) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VColorPickerCanvasBuilder) DotSize(v int) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(":dot-size", fmt.Sprint(v))
	return b
}

func (b *VColorPickerCanvasBuilder) Height(v int) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VColorPickerCanvasBuilder) Width(v int) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VColorPickerCanvasBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VColorPickerCanvasBuilder) Attr(vs ...interface{}) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VColorPickerCanvasBuilder) Children(children ...h.HTMLComponent) (r *VColorPickerCanvasBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VColorPickerCanvasBuilder) AppendChildren(children ...h.HTMLComponent) (r *VColorPickerCanvasBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VColorPickerCanvasBuilder) PrependChildren(children ...h.HTMLComponent) (r *VColorPickerCanvasBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VColorPickerCanvasBuilder) Class(names ...string) (r *VColorPickerCanvasBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VColorPickerCanvasBuilder) ClassIf(name string, add bool) (r *VColorPickerCanvasBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VColorPickerCanvasBuilder) On(name string, value string) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VColorPickerCanvasBuilder) Bind(name string, value string) (r *VColorPickerCanvasBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VColorPickerCanvasBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
