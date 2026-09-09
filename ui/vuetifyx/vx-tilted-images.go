package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXTiltedImagesBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXTiltedImages(children ...h.HTMLComponent) (r *VXTiltedImagesBuilder) {
	r = &VXTiltedImagesBuilder{
		tag: h.Tag("vx-tilted-images").Children(children...),
	}
	return
}

func (b *VXTiltedImagesBuilder) InitialRotateX(v float64) (r *VXTiltedImagesBuilder) {
	b.tag.Attr(":initial-rotate-x", fmt.Sprint(v))
	return b
}

func (b *VXTiltedImagesBuilder) InitialRotateY(v float64) (r *VXTiltedImagesBuilder) {
	b.tag.Attr(":initial-rotate-y", fmt.Sprint(v))
	return b
}

func (b *VXTiltedImagesBuilder) InitialTranslateX(v float64) (r *VXTiltedImagesBuilder) {
	b.tag.Attr(":initial-translate-x", fmt.Sprint(v))
	return b
}

func (b *VXTiltedImagesBuilder) InitialTranslateY(v float64) (r *VXTiltedImagesBuilder) {
	b.tag.Attr(":initial-translate-y", fmt.Sprint(v))
	return b
}

func (b *VXTiltedImagesBuilder) Attr(vs ...interface{}) (r *VXTiltedImagesBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTiltedImagesBuilder) Class(names ...string) (r *VXTiltedImagesBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXTiltedImagesBuilder) Style(v string) (r *VXTiltedImagesBuilder) {
	b.tag.Style(v)
	return b
}

func (b *VXTiltedImagesBuilder) Children(children ...h.HTMLComponent) (r *VXTiltedImagesBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXTiltedImagesBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXTiltedImagesBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXTiltedImagesBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXTiltedImagesBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXTiltedImagesBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
