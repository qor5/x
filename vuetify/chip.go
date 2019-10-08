package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VChipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VChip(children ...h.HTMLComponent) (r *VChipBuilder) {
	r = &VChipBuilder{
		tag: h.Tag("v-chip").Children(children...),
	}
	return
}

func (b *VChipBuilder) Close(v bool) (r *VChipBuilder) {
	b.tag.Attr(":close", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Color(v string) (r *VChipBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VChipBuilder) Dark(v bool) (r *VChipBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Disabled(v bool) (r *VChipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Label(v bool) (r *VChipBuilder) {
	b.tag.Attr(":label", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Light(v bool) (r *VChipBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Outline(v bool) (r *VChipBuilder) {
	b.tag.Attr(":outline", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Selected(v bool) (r *VChipBuilder) {
	b.tag.Attr(":selected", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Small(v bool) (r *VChipBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) TextColor(v string) (r *VChipBuilder) {
	b.tag.Attr("text-color", v)
	return b
}

func (b *VChipBuilder) Value(v bool) (r *VChipBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Class(names ...string) (r *VChipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VChipBuilder) ClassIf(name string, add bool) (r *VChipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VChipBuilder) On(name string, value string) (r *VChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VChipBuilder) Bind(name string, value string) (r *VChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VChipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
