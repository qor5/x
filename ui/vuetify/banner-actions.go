package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBannerActionsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBannerActions(children ...h.HTMLComponent) (r *VBannerActionsBuilder) {
	r = &VBannerActionsBuilder{
		tag: h.Tag("v-banner-actions").Children(children...),
	}
	return
}

func (b *VBannerActionsBuilder) Color(v string) (r *VBannerActionsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBannerActionsBuilder) Density(v string) (r *VBannerActionsBuilder) {
	b.tag.Attr("density", v)
	return b
}

func (b *VBannerActionsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBannerActionsBuilder) Attr(vs ...interface{}) (r *VBannerActionsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBannerActionsBuilder) Children(children ...h.HTMLComponent) (r *VBannerActionsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBannerActionsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBannerActionsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBannerActionsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBannerActionsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBannerActionsBuilder) Class(names ...string) (r *VBannerActionsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBannerActionsBuilder) ClassIf(name string, add bool) (r *VBannerActionsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBannerActionsBuilder) On(name string, value string) (r *VBannerActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBannerActionsBuilder) Bind(name string, value string) (r *VBannerActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBannerActionsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
