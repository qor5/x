package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VCardActionsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardActions(children ...h.HTMLComponent) (r *VCardActionsBuilder) {
	r = &VCardActionsBuilder{
		tag: h.Tag("v-card-actions").Children(children...),
	}
	return
}

func (b *VCardActionsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
