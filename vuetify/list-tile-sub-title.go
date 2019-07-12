package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VListTileSubTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTileSubTitle(children ...h.HTMLComponent) (r *VListTileSubTitleBuilder) {
	r = &VListTileSubTitleBuilder{
		tag: h.Tag("v-list-tile-sub-title").Children(children...),
	}
	return
}

func (b *VListTileSubTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
