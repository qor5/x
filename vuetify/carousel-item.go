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

func (b *VCarouselItemBuilder) ActiveClass(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VCarouselItemBuilder) Append(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Disabled(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Eager(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Exact(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) ExactActiveClass(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VCarouselItemBuilder) ExactPath(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":exact-path", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Href(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":href", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Link(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Nuxt(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Replace(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) ReverseTransition(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":reverse-transition", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Ripple(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Tag(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCarouselItemBuilder) Target(v string) (r *VCarouselItemBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VCarouselItemBuilder) To(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VCarouselItemBuilder) Transition(v bool) (r *VCarouselItemBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VCarouselItemBuilder) Value(v interface{}) (r *VCarouselItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
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
