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
