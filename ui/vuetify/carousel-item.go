package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCarouselItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCarouselItem(children ...h.HTMLComponent) (r *VCarouselItemBuilder) {
	r = &VCarouselItemBuilder{
		tag: h.Tag("v-carousel-item").Children(children...),
	}
	return
}

func (b *VCarouselItemBuilder) Alt(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("alt", v)
	return b
}

func (b *VCarouselItemBuilder) Cover(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":cover", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Color(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCarouselItemBuilder) Draggable(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":draggable", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Eager(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Gradient(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("gradient", v)
	return b
}

func (b *VCarouselItemBuilder) LazySrc(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("lazy-src", v)
	return b
}

func (b *VCarouselItemBuilder) Options(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Sizes(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("sizes", v)
	return b
}

func (b *VCarouselItemBuilder) Src(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Srcset(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("srcset", v)
	return b
}

func (b *VCarouselItemBuilder) Position(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("position", v)
	return b
}

func (b *VCarouselItemBuilder) AspectRatio(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":aspect-ratio", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) ContentClass(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Inline(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":inline", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Height(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) MaxHeight(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) MaxWidth(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) MinHeight(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) MinWidth(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Width(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Rounded(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Tile(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Transition(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Crossorigin(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":crossorigin", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Referrerpolicy(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":referrerpolicy", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) ReverseTransition(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":reverse-transition", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Value(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Disabled(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) SelectedClass(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VCarouselItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCarouselItemBuilder) Attr(vs ...interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCarouselItemBuilder) Children(children ...h.HTMLComponent) (r *VCarouselItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCarouselItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCarouselItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCarouselItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCarouselItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCarouselItemBuilder) Class(names ...string) (r *VCarouselItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCarouselItemBuilder) ClassIf(name string, add bool) (r *VCarouselItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCarouselItemBuilder) On(name string, value string) (r *VCarouselItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCarouselItemBuilder) Bind(name string, value string) (r *VCarouselItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCarouselItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
