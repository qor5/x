package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXMessageListenerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXMessageListener() (r *VXMessageListenerBuilder) {
	r = &VXMessageListenerBuilder{
		tag: h.Tag("vx-messagelistener"),
	}
	return
}

func (b *VXMessageListenerBuilder) ListenFunc(v string) (r *VXMessageListenerBuilder) {
	b.tag.Attr(":listen-func", v)
	return b
}

func (b *VXMessageListenerBuilder) Name(v string) (r *VXMessageListenerBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VXMessageListenerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {

	return b.tag.MarshalHTML(ctx)
}
