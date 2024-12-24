package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXScrollIframeBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXScrollIframe(children ...h.HTMLComponent) (r *VXScrollIframeBuilder) {
	r = &VXScrollIframeBuilder{
		tag: h.Tag("vx-scroll-iframe").Children(children...),
	}
	return
}

func (b *VXScrollIframeBuilder) Srcdoc(v interface{}) (r *VXScrollIframeBuilder) {
	b.tag.Attr(":srcdoc", h.JSONString(v))
	return b
}

func (b *VXScrollIframeBuilder) Width(v interface{}) (r *VXScrollIframeBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}
func (b *VXScrollIframeBuilder) VirtualElementText(v string) (r *VXScrollIframeBuilder) {
	b.tag.Attr(":virtual-element-text", h.JSONString(v))
	return b
}
func (b *VXScrollIframeBuilder) virtualElementHeight(v int) (r *VXScrollIframeBuilder) {
	b.tag.Attr(":virtual-element-height", h.JSONString(v))
	return b
}
func (b *VXScrollIframeBuilder) BackgroundColor(v string) (r *VXScrollIframeBuilder) {
	b.tag.Attr(":background-color", h.JSONString(v))
	return b
}
func (b *VXScrollIframeBuilder) UpdateDifferent(v bool) (r *VXScrollIframeBuilder) {
	b.tag.Attr(":update-different", h.JSONString(v))
	return b
}

func (b *VXScrollIframeBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXScrollIframeBuilder) Attr(vs ...interface{}) (r *VXScrollIframeBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXScrollIframeBuilder) Children(children ...h.HTMLComponent) (r *VXScrollIframeBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXScrollIframeBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXScrollIframeBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXScrollIframeBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXScrollIframeBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXScrollIframeBuilder) Class(names ...string) (r *VXScrollIframeBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXScrollIframeBuilder) ClassIf(name string, add bool) (r *VXScrollIframeBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXScrollIframeBuilder) On(name string, value string) (r *VXScrollIframeBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXScrollIframeBuilder) Bind(name string, value string) (r *VXScrollIframeBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXScrollIframeBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
