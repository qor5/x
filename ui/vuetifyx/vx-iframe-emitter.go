package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXIframeEmitterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXIframeEmitter(children ...h.HTMLComponent) (r *VXIframeEmitterBuilder) {
	r = &VXIframeEmitterBuilder{
		tag: h.Tag("vx-iframe-emitter").Children(children...),
	}
	return
}

func (b *VXIframeEmitterBuilder) Ref(v interface{}) (r *VXIframeEmitterBuilder) {
	b.tag.Attr(":ref", h.JSONString(v))
	return b
}

func (b *VXIframeEmitterBuilder) Src(v interface{}) (r *VXIframeEmitterBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VXIframeEmitterBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXIframeEmitterBuilder) Attr(vs ...interface{}) (r *VXIframeEmitterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXIframeEmitterBuilder) Children(children ...h.HTMLComponent) (r *VXIframeEmitterBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXIframeEmitterBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXIframeEmitterBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXIframeEmitterBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXIframeEmitterBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXIframeEmitterBuilder) Class(names ...string) (r *VXIframeEmitterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXIframeEmitterBuilder) Style(names ...string) (r *VXIframeEmitterBuilder) {
	for _, name := range names {
		b.tag.Style(name)
	}
	return b
}

func (b *VXIframeEmitterBuilder) ClassIf(name string, add bool) (r *VXIframeEmitterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXIframeEmitterBuilder) On(name string, value string) (r *VXIframeEmitterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXIframeEmitterBuilder) Bind(name string, value string) (r *VXIframeEmitterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXIframeEmitterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
