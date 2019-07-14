package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBreadcrumbsDividerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBreadcrumbsDivider() (r *VBreadcrumbsDividerBuilder) {
	r = &VBreadcrumbsDividerBuilder{
		tag: h.Tag("v-breadcrumbs-divider"),
	}
	return
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
