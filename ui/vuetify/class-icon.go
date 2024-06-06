package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VClassIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VClassIcon(children ...h.HTMLComponent) (r *VClassIconBuilder) {
	r = &VClassIconBuilder{
		tag: h.Tag("v-class-icon").Children(children...),
	}
	return
}

func (b *VClassIconBuilder) Icon(v interface{}) (r *VClassIconBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VClassIconBuilder) Tag(v string) (r *VClassIconBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VClassIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VClassIconBuilder) Attr(vs ...interface{}) (r *VClassIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VClassIconBuilder) Children(children ...h.HTMLComponent) (r *VClassIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VClassIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VClassIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VClassIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VClassIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VClassIconBuilder) Class(names ...string) (r *VClassIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VClassIconBuilder) ClassIf(name string, add bool) (r *VClassIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VClassIconBuilder) On(name string, value string) (r *VClassIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VClassIconBuilder) Bind(name string, value string) (r *VClassIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VClassIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
