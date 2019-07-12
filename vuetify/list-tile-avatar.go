package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListTileAvatarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListTileAvatar(children ...h.HTMLComponent) (r *VListTileAvatarBuilder) {
	r = &VListTileAvatarBuilder{
		tag: h.Tag("v-list-tile-avatar").Children(children...),
	}
	return
}

func (b *VListTileAvatarBuilder) Color(v string) (r *VListTileAvatarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListTileAvatarBuilder) Size(v int) (r *VListTileAvatarBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VListTileAvatarBuilder) Tile(v bool) (r *VListTileAvatarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VListTileAvatarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
