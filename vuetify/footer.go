package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFooterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFooter(children ...h.HTMLComponent) (r *VFooterBuilder) {
	r = &VFooterBuilder{
		tag: h.Tag("v-footer").Children(children...),
	}
	return
}

func (b *VFooterBuilder) Absolute(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) App(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Color(v string) (r *VFooterBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VFooterBuilder) Dark(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Elevation(v int) (r *VFooterBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Fixed(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Height(v int) (r *VFooterBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Inset(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Light(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) MaxHeight(v int) (r *VFooterBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) MaxWidth(v int) (r *VFooterBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) MinHeight(v int) (r *VFooterBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) MinWidth(v int) (r *VFooterBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Outlined(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Padless(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":padless", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Rounded(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Shaped(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Tag(v string) (r *VFooterBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VFooterBuilder) Tile(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Width(v int) (r *VFooterBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFooterBuilder) Attr(vs ...interface{}) (r *VFooterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFooterBuilder) Children(children ...h.HTMLComponent) (r *VFooterBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFooterBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFooterBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFooterBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFooterBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFooterBuilder) Class(names ...string) (r *VFooterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFooterBuilder) ClassIf(name string, add bool) (r *VFooterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFooterBuilder) On(name string, value string) (r *VFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFooterBuilder) Bind(name string, value string) (r *VFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFooterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
