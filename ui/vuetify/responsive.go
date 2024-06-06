package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VResponsiveBuilder struct {
	tag *h.HTMLTagBuilder
}

func VResponsive(children ...h.HTMLComponent) (r *VResponsiveBuilder) {
	r = &VResponsiveBuilder{
		tag: h.Tag("v-responsive").Children(children...),
	}
	return
}

func (b *VResponsiveBuilder) AspectRatio(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":aspect-ratio", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) ContentClass(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) Inline(v bool) (r *VResponsiveBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VResponsiveBuilder) Height(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) MaxHeight(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) MaxWidth(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) MinHeight(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) MinWidth(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) Width(v interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VResponsiveBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VResponsiveBuilder) Attr(vs ...interface{}) (r *VResponsiveBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VResponsiveBuilder) Children(children ...h.HTMLComponent) (r *VResponsiveBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VResponsiveBuilder) AppendChildren(children ...h.HTMLComponent) (r *VResponsiveBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VResponsiveBuilder) PrependChildren(children ...h.HTMLComponent) (r *VResponsiveBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VResponsiveBuilder) Class(names ...string) (r *VResponsiveBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VResponsiveBuilder) ClassIf(name string, add bool) (r *VResponsiveBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VResponsiveBuilder) On(name string, value string) (r *VResponsiveBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VResponsiveBuilder) Bind(name string, value string) (r *VResponsiveBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VResponsiveBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
