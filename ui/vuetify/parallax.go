package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VParallaxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VParallax(children ...h.HTMLComponent) (r *VParallaxBuilder) {
	r = &VParallaxBuilder{
		tag: h.Tag("v-parallax").Children(children...),
	}
	return
}

func (b *VParallaxBuilder) Scale(v interface{}) (r *VParallaxBuilder) {
	b.tag.Attr(":scale", h.JSONString(v))
	return b
}

func (b *VParallaxBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VParallaxBuilder) Attr(vs ...interface{}) (r *VParallaxBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VParallaxBuilder) Children(children ...h.HTMLComponent) (r *VParallaxBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VParallaxBuilder) AppendChildren(children ...h.HTMLComponent) (r *VParallaxBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VParallaxBuilder) PrependChildren(children ...h.HTMLComponent) (r *VParallaxBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VParallaxBuilder) Class(names ...string) (r *VParallaxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VParallaxBuilder) ClassIf(name string, add bool) (r *VParallaxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VParallaxBuilder) On(name string, value string) (r *VParallaxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VParallaxBuilder) Bind(name string, value string) (r *VParallaxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VParallaxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
