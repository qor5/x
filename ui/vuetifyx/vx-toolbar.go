package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXToolbarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXToolbar(children ...h.HTMLComponent) (r *VXToolbarBuilder) {
	r = &VXToolbarBuilder{
		tag: h.Tag("vx-toolbar").Children(children...),
	}
	return
}

func (b *VXToolbarBuilder) Text(v string) (r *VXToolbarBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VXToolbarBuilder) Placeholder(v string) (r *VXToolbarBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VXToolbarBuilder) Attr(vs ...interface{}) (r *VXToolbarBuilder) {
	b.tag.Attr(vs...)
	return b
}
func (b *VXToolbarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXToolbarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
