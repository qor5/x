package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXOverlayBuilder struct {
	tag      *h.HTMLTagBuilder
	maxWidth int
}

func VXOverlay(children ...h.HTMLComponent) (r *VXOverlayBuilder) {
	r = &VXOverlayBuilder{
		tag: h.Tag("vx-overlay").Children(children...),
	}
	return
}

func (b *VXOverlayBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXOverlayBuilder) Attr(vs ...interface{}) (r *VXOverlayBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXOverlayBuilder) MaxWidth(v int) (r *VXOverlayBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VXOverlayBuilder) Children(children ...h.HTMLComponent) (r *VXOverlayBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXOverlayBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXOverlayBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXOverlayBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXOverlayBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXOverlayBuilder) Class(names ...string) (r *VXOverlayBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXOverlayBuilder) ClassIf(name string, add bool) (r *VXOverlayBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXOverlayBuilder) On(name string, value string) (r *VXOverlayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXOverlayBuilder) Bind(name string, value string) (r *VXOverlayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXOverlayBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
