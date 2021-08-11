package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VImgBuilder struct {
	tag *h.HTMLTagBuilder
}

func VImg(children ...h.HTMLComponent) (r *VImgBuilder) {
	r = &VImgBuilder{
		tag: h.Tag("v-img").Children(children...),
	}
	return
}

func (b *VImgBuilder) Alt(v string) (r *VImgBuilder) {
	b.tag.Attr("alt", v)
	return b
}

func (b *VImgBuilder) AspectRatio(v string) (r *VImgBuilder) {
	b.tag.Attr("aspect-ratio", v)
	return b
}

func (b *VImgBuilder) Contain(v bool) (r *VImgBuilder) {
	b.tag.Attr(":contain", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) ContentClass(v string) (r *VImgBuilder) {
	b.tag.Attr("content-class", v)
	return b
}

func (b *VImgBuilder) Dark(v bool) (r *VImgBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Eager(v bool) (r *VImgBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Gradient(v string) (r *VImgBuilder) {
	b.tag.Attr("gradient", v)
	return b
}

func (b *VImgBuilder) Height(v int) (r *VImgBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) LazySrc(v string) (r *VImgBuilder) {
	b.tag.Attr("lazy-src", v)
	return b
}

func (b *VImgBuilder) Light(v bool) (r *VImgBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) MaxHeight(v int) (r *VImgBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) MaxWidth(v int) (r *VImgBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) MinHeight(v int) (r *VImgBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) MinWidth(v int) (r *VImgBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Options(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Position(v string) (r *VImgBuilder) {
	b.tag.Attr("position", v)
	return b
}

func (b *VImgBuilder) Sizes(v string) (r *VImgBuilder) {
	b.tag.Attr("sizes", v)
	return b
}

func (b *VImgBuilder) Src(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Srcset(v string) (r *VImgBuilder) {
	b.tag.Attr("srcset", v)
	return b
}

func (b *VImgBuilder) Transition(v bool) (r *VImgBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Width(v int) (r *VImgBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VImgBuilder) Attr(vs ...interface{}) (r *VImgBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VImgBuilder) Children(children ...h.HTMLComponent) (r *VImgBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VImgBuilder) AppendChildren(children ...h.HTMLComponent) (r *VImgBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VImgBuilder) PrependChildren(children ...h.HTMLComponent) (r *VImgBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VImgBuilder) Class(names ...string) (r *VImgBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VImgBuilder) ClassIf(name string, add bool) (r *VImgBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VImgBuilder) On(name string, value string) (r *VImgBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VImgBuilder) Bind(name string, value string) (r *VImgBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VImgBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
