package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSpacerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSpacer(children ...h.HTMLComponent) (r *VSpacerBuilder) {
	r = &VSpacerBuilder{
		tag: h.Tag("v-spacer").Children(children...),
	}
	return
}

func (b *VSpacerBuilder) Tag(v string) (r *VSpacerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VSpacerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSpacerBuilder) Attr(vs ...interface{}) (r *VSpacerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSpacerBuilder) Children(children ...h.HTMLComponent) (r *VSpacerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSpacerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSpacerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSpacerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSpacerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSpacerBuilder) Class(names ...string) (r *VSpacerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSpacerBuilder) ClassIf(name string, add bool) (r *VSpacerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSpacerBuilder) On(name string, value string) (r *VSpacerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSpacerBuilder) Bind(name string, value string) (r *VSpacerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSpacerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
