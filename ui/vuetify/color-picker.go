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

func (b *VColorPickerBuilder) CanvasHeight(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":canvas-height", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Disabled(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) DotSize(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":dot-size", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) HideCanvas(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-canvas", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) HideSliders(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-sliders", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) HideInputs(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":hide-inputs", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Mode(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":mode", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Modes(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":modes", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) ShowSwatches(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":show-swatches", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) SwatchesMaxHeight(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":swatches-max-height", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) ModelValue(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Color(v string) (r *VColorPickerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VColorPickerBuilder) Border(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Width(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Elevation(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Position(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Rounded(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VColorPickerBuilder) Tile(v bool) (r *VColorPickerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VColorPickerBuilder) Tag(v string) (r *VColorPickerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VColorPickerBuilder) Theme(v string) (r *VColorPickerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VColorPickerBuilder) Swatches(v interface{}) (r *VColorPickerBuilder) {
	b.tag.Attr(":swatches", h.JSONString(v))
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
