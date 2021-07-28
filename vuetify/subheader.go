package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSubheaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSubheader(children ...h.HTMLComponent) (r *VSubheaderBuilder) {
	r = &VSubheaderBuilder{
		tag: h.Tag("v-subheader").Children(children...),
	}
	return
}

func (b *VSubheaderBuilder) Dark(v bool) (r *VSubheaderBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSubheaderBuilder) Inset(v bool) (r *VSubheaderBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VSubheaderBuilder) Light(v bool) (r *VSubheaderBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSubheaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSubheaderBuilder) Attr(vs ...interface{}) (r *VSubheaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSubheaderBuilder) Children(children ...h.HTMLComponent) (r *VSubheaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSubheaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSubheaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSubheaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSubheaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSubheaderBuilder) Class(names ...string) (r *VSubheaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSubheaderBuilder) ClassIf(name string, add bool) (r *VSubheaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSubheaderBuilder) On(name string, value string) (r *VSubheaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSubheaderBuilder) Bind(name string, value string) (r *VSubheaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSubheaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
