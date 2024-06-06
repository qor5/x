package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardTitle(children ...h.HTMLComponent) (r *VCardTitleBuilder) {
	r = &VCardTitleBuilder{
		tag: h.Tag("v-card-title").Children(children...),
	}
	return
}

func (b *VCardTitleBuilder) Tag(v string) (r *VCardTitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCardTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCardTitleBuilder) Attr(vs ...interface{}) (r *VCardTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCardTitleBuilder) Children(children ...h.HTMLComponent) (r *VCardTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCardTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCardTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCardTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCardTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCardTitleBuilder) Class(names ...string) (r *VCardTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardTitleBuilder) ClassIf(name string, add bool) (r *VCardTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardTitleBuilder) On(name string, value string) (r *VCardTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardTitleBuilder) Bind(name string, value string) (r *VCardTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
