package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAppBuilder struct {
	tag *h.HTMLTagBuilder
}

func VApp(children ...h.HTMLComponent) (r *VAppBuilder) {
	r = &VAppBuilder{
		tag: h.Tag("v-app").Children(children...),
	}
	return
}

func (b *VAppBuilder) FullHeight(v bool) (r *VAppBuilder) {
	b.tag.Attr(":full-height", fmt.Sprint(v))
	return b
}

func (b *VAppBuilder) Overlaps(v interface{}) (r *VAppBuilder) {
	b.tag.Attr(":overlaps", h.JSONString(v))
	return b
}

func (b *VAppBuilder) Theme(v string) (r *VAppBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VAppBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAppBuilder) Attr(vs ...interface{}) (r *VAppBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAppBuilder) Children(children ...h.HTMLComponent) (r *VAppBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAppBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAppBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAppBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAppBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAppBuilder) Class(names ...string) (r *VAppBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAppBuilder) ClassIf(name string, add bool) (r *VAppBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAppBuilder) On(name string, value string) (r *VAppBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAppBuilder) Bind(name string, value string) (r *VAppBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAppBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
