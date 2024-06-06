package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListImgBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListImg(children ...h.HTMLComponent) (r *VListImgBuilder) {
	r = &VListImgBuilder{
		tag: h.Tag("v-list-img").Children(children...),
	}
	return
}

func (b *VListImgBuilder) Tag(v string) (r *VListImgBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListImgBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListImgBuilder) Attr(vs ...interface{}) (r *VListImgBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListImgBuilder) Children(children ...h.HTMLComponent) (r *VListImgBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListImgBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListImgBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListImgBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListImgBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListImgBuilder) Class(names ...string) (r *VListImgBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListImgBuilder) ClassIf(name string, add bool) (r *VListImgBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListImgBuilder) On(name string, value string) (r *VListImgBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListImgBuilder) Bind(name string, value string) (r *VListImgBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListImgBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
