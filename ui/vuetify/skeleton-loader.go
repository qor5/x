package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSkeletonLoaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSkeletonLoader(children ...h.HTMLComponent) (r *VSkeletonLoaderBuilder) {
	r = &VSkeletonLoaderBuilder{
		tag: h.Tag("v-skeleton-loader").Children(children...),
	}
	return
}

func (b *VSkeletonLoaderBuilder) Boilerplate(v bool) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":boilerplate", fmt.Sprint(v))
	return b
}

func (b *VSkeletonLoaderBuilder) Color(v string) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSkeletonLoaderBuilder) Loading(v bool) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VSkeletonLoaderBuilder) LoadingText(v string) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr("loading-text", v)
	return b
}

func (b *VSkeletonLoaderBuilder) Type(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":type", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) Height(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) MaxHeight(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) MaxWidth(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) MinHeight(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) MinWidth(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) Width(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) Elevation(v interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VSkeletonLoaderBuilder) Theme(v string) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSkeletonLoaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSkeletonLoaderBuilder) Attr(vs ...interface{}) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSkeletonLoaderBuilder) Children(children ...h.HTMLComponent) (r *VSkeletonLoaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSkeletonLoaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSkeletonLoaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSkeletonLoaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSkeletonLoaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSkeletonLoaderBuilder) Class(names ...string) (r *VSkeletonLoaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSkeletonLoaderBuilder) ClassIf(name string, add bool) (r *VSkeletonLoaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSkeletonLoaderBuilder) On(name string, value string) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSkeletonLoaderBuilder) Bind(name string, value string) (r *VSkeletonLoaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSkeletonLoaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
