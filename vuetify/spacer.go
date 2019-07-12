package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VSpacerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSpacer() (r *VSpacerBuilder) {
	r = &VSpacerBuilder{
		tag: h.Tag("v-spacer"),
	}
	return
}

func (b *VSpacerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
