package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VToolbarTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VToolbarTitle(text string) (r *VToolbarTitleBuilder) {
	r = &VToolbarTitleBuilder{
		tag: h.Tag("v-toolbar-title").Text(text),
	}
	return
}

func (b *VToolbarTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
