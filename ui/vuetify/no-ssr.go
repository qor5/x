package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VNoSsrBuilder struct {
	tag *h.HTMLTagBuilder
}

func VNoSsr(children ...h.HTMLComponent) (r *VNoSsrBuilder) {
	r = &VNoSsrBuilder{
		tag: h.Tag("v-no-ssr").Children(children...),
	}
	return
}

func (b *VNoSsrBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VNoSsrBuilder) Attr(vs ...interface{}) (r *VNoSsrBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VNoSsrBuilder) Children(children ...h.HTMLComponent) (r *VNoSsrBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VNoSsrBuilder) AppendChildren(children ...h.HTMLComponent) (r *VNoSsrBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VNoSsrBuilder) PrependChildren(children ...h.HTMLComponent) (r *VNoSsrBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VNoSsrBuilder) Class(names ...string) (r *VNoSsrBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VNoSsrBuilder) ClassIf(name string, add bool) (r *VNoSsrBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VNoSsrBuilder) On(name string, value string) (r *VNoSsrBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VNoSsrBuilder) Bind(name string, value string) (r *VNoSsrBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VNoSsrBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
