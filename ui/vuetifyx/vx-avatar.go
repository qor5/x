package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXAvatarBuilder struct {
	tag           *h.HTMLTagBuilder
}

func VXAvatar(children ...h.HTMLComponent) (r *VXAvatarBuilder) {
	r = &VXAvatarBuilder{
		tag: h.Tag("vx-avatar").Children(children...),
	}
	return
}

func (b *VXAvatarBuilder) Name(v string) (r *VXAvatarBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VXAvatarBuilder) Size(v string) (r *VXAvatarBuilder) {
	b.tag.Attr("size", v)
	return b
}

func (b *VXAvatarBuilder) Img(v string) (r *VXAvatarBuilder) {
	b.tag.Attr("img", v)
	return b
}

func (b *VXAvatarBuilder) Class(v string) (r *VXAvatarBuilder) {
	b.tag.Attr("class", v)
	return b
}

func (b *VXAvatarBuilder) Slot(name string, children ...h.HTMLComponent) *VXAvatarBuilder {
	slotTemplate := h.Tag("template").Attr("#"+name).Children(children...)
	b.tag.Children(slotTemplate)
	return b
}

func (b *VXAvatarBuilder) Attr(vs ...interface{}) (r *VXAvatarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXAvatarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
