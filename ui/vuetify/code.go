package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCodeBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCode(children ...h.HTMLComponent) (r *VCodeBuilder) {
	r = &VCodeBuilder{
		tag: h.Tag("v-code").Children(children...),
	}
	return
}

func (b *VCodeBuilder) Tag(v string) (r *VCodeBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCodeBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCodeBuilder) Attr(vs ...interface{}) (r *VCodeBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCodeBuilder) Children(children ...h.HTMLComponent) (r *VCodeBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCodeBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCodeBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCodeBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCodeBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCodeBuilder) Class(names ...string) (r *VCodeBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCodeBuilder) ClassIf(name string, add bool) (r *VCodeBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCodeBuilder) On(name string, value string) (r *VCodeBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCodeBuilder) Bind(name string, value string) (r *VCodeBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCodeBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
