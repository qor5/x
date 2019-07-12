package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDividerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDivider() (r *VDividerBuilder) {
	r = &VDividerBuilder{
		tag: h.Tag("v-divider"),
	}
	return
}

func (b *VDividerBuilder) Dark(v bool) (r *VDividerBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VDividerBuilder) Inset(v bool) (r *VDividerBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VDividerBuilder) Light(v bool) (r *VDividerBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VDividerBuilder) Vertical(v bool) (r *VDividerBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VDividerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
