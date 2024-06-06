package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardTextBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardText(children ...h.HTMLComponent) (r *VCardTextBuilder) {
	r = &VCardTextBuilder{
		tag: h.Tag("v-card-text").Children(children...),
	}
	return
}

func (b *VCardTextBuilder) Opacity(v interface{}) (r *VCardTextBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VCardTextBuilder) Tag(v string) (r *VCardTextBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCardTextBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCardTextBuilder) Attr(vs ...interface{}) (r *VCardTextBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCardTextBuilder) Children(children ...h.HTMLComponent) (r *VCardTextBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCardTextBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCardTextBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCardTextBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCardTextBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCardTextBuilder) Class(names ...string) (r *VCardTextBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardTextBuilder) ClassIf(name string, add bool) (r *VCardTextBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardTextBuilder) On(name string, value string) (r *VCardTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardTextBuilder) Bind(name string, value string) (r *VCardTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardTextBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
