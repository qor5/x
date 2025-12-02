package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXVideoBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXVideo() (r *VXVideoBuilder) {
	r = &VXVideoBuilder{
		tag: h.Tag("vx-video"),
	}
	return
}

func (b *VXVideoBuilder) Src(v string) (r *VXVideoBuilder) {
	b.tag.Attr("src", v)
	return b
}

func (b *VXVideoBuilder) Poster(v string) (r *VXVideoBuilder) {
	b.tag.Attr("poster", v)
	return b
}

func (b *VXVideoBuilder) Width(v string) (r *VXVideoBuilder) {
	b.tag.Attr("width", v)
	return b
}

func (b *VXVideoBuilder) Height(v string) (r *VXVideoBuilder) {
	b.tag.Attr("height", v)
	return b
}

func (b *VXVideoBuilder) Class(names ...string) (r *VXVideoBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXVideoBuilder) ClassIf(name string, add bool) (r *VXVideoBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXVideoBuilder) Attr(vs ...interface{}) (r *VXVideoBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXVideoBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
