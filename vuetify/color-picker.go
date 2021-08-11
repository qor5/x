package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VColorPickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VColorPicker(children ...h.HTMLComponent) (r *VColorPickerBuilder) {
	r = &VColorPickerBuilder{
		tag: h.Tag("v-color-picker").Children(children...),
	}
	return
}

func (b *VColorPickerBuilder) CanvasHeight(v string) (r *VColorPickerBuilder) {
	b.tag.Attr("canvas-height", v)
	return b
}

func (b *VColorPickerBuilder) Dark(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Disabled(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) DotSize(v int) (r *VColorPickerBuilder) {
	b.tag.Attr(":dot-size", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Elevation(v int) (r *VColorPickerBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Flat(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) HideCanvas(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-canvas", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) HideInputs(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-inputs", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) HideModeSwitch(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-mode-switch", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) HideSliders(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-sliders", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Light(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Mode(v string) (r *VColorPickerBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VColorPickerBuilder) ShowSwatches(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":show-swatches", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Swatches(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":swatches", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) SwatchesMaxHeight(v int) (r *VColorPickerBuilder) {
	b.tag.Attr(":swatches-max-height", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Value(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Width(v int) (r *VColorPickerBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VColorPickerBuilder) Attr(vs ...interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VColorPickerBuilder) Children(children ...h.HTMLComponent) (r *VColorPickerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VColorPickerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VColorPickerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VColorPickerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VColorPickerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VColorPickerBuilder) Class(names ...string) (r *VColorPickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VColorPickerBuilder) ClassIf(name string, add bool) (r *VColorPickerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VColorPickerBuilder) On(name string, value string) (r *VColorPickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VColorPickerBuilder) Bind(name string, value string) (r *VColorPickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VColorPickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
