package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VCardTextBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardText(children ...h.HTMLComponent) (r *VCardTextBuilder) {
	r = &VCardTextBuilder{
		tag: h.Tag("v-card-text").Children(children...),
	}
	return
}

func (b *VCardTextBuilder) Class(names ...string) (r *VCardTextBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardTextBuilder) ClassIf(name string, add bool) (r *VCardTextBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardTextBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
