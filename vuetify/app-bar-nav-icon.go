package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAppBarNavIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAppBarNavIcon(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	r = &VAppBarNavIconBuilder{
		tag: h.Tag("v-app-bar-nav-icon").Children(children...),
	}
	return
}

func (b *VAppBarNavIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAppBarNavIconBuilder) Attr(vs ...interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAppBarNavIconBuilder) Children(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAppBarNavIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAppBarNavIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAppBarNavIconBuilder) Class(names ...string) (r *VAppBarNavIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAppBarNavIconBuilder) ClassIf(name string, add bool) (r *VAppBarNavIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAppBarNavIconBuilder) On(name string, value string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAppBarNavIconBuilder) Bind(name string, value string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAppBarNavIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
