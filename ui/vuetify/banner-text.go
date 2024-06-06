package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBannerTextBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBannerText(children ...h.HTMLComponent) (r *VBannerTextBuilder) {
	r = &VBannerTextBuilder{
		tag: h.Tag("v-banner-text").Children(children...),
	}
	return
}

func (b *VBannerTextBuilder) Tag(v string) (r *VBannerTextBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBannerTextBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBannerTextBuilder) Attr(vs ...interface{}) (r *VBannerTextBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBannerTextBuilder) Children(children ...h.HTMLComponent) (r *VBannerTextBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBannerTextBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBannerTextBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBannerTextBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBannerTextBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBannerTextBuilder) Class(names ...string) (r *VBannerTextBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBannerTextBuilder) ClassIf(name string, add bool) (r *VBannerTextBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBannerTextBuilder) On(name string, value string) (r *VBannerTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBannerTextBuilder) Bind(name string, value string) (r *VBannerTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBannerTextBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
