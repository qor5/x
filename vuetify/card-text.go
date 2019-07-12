package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VCardTextBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardText(text string) (r *VCardTextBuilder) {
	r = &VCardTextBuilder{
		tag: h.Tag("v-card-text").Text(text),
	}
	return
}

func (b *VCardTextBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
