package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRadioBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRadio() (r *VRadioBuilder) {
	r = &VRadioBuilder{
		tag: h.Tag("v-radio"),
	}
	return
}

func (b *VRadioBuilder) Color(v string) (r *VRadioBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRadioBuilder) Dark(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Disabled(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Label(v string) (r *VRadioBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRadioBuilder) Light(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) OffIcon(v string) (r *VRadioBuilder) {
	b.tag.Attr("off-icon", v)
	return b
}

func (b *VRadioBuilder) OnIcon(v string) (r *VRadioBuilder) {
	b.tag.Attr("on-icon", v)
	return b
}

func (b *VRadioBuilder) Readonly(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Ripple(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Value(v string) (r *VRadioBuilder) {
	b.tag.Attr("value", v)
	return b
}

func (b *VRadioBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
