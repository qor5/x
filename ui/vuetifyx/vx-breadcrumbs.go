package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXBreadcrumbsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXBreadcrumbs(children ...h.HTMLComponent) (r *VXBreadcrumbsBuilder) {
	r = &VXBreadcrumbsBuilder{
		tag: h.Tag("vx-breadcrumbs").Children(children...),
	}
	return
}

func (b *VXBreadcrumbsBuilder) Divider(v string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr("divider", v)
	return b
}

func (b *VXBreadcrumbsBuilder) ActiveClass(v string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VXBreadcrumbsBuilder) ActiveColor(v string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VXBreadcrumbsBuilder) BgColor(v string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VXBreadcrumbsBuilder) Color(v string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXBreadcrumbsBuilder) Disabled(v bool) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXBreadcrumbsBuilder) Icon(v interface{}) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VXBreadcrumbsBuilder) Items(v interface{}) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VXBreadcrumbsBuilder) Density(v interface{}) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXBreadcrumbsBuilder) Rounded(v interface{}) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXBreadcrumbsBuilder) Tile(v bool) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXBreadcrumbsBuilder) Tag(v string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXBreadcrumbsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXBreadcrumbsBuilder) Attr(vs ...interface{}) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXBreadcrumbsBuilder) Children(children ...h.HTMLComponent) (r *VXBreadcrumbsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXBreadcrumbsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXBreadcrumbsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXBreadcrumbsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXBreadcrumbsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXBreadcrumbsBuilder) Class(names ...string) (r *VXBreadcrumbsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXBreadcrumbsBuilder) ClassIf(name string, add bool) (r *VXBreadcrumbsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXBreadcrumbsBuilder) On(name string, value string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXBreadcrumbsBuilder) Bind(name string, value string) (r *VXBreadcrumbsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXBreadcrumbsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
