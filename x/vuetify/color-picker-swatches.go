package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VColorPickerSwatchesBuilder struct {
	tag *h.HTMLTagBuilder
}

func VColorPickerSwatches() (r *VColorPickerSwatchesBuilder) {
	r = &VColorPickerSwatchesBuilder{
		tag: h.Tag("v-color-picker-swatches"),
	}
	return
}

func (b *VColorPickerSwatchesBuilder) Color(v interface{}) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(":color", v)
	return b
}

func (b *VColorPickerSwatchesBuilder) Dark(v bool) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VColorPickerSwatchesBuilder) Light(v bool) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VColorPickerSwatchesBuilder) MaxHeight(v int) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VColorPickerSwatchesBuilder) MaxWidth(v int) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VColorPickerSwatchesBuilder) Swatches(v []string) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(":swatches", v)
	return b
}

func (b *VColorPickerSwatchesBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VColorPickerSwatchesBuilder) Attr(vs ...interface{}) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VColorPickerSwatchesBuilder) Children(children ...h.HTMLComponent) (r *VColorPickerSwatchesBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VColorPickerSwatchesBuilder) AppendChildren(children ...h.HTMLComponent) (r *VColorPickerSwatchesBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VColorPickerSwatchesBuilder) PrependChildren(children ...h.HTMLComponent) (r *VColorPickerSwatchesBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VColorPickerSwatchesBuilder) Class(names ...string) (r *VColorPickerSwatchesBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VColorPickerSwatchesBuilder) ClassIf(name string, add bool) (r *VColorPickerSwatchesBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VColorPickerSwatchesBuilder) On(name string, value string) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VColorPickerSwatchesBuilder) Bind(name string, value string) (r *VColorPickerSwatchesBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VColorPickerSwatchesBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
