package vuetify

import (
	"context"
	"fmt"
	h "github.com/theplant/htmlgo"
)

type VCardBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCard(children ...h.HTMLComponent) (r *VCardBuilder) {
	r = &VCardBuilder{
		tag: h.Tag("v-card").Children(children...),
	}
	return
}
func (b *VCardBuilder) ActiveClass(v string) (r *VCardBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VCardBuilder) Append(v bool) (r *VCardBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Color(v string) (r *VCardBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCardBuilder) Dark(v bool) (r *VCardBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Disabled(v bool) (r *VCardBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Elevation(v int) (r *VCardBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Exact(v bool) (r *VCardBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) ExactActiveClass(v string) (r *VCardBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VCardBuilder) Flat(v bool) (r *VCardBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Height(v int) (r *VCardBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Hover(v bool) (r *VCardBuilder) {
	b.tag.Attr(":hover", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Href(v string) (r *VCardBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VCardBuilder) Img(v string) (r *VCardBuilder) {
	b.tag.Attr("img", v)
	return b
}

func (b *VCardBuilder) Light(v bool) (r *VCardBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) MaxHeight(v int) (r *VCardBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) MaxWidth(v int) (r *VCardBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) MinHeight(v int) (r *VCardBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) MinWidth(v int) (r *VCardBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Nuxt(v bool) (r *VCardBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Raised(v bool) (r *VCardBuilder) {
	b.tag.Attr(":raised", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Replace(v bool) (r *VCardBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Ripple(v bool) (r *VCardBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) Tag(v string) (r *VCardBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCardBuilder) Target(v string) (r *VCardBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VCardBuilder) Tile(v bool) (r *VCardBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) To(v string) (r *VCardBuilder) {
	b.tag.Attr("to", v)
	return b
}

func (b *VCardBuilder) Width(v int) (r *VCardBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VCardBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
