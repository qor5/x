package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXModelProxyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXModelProxy() (r *VXModelProxyBuilder) {
	r = &VXModelProxyBuilder{
		tag: h.Tag("vx-model-proxy"),
	}
	return
}

func (b *VXModelProxyBuilder) ModelValue(v interface{}) (r *VXModelProxyBuilder) {
	b.tag.Attr("v-model", v)
	return b
}

func (b *VXModelProxyBuilder) FormatModel(v string) (r *VXModelProxyBuilder) {
	b.tag.Attr("format-model", v)
	return b
}

func (b *VXModelProxyBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXModelProxyBuilder) Attr(vs ...interface{}) (r *VXModelProxyBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXModelProxyBuilder) Children(children ...h.HTMLComponent) (r *VXModelProxyBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXModelProxyBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXModelProxyBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXModelProxyBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXModelProxyBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXModelProxyBuilder) Class(names ...string) (r *VXModelProxyBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXModelProxyBuilder) ClassIf(name string, add bool) (r *VXModelProxyBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXModelProxyBuilder) On(name string, value string) (r *VXModelProxyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXModelProxyBuilder) Bind(name string, value string) (r *VXModelProxyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXModelProxyBuilder) Slot(name string, value interface{}) (r *VXModelProxyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-slot:%s", name), value)
	return b
}

func (b *VXModelProxyBuilder) DefaultSlot(value interface{}) (r *VXModelProxyBuilder) {
	b.tag.Attr("v-slot", value)
	return b
}

func (b *VXModelProxyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
