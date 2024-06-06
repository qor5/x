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

func (b *VImgBuilder) Cover(v bool) (r *VImgBuilder) {
	b.tag.Attr(":cover", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Color(v string) (r *VImgBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VImgBuilder) Draggable(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":draggable", h.JSONString(v))
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

func (b *VImgBuilder) LazySrc(v string) (r *VImgBuilder) {
	b.tag.Attr("lazy-src", v)
	return b
}

func (b *VImgBuilder) Options(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
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

func (b *VImgBuilder) Position(v string) (r *VImgBuilder) {
	b.tag.Attr("position", v)
	return b
}

func (b *VImgBuilder) AspectRatio(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":aspect-ratio", h.JSONString(v))
	return b
}

func (b *VImgBuilder) ContentClass(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Inline(v bool) (r *VImgBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Height(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VImgBuilder) MaxHeight(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VImgBuilder) MaxWidth(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VImgBuilder) MinHeight(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VImgBuilder) MinWidth(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Width(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Rounded(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Tile(v bool) (r *VImgBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VImgBuilder) Transition(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Crossorigin(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":crossorigin", h.JSONString(v))
	return b
}

func (b *VImgBuilder) Referrerpolicy(v interface{}) (r *VImgBuilder) {
	b.tag.Attr(":referrerpolicy", h.JSONString(v))
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
