package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VKbdBuilder struct {
	tag *h.HTMLTagBuilder
}

func VKbd(children ...h.HTMLComponent) (r *VKbdBuilder) {
	r = &VKbdBuilder{
		tag: h.Tag("v-kbd").Children(children...),
	}
	return
}

func (b *VKbdBuilder) Tag(v string) (r *VKbdBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VKbdBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VKbdBuilder) Attr(vs ...interface{}) (r *VKbdBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VKbdBuilder) Children(children ...h.HTMLComponent) (r *VKbdBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VKbdBuilder) AppendChildren(children ...h.HTMLComponent) (r *VKbdBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VKbdBuilder) PrependChildren(children ...h.HTMLComponent) (r *VKbdBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VKbdBuilder) Class(names ...string) (r *VKbdBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VKbdBuilder) ClassIf(name string, add bool) (r *VKbdBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VKbdBuilder) On(name string, value string) (r *VKbdBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VKbdBuilder) Bind(name string, value string) (r *VKbdBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VKbdBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
