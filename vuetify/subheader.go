package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSubheaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSubheader(children ...h.HTMLComponent) (r *VSubheaderBuilder) {
	r = &VSubheaderBuilder{
		tag: h.Tag("v-subheader").Children(children...),
	}
	return
}

func (b *VSubheaderBuilder) Dark(v bool) (r *VSubheaderBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSubheaderBuilder) Inset(v bool) (r *VSubheaderBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VSubheaderBuilder) Light(v bool) (r *VSubheaderBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSubheaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
