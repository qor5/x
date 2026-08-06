package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXImageGalleryBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXImageGallery() (r *VXImageGalleryBuilder) {
	r = &VXImageGalleryBuilder{
		tag: h.Tag("vx-image-gallery"),
	}
	return
}

func (b *VXImageGalleryBuilder) Items(v interface{}) (r *VXImageGalleryBuilder) {
	b.tag.Attr(":items", v)
	return b
}

func (b *VXImageGalleryBuilder) Height(v interface{}) (r *VXImageGalleryBuilder) {
	b.tag.Attr(":height", v)
	return b
}

func (b *VXImageGalleryBuilder) Width(v interface{}) (r *VXImageGalleryBuilder) {
	b.tag.Attr(":width", v)
	return b
}

func (b *VXImageGalleryBuilder) Attr(vs ...interface{}) (r *VXImageGalleryBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXImageGalleryBuilder) Class(names ...string) (r *VXImageGalleryBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXImageGalleryBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
