package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VListTileTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTileTitle(children ...h.HTMLComponent) (r *VListTileTitleBuilder) {
	r = &VListTileTitleBuilder{
		tag: h.Tag("v-list-tile-title").Children(children...),
	}
	return
}

func (b *VListTileTitleBuilder) Class(names ...string) (r *VListTileTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListTileTitleBuilder) ClassIf(name string, add bool) (r *VListTileTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListTileTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
