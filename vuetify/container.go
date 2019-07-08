package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VContainerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VContainer(children ...h.HTMLComponent) (r *VContainerBuilder) {
	r = &VContainerBuilder{
		tag: h.Tag("v-container"),
	}
	r.tag.Children(children...)
	return
}

func (b *VContainerBuilder) AlignCenter(v bool) (r *VContainerBuilder) {
	b.tag.Attr("align-center", v)
	return b
}

func (b *VContainerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
