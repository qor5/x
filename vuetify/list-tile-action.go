package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VListTileActionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTileAction(children ...h.HTMLComponent) (r *VListTileActionBuilder) {
	r = &VListTileActionBuilder{
		tag: h.Tag("v-list-tile-action").Children(children...),
	}
	return
}

func (b *VListTileActionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
