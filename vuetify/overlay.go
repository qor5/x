package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VOverlayBuilder struct {
	tag *h.HTMLTagBuilder
}

func VOverlay(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	r = &VOverlayBuilder{
		tag: h.Tag("v-overlay").Children(children...),
	}
	return
}

func (b *VOverlayBuilder) Absolute(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Color(v string) (r *VOverlayBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VOverlayBuilder) Dark(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Light(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Opacity(v int) (r *VOverlayBuilder) {
	b.tag.Attr(":opacity", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Value(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) ZIndex(v int) (r *VOverlayBuilder) {
	b.tag.Attr(":z-index", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VOverlayBuilder) Attr(vs ...interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VOverlayBuilder) Children(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VOverlayBuilder) AppendChildren(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VOverlayBuilder) PrependChildren(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VOverlayBuilder) Class(names ...string) (r *VOverlayBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VOverlayBuilder) ClassIf(name string, add bool) (r *VOverlayBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VOverlayBuilder) On(name string, value string) (r *VOverlayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VOverlayBuilder) Bind(name string, value string) (r *VOverlayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VOverlayBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
