package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VToolbarSideIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VToolbarSideIcon() (r *VToolbarSideIconBuilder) {
	r = &VToolbarSideIconBuilder{
		tag: h.Tag("v-toolbar-side-icon"),
	}
	return
}

func (b *VToolbarSideIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
