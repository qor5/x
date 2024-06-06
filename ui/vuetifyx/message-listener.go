package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXMessageListenerBuilder struct {
	tag        *h.HTMLTagBuilder
	listenFunc string
}

func VXMessageListener() (r *VXMessageListenerBuilder) {
	r = &VXMessageListenerBuilder{
		tag: h.Tag("vx-messagelistener"),
	}
	return
}

func (b *VXMessageListenerBuilder) ListenFunc(v string) (r *VXMessageListenerBuilder) {
	b.listenFunc = v
	return b
}

func (b *VXMessageListenerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.listenFunc != "" {
		b.tag.Attr(":listen-func", b.listenFunc)
	}

	return b.tag.MarshalHTML(ctx)
}
