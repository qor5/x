package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLazyBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLazy(children ...h.HTMLComponent) (r *VLazyBuilder) {
	r = &VLazyBuilder{
		tag: h.Tag("v-lazy").Children(children...),
	}
	return
}

func (b *VLazyBuilder) ModelValue(v bool) (r *VLazyBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VLazyBuilder) Options(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) Height(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) MaxHeight(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) MaxWidth(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) MinHeight(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) MinWidth(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) Width(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) Tag(v string) (r *VLazyBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VLazyBuilder) Transition(v interface{}) (r *VLazyBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VLazyBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VLazyBuilder) Attr(vs ...interface{}) (r *VLazyBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VLazyBuilder) Children(children ...h.HTMLComponent) (r *VLazyBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VLazyBuilder) AppendChildren(children ...h.HTMLComponent) (r *VLazyBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VLazyBuilder) PrependChildren(children ...h.HTMLComponent) (r *VLazyBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VLazyBuilder) Class(names ...string) (r *VLazyBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLazyBuilder) ClassIf(name string, add bool) (r *VLazyBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLazyBuilder) On(name string, value string) (r *VLazyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLazyBuilder) Bind(name string, value string) (r *VLazyBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLazyBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
