package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VContentBuilder struct {
	tag *h.HTMLTagBuilder
}

func VContent(children ...h.HTMLComponent) (r *VContentBuilder) {
	r = &VContentBuilder{
		tag: h.Tag("v-content").Children(children...),
	}
	return
}

func (b *VContentBuilder) Tag(v string) (r *VContentBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VContentBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VContentBuilder) Attr(vs ...interface{}) (r *VContentBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VContentBuilder) Children(children ...h.HTMLComponent) (r *VContentBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VContentBuilder) AppendChildren(children ...h.HTMLComponent) (r *VContentBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VContentBuilder) PrependChildren(children ...h.HTMLComponent) (r *VContentBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VContentBuilder) Class(names ...string) (r *VContentBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VContentBuilder) ClassIf(name string, add bool) (r *VContentBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VContentBuilder) On(name string, value string) (r *VContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VContentBuilder) Bind(name string, value string) (r *VContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VContentBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
