package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VListTileContentBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTileContent(children ...h.HTMLComponent) (r *VListTileContentBuilder) {
	r = &VListTileContentBuilder{
		tag: h.Tag("v-list-tile-content").Children(children...),
	}
	return
}

func (b *VListTileContentBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
