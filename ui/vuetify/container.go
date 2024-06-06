package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VContainerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VContainer(children ...h.HTMLComponent) (r *VContainerBuilder) {
	r = &VContainerBuilder{
		tag: h.Tag("v-container").Children(children...),
	}
	return
}

func (b *VContainerBuilder) Fluid(v bool) (r *VContainerBuilder) {
	b.tag.Attr(":fluid", fmt.Sprint(v))
	return b
}

func (b *VContainerBuilder) Tag(v string) (r *VContainerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VContainerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VContainerBuilder) Attr(vs ...interface{}) (r *VContainerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VContainerBuilder) Children(children ...h.HTMLComponent) (r *VContainerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VContainerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VContainerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VContainerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VContainerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VContainerBuilder) Class(names ...string) (r *VContainerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VContainerBuilder) ClassIf(name string, add bool) (r *VContainerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VContainerBuilder) On(name string, value string) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VContainerBuilder) Bind(name string, value string) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VContainerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
