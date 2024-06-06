package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VComponentIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VComponentIcon(children ...h.HTMLComponent) (r *VComponentIconBuilder) {
	r = &VComponentIconBuilder{
		tag: h.Tag("v-component-icon").Children(children...),
	}
	return
}

func (b *VComponentIconBuilder) Icon(v interface{}) (r *VComponentIconBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VComponentIconBuilder) Tag(v string) (r *VComponentIconBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VComponentIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VComponentIconBuilder) Attr(vs ...interface{}) (r *VComponentIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VComponentIconBuilder) Children(children ...h.HTMLComponent) (r *VComponentIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VComponentIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VComponentIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VComponentIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VComponentIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VComponentIconBuilder) Class(names ...string) (r *VComponentIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VComponentIconBuilder) ClassIf(name string, add bool) (r *VComponentIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VComponentIconBuilder) On(name string, value string) (r *VComponentIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VComponentIconBuilder) Bind(name string, value string) (r *VComponentIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VComponentIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
