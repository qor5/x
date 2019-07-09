package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VContentBuilder struct {
	tag *h.HTMLTagBuilder
}

func VContent(children ...h.HTMLComponent) (r *VContentBuilder) {
	r = &VContentBuilder{
		tag: h.Tag("v-content").Children(children...),
	}
	return
}

func (b *VContentBuilder) Tag(v string) (r *VContentBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VContentBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
