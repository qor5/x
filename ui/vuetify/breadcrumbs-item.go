package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBreadcrumbsItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBreadcrumbsItem(children ...h.HTMLComponent) (r *VBreadcrumbsItemBuilder) {
	r = &VBreadcrumbsItemBuilder{
		tag: h.Tag("v-breadcrumbs-item").Children(children...),
	}
	return
}

func (b *VBreadcrumbsItemBuilder) Active(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) ActiveClass(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) ActiveColor(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Color(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Disabled(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Title(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Href(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Replace(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Exact(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) To(v interface{}) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Tag(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBreadcrumbsItemBuilder) Attr(vs ...interface{}) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBreadcrumbsItemBuilder) Children(children ...h.HTMLComponent) (r *VBreadcrumbsItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBreadcrumbsItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBreadcrumbsItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBreadcrumbsItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBreadcrumbsItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBreadcrumbsItemBuilder) Class(names ...string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBreadcrumbsItemBuilder) ClassIf(name string, add bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBreadcrumbsItemBuilder) On(name string, value string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBreadcrumbsItemBuilder) Bind(name string, value string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBreadcrumbsItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
