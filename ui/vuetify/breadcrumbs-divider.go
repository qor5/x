package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBreadcrumbsDividerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBreadcrumbsDivider(children ...h.HTMLComponent) (r *VBreadcrumbsDividerBuilder) {
	r = &VBreadcrumbsDividerBuilder{
		tag: h.Tag("v-breadcrumbs-divider").Children(children...),
	}
	return
}

func (b *VBreadcrumbsDividerBuilder) Divider(v interface{}) (r *VBreadcrumbsDividerBuilder) {
	b.tag.Attr(":divider", h.JSONString(v))
	return b
}

func (b *VBreadcrumbsDividerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBreadcrumbsDividerBuilder) Attr(vs ...interface{}) (r *VBreadcrumbsDividerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBreadcrumbsDividerBuilder) Children(children ...h.HTMLComponent) (r *VBreadcrumbsDividerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBreadcrumbsDividerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBreadcrumbsDividerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBreadcrumbsDividerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBreadcrumbsDividerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBreadcrumbsDividerBuilder) Class(names ...string) (r *VBreadcrumbsDividerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBreadcrumbsDividerBuilder) ClassIf(name string, add bool) (r *VBreadcrumbsDividerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBreadcrumbsDividerBuilder) On(name string, value string) (r *VBreadcrumbsDividerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBreadcrumbsDividerBuilder) Bind(name string, value string) (r *VBreadcrumbsDividerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBreadcrumbsDividerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
