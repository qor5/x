package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLabelBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLabel() (r *VLabelBuilder) {
	r = &VLabelBuilder{
		tag: h.Tag("v-label"),
	}
	return
}

func (b *VLabelBuilder) Absolute(v bool) (r *VLabelBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Color(v string) (r *VLabelBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VLabelBuilder) Dark(v bool) (r *VLabelBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Disabled(v bool) (r *VLabelBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Focused(v bool) (r *VLabelBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) For(v string) (r *VLabelBuilder) {
	b.tag.Attr("for", v)
	return b
}

func (b *VLabelBuilder) Left(v int) (r *VLabelBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Light(v bool) (r *VLabelBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Right(v int) (r *VLabelBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Value(v bool) (r *VLabelBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VLabelBuilder) Class(names ...string) (r *VLabelBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLabelBuilder) ClassIf(name string, add bool) (r *VLabelBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLabelBuilder) On(name string, value string) (r *VLabelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLabelBuilder) Bind(name string, value string) (r *VLabelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLabelBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
